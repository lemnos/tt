package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/gdamore/tcell"
)

const (
	TyperComplete = iota
	TyperSigInt
	TyperEscape
	TyperPrevious
	TyperNext
	TyperResize
)

type segment struct {
	Text        string `json:"text"`
	Attribution string `json:"attribution"`
	renderForm  string
}

func (s *segment) Render() string {
	if s.renderForm == "" {
		return s.Text
	}
	return s.renderForm
}

func (s *segment) SetRenderForm(form string) {
	s.renderForm = form
}

func (s *segment) SetContent(content string) {
	s.Text = content
}

type mistake struct {
	Word  string `json:"word"`
	Typed string `json:"typed"`
}

type typer struct {
	Scr              tcell.Screen
	OnStart          func()
	SkipWord         bool
	ShowWpm          bool
	DisableBackspace bool
	BlockCursor      bool
	tty              io.Writer

	currentWordStyle    tcell.Style
	nextWordStyle       tcell.Style
	incorrectSpaceStyle tcell.Style
	incorrectStyle      tcell.Style
	correctStyle        tcell.Style
	defaultStyle        tcell.Style
	maskedWsStyle       tcell.Style
	passedWsStyle       tcell.Style
}

// position key is [position in rendered text], position in textbox is an array field of cell
type textBox map[int]cell

func (txbx textBox) ShowCursor(cursor int, t *typer, rawmode bool) int {
	var skips int
	w, z := txbx[cursor].ScreenPos[0], txbx[cursor].ScreenPos[1]
	if w <= 0 || z <= 0 {
		if !rawmode && txbx[cursor].c != 0 {
			t.Scr.HideCursor()
		} else if rawmode && txbx[cursor].c != 0 {
			skips += txbx.ShowCursor(cursor+1, t, rawmode)
			return skips
		}
		w, z = txbx[cursor-1].ScreenPos[0], txbx[cursor-1].ScreenPos[1]
		w++
	}
	t.Scr.ShowCursor(w, z)
	return skips
}

func NewTyper(scr tcell.Screen, emboldenTypedText bool, fgcol, bgcol, hicol, hicol2, hicol3, errcol tcell.Color) *typer {
	var tty io.Writer
	def := tcell.StyleDefault.
		Foreground(fgcol).
		Background(bgcol)

	tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	//Will fail on windows, but tt is still mostly usable via tcell
	if err != nil {
		tty = ioutil.Discard
	}

	correctStyle := def.Foreground(hicol)
	if emboldenTypedText {
		correctStyle = correctStyle.Bold(true)
	}
	_, xBg, _ := def.Decompose()
	xFg := tcell.Color(DimColor(hicol2))
	wsChar := def.Foreground(DimColor(xBg)).Background(bgcol)
	return &typer{
		Scr:      scr,
		SkipWord: true,
		tty:      tty,

		defaultStyle:        def,
		correctStyle:        correctStyle,
		currentWordStyle:    def.Foreground(hicol2),
		nextWordStyle:       def.Foreground(hicol3),
		incorrectStyle:      def.Foreground(errcol),
		incorrectSpaceStyle: def.Background(errcol),
		maskedWsStyle:       def.Foreground(xFg),
		passedWsStyle:       wsChar,
	}
}

func (t *typer) Start(test []segment, timeout time.Duration, rawMode bool) (nerrs, ncorrect int, duration time.Duration, rc int, mistakes []mistake) {
	timeLeft := timeout
	for i, segment := range test {
		startImmediately := true
		var d time.Duration
		var e, c int
		var m []mistake

		if i == 0 {
			startImmediately = false
		}

		e, c, rc, d, m = t.processSegment(segment, timeLeft, startImmediately, rawMode) //errors, correct, return code, duration, mistakes

		nerrs += e
		ncorrect += c
		duration += d
		mistakes = append(mistakes, m...)
		if timeLeft != -1 {
			timeLeft -= d
		}
		if timeLeft <= 0 && timeout > 0 {
			break
		}
		if rc != TyperComplete {
			return
		}
	}
	return
}

func extractMistypedWords(text []rune, typed []rune, rawMode bool) (mistakes []mistake) {
	mistakes = make([]mistake, 0, 20)
	var w []rune
	var t []rune
	f := false

	for i := range text {
		if text[i] == 32 || text[i] == 10 || (rawMode && (text[i] == 32 || text[i] == 9 || text[i] == 10)) {
			//save and reset
			if f {
				mistakes = append(mistakes, mistake{string(w), string(t)})
			}

			w = w[:0]
			t = t[:0]
			f = false
			if !rawMode {
				continue
			}
		}
		if text[i] != typed[i] && text[i] == 10 && typed[i] == 13 {
			typed[i] = '\n'
		} else if text[i] != typed[i] {
			f = true
		}

		if text[i] == 0 {
			w = append(w, '_')
		} else {
			w = append(w, text[i])
		}

		if typed[i] == 0 {
			t = append(t, '_')
		} else {
			t = append(t, typed[i])
		}
	}

	if f {
		mistakes = append(mistakes, mistake{string(w), string(t)})
	}

	return
}

