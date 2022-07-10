package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
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
	}
}

func (t *typer) Start(text []segment, timeout time.Duration) (nerrs, ncorrect int, duration time.Duration, rc int, mistakes []mistake) {
	timeLeft := timeout

	for i, s := range text {
		startImmediately := true
		var d time.Duration
		var e, c int
		var m []mistake

		if i == 0 {
			startImmediately = false
		}

		e, c, rc, d, m = t.start(s.Text, timeLeft, startImmediately, s.Attribution)

		nerrs += e
		ncorrect += c
		duration += d
		mistakes = append(mistakes, m...)

		if timeout != -1 {
			timeLeft -= d
			if timeLeft <= 0 {
				return
			}
		}

		if rc != TyperComplete {
			return
		}
	}

	return
}

func extractMistypedWords(text []rune, typed []rune) (mistakes []mistake) {
	var w []rune
	var t []rune
	f := false

	for i := range text {
		if text[i] == ' ' {
			if f {
				mistakes = append(mistakes, mistake{string(w), string(t)})
			}

			w = w[:0]
			t = t[:0]
			f = false
			continue
		}

		if text[i] != typed[i] {
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

func (t *typer) start(s string, timeLimit time.Duration, startImmediately bool, attribution string) (nerrs int, ncorrect int, rc int, duration time.Duration, mistakes []mistake) {
	var startTime time.Time
	text := []rune(s)
	typed := make([]rune, len(text))

	sw, sh := scr.Size()
	nc, nr := calcStringDimensions(s)
	x := (sw - nc) / 2
	y := (sh - nr) / 2

	if !t.BlockCursor {
		t.tty.Write([]byte("\033[5 q"))

		//Assumes original cursor shape was a block (the one true cursor shape), there doesn't appear to be a
		//good way to save/restore the shape if the user has changed it from the otcs.
		defer t.tty.Write([]byte("\033[2 q"))
	}

	t.Scr.SetStyle(t.defaultStyle)
	idx := 0

	calcStats := func() {
		nerrs = 0
		ncorrect = 0

		mistakes = extractMistypedWords(text[:idx], typed[:idx])

		for i := 0; i < idx; i++ {
			if text[i] != '\n' {
				if text[i] != typed[i] {
					nerrs++
				} else {
					ncorrect++
				}
			}
		}

		rc = TyperComplete
		duration = time.Now().Sub(startTime)
	}

	redraw := func() {
		cx := x
		cy := y
		inword := -1

		for i := range text {
			style := t.defaultStyle

			if text[i] == '\n' {
				cy++
				cx = x
				if inword != -1 {
					inword++
				}
				continue
			}

			if i == idx {
				scr.ShowCursor(cx, cy)
				inword = 0
			}

			if i >= idx {
				if text[i] == ' ' {
					inword++
				} else if inword == 0 {
					style = t.currentWordStyle
				} else if inword == 1 {
					style = t.nextWordStyle
				} else {
					style = t.defaultStyle
				}
			} else if text[i] != typed[i] {
				if text[i] == ' ' {
					style = t.incorrectSpaceStyle
				} else {
					style = t.incorrectStyle
				}
			} else {
				style = t.correctStyle
			}

			scr.SetContent(cx, cy, text[i], nil, style)
			cx++
		}

		aw, ah := calcStringDimensions(attribution)
		drawString(t.Scr, x+nc-aw, y+nr+1, attribution, -1, t.defaultStyle)

		if timeLimit != -1 && !startTime.IsZero() {
			remaining := timeLimit - time.Now().Sub(startTime)
			drawString(t.Scr, x+nc/2, y+nr+ah+1, "      ", -1, t.defaultStyle)
			drawString(t.Scr, x+nc/2, y+nr+ah+1, strconv.Itoa(int(remaining/1e9)+1), -1, t.defaultStyle)
		}

		if t.ShowWpm && !startTime.IsZero() {
			calcStats()
			if duration > 1e7 { //Avoid flashing large numbers on test start.
				wpm := int((float64(ncorrect) / 5) / (float64(duration) / 60e9))
				drawString(t.Scr, x+nc/2-4, y-2, fmt.Sprintf("WPM: %-10d\n", wpm), -1, t.defaultStyle)
			}
		}

		//Potentially inefficient, but seems to be good enough

		t.Scr.Show()
	}

	deleteWord := func() {
		if idx == 0 {
			return
		}

		idx--

		for idx > 0 && (text[idx] == ' ' || text[idx] == '\n') {
			idx--
		}

		for idx > 0 && text[idx] != ' ' && text[idx] != '\n' {
			idx--
		}

		if text[idx] == ' ' || text[idx] == '\n' {
			typed[idx] = text[idx]
			idx++
		}
	}

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

	if startImmediately {
		startTime = time.Now()
	}

	t.Scr.Clear()
	for {
		redraw()

		ev := t.Scr.PollEvent()

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
						if idx == 0 {
							break
						}

						idx--

						for idx > 0 && text[idx] == '\n' {
							idx--
						}
					}
				}
			case tcell.KeyRune:
				if idx < len(text) {
					if t.SkipWord && ev.Rune() == ' ' {
						if idx > 0 && text[idx-1] == ' ' && text[idx] != ' ' { //Do nothing on word boundaries.
							break
						}

						for idx < len(text) && text[idx] != ' ' && text[idx] != '\n' {
							typed[idx] = 0
							idx++
						}

						if idx < len(text) {
							typed[idx] = text[idx]
							idx++
						}
					} else {
						typed[idx] = ev.Rune()
						idx++
					}

					for idx < len(text) && text[idx] == '\n' {
						typed[idx] = text[idx]
						idx++
					}
				}

				if idx == len(text) {
					calcStats()
					return
				}
			}
		default: //tick
			if timeLimit != -1 && !startTime.IsZero() && timeLimit <= time.Now().Sub(startTime) {
				calcStats()
				return
			}

			redraw()
		}
	}
}