func (t *typer) processSegment(excersice segment, timeLimit time.Duration, startImmediately bool, rawMode bool) (nerrs int, ncorrect int, rc int, duration time.Duration, mistakes []mistake) {
	//attribution := excersice.Attribution
	//remove newlines from text
	var rubric = []rune(excersice.Text)
	text := []rune(excersice.Render())
	typed := make([]rune, len(text))
	var startTime time.Time
	sw, sh := scr.Size()
	nc, nr := calcStringDimensions_raw(excersice.Render())
	x := (sw - nc) / 2
	y := (sh - nr) / 2

	typingExrc := typeSet(text, x, y, t, rawMode) //state

	if !t.BlockCursor {
		t.tty.Write([]byte("\033[5 q"))

		//Assumes original cursor shape was a block (the one true cursor shape), there doesn't appear to be a
		//good way to save/restore the shape if the user has changed it from the otcs.
		defer t.tty.Write([]byte("\033[2 q"))
	}

	t.Scr.SetStyle(t.defaultStyle)

	pointer := 0 //Tracks typed and rubric
	cursor := 0  //Tracks display text, do not use for evaluations

	calcStats := func() {
		nerrs = 0
		ncorrect = 0

		mistakes = extractMistypedWords(rubric[:pointer], typed[:pointer], rawMode)

		for iter := 0; iter < pointer; iter++ {
			if rubric[iter] != typed[iter] {
				nerrs++
			} else {
				ncorrect++
			}

			if rubric[iter] != typed[iter] && rubric[iter] == '\n' && typed[iter] == '\r' {
				nerrs--
				ncorrect++
			}

		}
	}

	draw := func(txb textBox) {
		for _, content := range txb {
			if content.c == '\r' {
				continue
			}
			t.Scr.SetContent(content.ScreenPos[0], content.ScreenPos[1], content.c, nil, content.style)
		}
		t.Scr.Show()
	}

	advance := func(ca int, r rune) { //[c]ursor [a]dvance
		typed[pointer] = r
		pointer++
		cursor += ca

	}

	deleteWord := func() {
		if cursor == 0 {
			return
		}
		cursor--

		for cursor > 0 && text[cursor] != ' ' && text[cursor] != '\n' && text[cursor] != '·' && text[cursor] != '↩' && text[cursor] != '\r' {
			cursor--
		}
		if text[cursor] == '·' || text[cursor] == '\n' {
			cursor -= 2
		}
	}

	redrawChanges := func() bool {
		var ss, sa int // [s]tate [s]tyle and [s]lot [a]djustment
		rltvPos := []int{wordBefor, wordBefor, wordBefor, currentWord, nextWord, scndWord, unstagedWord, unstagedWord}
		//working with character windows helps
		wbIndex := defineWindow(rubric, pointer) //wordBoundary index
		if wbIndex[0] == -1 {
			return true
		}

		//left truncate the relative position array to the available window
		if len(wbIndex) < len(rltvPos) {
			if pointer-1 < wbIndex[0] {
				sa = 0
				//prepend 0 to wbIndex
				wbIndex = append(wbIndex, 0)
				copy(wbIndex[1:], wbIndex)
				wbIndex[0] = 0
			} else {
				for iter := range wbIndex {
					if pointer <= wbIndex[iter+1] && pointer >= wbIndex[iter] {
						sa = iter
						break
					}
				}
			}
			rltvPos = rltvPos[3-sa:]
		}

		pos := wbIndex[0] //tracks CURSOR
		pnt := wbIndex[0] //tracks POINTER
		// traverse each of the word boundaries
		// search for the current cell where cc.RubricInd == pointer
		for i := wbIndex[0]; i < wbIndex[len(wbIndex)-1]; i++ {
			if typingExrc[i].RubricInd == pnt {
				pos = i
				break
			}
		}
	wb:
		for i := range wbIndex[:len(wbIndex)-1] {
		word:
			for pnt >= wbIndex[i] && pnt <= wbIndex[i+1] || pos == len(text)-1 {
				ss = rltvPos[i] //relative position of the word to the cursor
				if ss == currentWord && pointer > pnt {
					ss = wordBefor
				}
				cc := typingExrc[pos] //current cell
				if cc.c == 0 {
					pnt++
					pos++
					continue word
				}
				cc.stylize(ss, typed[cc.RubricInd] == rubric[cc.RubricInd])
				typingExrc[pos] = cc // update state
				pos++
				if cc.Format == 0 {
					pnt++
					continue word
				}
				//when masked process accordingly depending on what symbol comes first '›' or  '↩' use a switch statements
				switch cc.c {
				case '›':
					for mask := 0; mask < 2; mask++ {
						cc := typingExrc[pos] //current cell
						cc.stylize(ss, typed[cc.RubricInd] == rubric[cc.RubricInd])
						typingExrc[pos] = cc // update state
						pos++
					}
					if rubric[pnt] != 9 && rubric[pnt] != 10 {
						ss++
					} else {
						continue wb
					}

				case '↩':
					pnt++
					pos++
					ss++
				}

			}
			if ss < 4 {
				ss++
			}
		}
		draw(typingExrc)
		return false
	}

	///**************  MAIN EVENT LOOP  *****************///

	tickerCloser := make(chan bool)

	//Inject nil events into the main event loop at regular invervals to force an update
	ticker := func() {
		for {
			select {
			case <-tickerCloser:
				return
			default:
			}

			time.Sleep(time.Duration(5e8))
			t.Scr.PostEventWait(nil)
		}
	}

	go ticker()
	defer close(tickerCloser)

	startTime = time.Now()
	rc = TyperComplete
	//if starttime is set, then the timer is running
	// duration = time.Since(startTime)

	//redraw function is called for every non- nil event
	t.Scr.Clear()
	t.Scr.ShowCursor(x, y)
	_ = t.Scr.PollEvent() // ignores unnecessary resize event at the start
	draw(typingExrc)

	var (
		started     bool
		wpsTicker   *time.Ticker = time.NewTicker(time.Second * 1e9)
		charsNow    int
		charsBefore int
	)
listening:
	for {
		ev := t.Scr.PollEvent()
		if t.ShowWpm {
			select {
			case <-wpsTicker.C:
				charsNow = pointer
				charsP2S := charsNow - charsBefore
				//wpm is calculated based on the groups of 5characters typed in 2 seconds
				// x chars/2 seconds * 60 sec*word/5 chars*min
				wpm := charsP2S * 6
				drawString(t.Scr, x+nc/2-4, y-2, fmt.Sprintf("WPM: %-10d\n", wpm), -1, t.defaultStyle)
				charsBefore = charsNow
				continue listening

			default:
				if ev == nil {
					continue listening
				}
			}
		} else if ev == nil {
			continue listening
		}
		if !started {
			wpsTicker = time.NewTicker(time.Second * 2)
			started = true
		}
		switch ev := ev.(type) {
		case *tcell.EventResize:
			rc = TyperResize
			return
		case *tcell.EventKey:
			if runtime.GOOS != "windows" && ev.Key() == tcell.KeyBackspace { //Control+backspace on unix terms
				if !t.DisableBackspace {
					deleteWord()
				}
				continue
			}

			if startTime.IsZero() {
				startTime = time.Now()
			}

			switch key := ev.Key(); key {
			case tcell.KeyCtrlC:
				rc = TyperSigInt
				return

			case tcell.KeyEscape:
				rc = TyperEscape
				return

			case tcell.KeyCtrlL:
				t.Scr.Sync()

			case tcell.KeyRight:
				rc = TyperNext
				return

			case tcell.KeyLeft:
				rc = TyperPrevious
				return

			case tcell.KeyCtrlW:
				if !t.DisableBackspace {
					deleteWord()
				}

			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if !t.DisableBackspace {
					if ev.Modifiers() == tcell.ModAlt || ev.Modifiers() == tcell.ModCtrl {
						deleteWord()
					} else {
						if cursor == 0 {
							break
						}

						cursor--
						pointer--

						for cursor > 0 && text[cursor] == '\n' {
							cursor--
							pointer--
						}

						if rawMode && text[cursor] == '·' {
							cursor -= 2
						}

					}
				}
			case tcell.KeyRune, tcell.KeyCtrlI, tcell.KeyEnter:
				//ignore tab and new line when not in raw mode
				if !rawMode && (ev.Rune() == 9 || ev.Rune() == 10) {
					continue
				}
				if cursor < len(text) {
					if t.SkipWord && ev.Rune() == 32 {
						if cursor > 0 && text[cursor-1] == 32 && text[cursor] != 32 { //Do nothing on word boundaries.
							break
						} else if text[cursor] == '›' {
							advance(3, 32)

						}
						for cursor < len(text) && text[cursor] != 32 && text[cursor] != '\n' {
							advance(1, 0)
						}

						if cursor < len(text) {
							advance(1, text[pointer])
						}
					} else {
						if rawMode {
							if rubric[pointer] == 9 {
								advance(3, 9)
							} else if rubric[pointer] == 10 || rubric[pointer] == 13 {
								advance(2, 10)
							} else {
								advance(1, ev.Rune())
							}

						} else {
							advance(1, ev.Rune())

						}
					}
				}
			}

			cursor += typingExrc.ShowCursor(cursor, t, rawMode)

			fini := redrawChanges()
			if fini {
				duration = time.Since(startTime)
				calcStats()
				return
			}
		}
	}
}

func typeSet(S []rune, x, y int, t *typer, isRaw bool) textBox {

	whatStyles := func(s rune) (waiting tcell.Style, compared tcell.Style, incorrect tcell.Style) {
		switch s {
		case '›', '·', '↩', '\n', ' ':
			return t.maskedWsStyle, t.passedWsStyle, t.incorrectSpaceStyle
		default:
			return t.defaultStyle, t.correctStyle, t.incorrectStyle

		}

	}

	whatFunction := func(s rune) int {
		const (
			normal int = iota
			mask
		)
		switch s {
		case '›', '·', '↩', '\n':
			return mask
		default:
			return normal
		}
	}

	var current cell
	var r textBox = make(map[int]cell)
	var breakchar rune
	var cx, cy int = x, y
	if isRaw {
		breakchar = '\r'
	} else {
		breakchar = '\n'
	}

	rubricPos := 0
	wordNum := 1
	for iter := range S {
		//s is prerendered and masking wont be affected
		if S[iter] == breakchar || S[iter] == '\n' {
			cx = x
			cy++
			wordNum++
			rubricPos++
			continue
		} else if S[iter] == ' ' {
			wordNum++
		}

		current.c = S[iter]
		current.WaitingStyle, current.ComparedStyle, current.WrongStyle = whatStyles(S[iter])
		current.CursorStyle = t.currentWordStyle
		current.AfterCursorStyle = t.nextWordStyle
		current.Format = whatFunction(S[iter])
		current.ScreenPos = [2]int{cx, cy}
		current.RubricInd = rubricPos
		if wordNum > 4 {
			current.stylize(unstagedWord, false)
		} else {
			current.stylize(wordNum, false)
		}
		r[iter] = current
		cx++
		//skips masks
		if current.Format == 1 {
			switch S[iter] {
			case '›', '↩':
				continue
			case '·':
				if S[iter-1] == '›' {
					continue
				}
			}

		}
		rubricPos++

	}
	return r
}

// This is the simple case.
// this includes↩\na newline and›••tabs.
func defineWindow(rubric []rune, pointer int) (wordBoundaryIndex []int) {
	if pointer > len(rubric)-1 {
		return []int{-1}
	}
	counti := -1 //value of counts is the number of words
	countj := -1
	i := pointer
	j := pointer
	//check if previous and current cursor are word boundaries

	for ; j < len(rubric)-1; j++ {
		if runeInSlice(rubric[j], []rune{' ', '↩', '\t', '\r', '›', '\n'}) {
			countj++
			wordBoundaryIndex = append(wordBoundaryIndex, j)
			if countj == 3 {
				break
			}
		}
	}
	if i < 0 {
		i = 0
	}
	wordBoundaryIndex = append(wordBoundaryIndex, j)
	for ; i > 0; i-- {
		if runeInSlice(rubric[i], []rune{' ', '\t', '\r', '·', '\n'}) {
			if i > 0 && rubric[i-1] == '›' {
				continue
			}
			counti++
			wordBoundaryIndex = append(wordBoundaryIndex, i)
			if counti == 2 {
				break
			}
		}
	}
	if i < 0 {
		wordBoundaryIndex = append(wordBoundaryIndex, i)
	}

	//make map with word boundaries
	x := map[int]bool{}
	for _, v := range wordBoundaryIndex {
		x[v] = true
	}
	wordBoundaryIndex = wordBoundaryIndex[:0]
	for k := range x {
		wordBoundaryIndex = append(wordBoundaryIndex, k)
	}
	sort.Ints(wordBoundaryIndex)
	return
}

func runeInSlice(r rune, slice []rune) bool {
	for _, s := range slice {
		if r == s {
			return true
		}
	}
	return false
}

const (
	wordBefor int = iota
	currentWord
	nextWord
	scndWord
	unstagedWord
)

func (cx *cell) stylize(rltv int, correct bool) {
	var change tcell.Style

	styler := []tcell.Style{cx.ComparedStyle, cx.CursorStyle, cx.AfterCursorStyle, cx.WaitingStyle, cx.WaitingStyle}

	if rltv == wordBefor && !correct {
		change = cx.WrongStyle
	} else {
		change = styler[rltv]
	}
	cx.style = change
}
