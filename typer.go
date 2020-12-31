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
	TyperResize
)

type typer struct {
	Scr      tcell.Screen
	OnStart  func()
	SkipWord bool
	ShowWpm  bool
	tty      io.Writer

	currentWordStyle    tcell.Style
	nextWordStyle       tcell.Style
	incorrectSpaceStyle tcell.Style
	incorrectStyle      tcell.Style
	correctStyle        tcell.Style
	backgroundStyle     tcell.Style
}

func NewTyper(scr tcell.Screen, fgcol, bgcol, hicol, hicol2, hicol3, errcol tcell.Color) *typer {
	var tty io.Writer
	def := tcell.StyleDefault.
		Foreground(fgcol).
		Background(bgcol)

	tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	//Will fail on windows, but tt is still mostly usable via tcell
	if err != nil {
		tty = ioutil.Discard
	}

	return &typer{
		Scr:      scr,
		SkipWord: true,
		tty:      tty,

		backgroundStyle:     def,
		correctStyle:        def.Foreground(hicol),
		currentWordStyle:    def.Foreground(hicol2),
		nextWordStyle:       def.Foreground(hicol3),
		incorrectStyle:      def.Foreground(errcol),
		incorrectSpaceStyle: def.Background(errcol),
	}
}

func (t *typer) highlight(text []cell, idx int, currentWordStyle, nextWordStyle tcell.Style) {
	for ; idx < len(text) && text[idx].c != ' ' && text[idx].c != '\n'; idx++ {
		text[idx].style = currentWordStyle
	}

	for ; idx < len(text) && (text[idx].c == ' ' || text[idx].c == '\n'); idx++ {
	}

	for ; idx < len(text) && text[idx].c != ' ' && text[idx].c != '\n'; idx++ {
		text[idx].style = nextWordStyle
	}
}

func (t *typer) Start(text []string, timeout time.Duration) (nerrs, ncorrect int, duration time.Duration, rc int) {
	timeLeft := timeout

	for i, p := range text {
		startImmediately := true
		var d time.Duration
		var e, c int

		if i == 0 {
			startImmediately = false
		}

		e, c, rc, d = t.start(p, timeLeft, startImmediately)

		nerrs += e
		ncorrect += c
		duration += d

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

func (t *typer) start(s string, timeLimit time.Duration, startImmediately bool) (nerrs int, ncorrect int, rc int, duration time.Duration) {
	var startTime time.Time
	text := stringToCells(s)

	sw, sh := scr.Size()
	nc, nr := calcStringDimensions(s)
	x := (sw - nc) / 2
	y := (sh - nr) / 2

	for i, _ := range text {
		text[i].style = t.backgroundStyle
	}

	t.tty.Write([]byte("\033[5 q"))

	//Assumes original cursor shape was a block (the one true cursor shape), there doesn't appear to be a
	//good way to save/restore the shape if the user has changed it from the otcs.
	defer t.tty.Write([]byte("\033[2 q"))

	t.Scr.SetStyle(t.backgroundStyle)
	idx := 0

	calcStats := func() {
		nerrs = 0
		ncorrect = 0

		for _, c := range text {
			if c.style == t.incorrectStyle || c.style == t.incorrectSpaceStyle {
				nerrs++
			} else if c.style == t.correctStyle {
				ncorrect++
			}
		}

		rc = TyperComplete
		duration = time.Now().Sub(startTime)
	}

	redraw := func() {
		if timeLimit != -1 && !startTime.IsZero() {
			remaining := timeLimit - time.Now().Sub(startTime)
			drawString(t.Scr, x+nc/2, y+nr+1, "      ", -1, t.backgroundStyle)
			drawString(t.Scr, x+nc/2, y+nr+1, strconv.Itoa(int(remaining/1E9)), -1, t.backgroundStyle)
		}

		if t.ShowWpm && !startTime.IsZero() {
			calcStats()
			if duration > 1E7 { //Avoid flashing large numbers on test start.
				wpm := int((float64(ncorrect) / 5) / (float64(duration) / 60E9))
				drawString(t.Scr, x+nc/2-4, y-2, fmt.Sprintf("WPM: %-10d\n", wpm), -1, t.backgroundStyle)
			}
		}

		//Potentially inefficient, but seems to be good enough

		drawCells(t.Scr, x, y, text, idx)

		t.Scr.Show()
	}

	deleteWord := func() {
		t.highlight(text, idx, t.backgroundStyle, t.backgroundStyle)

		if idx == 0 {
			return
		}

		idx--

		for idx > 0 && (text[idx].c == ' ' || text[idx].c == '\n') {
			text[idx].style = t.backgroundStyle
			idx--
		}

		for idx > 0 && text[idx].c != ' ' && text[idx].c != '\n' {
			text[idx].style = t.backgroundStyle
			idx--
		}

		if text[idx].c == ' ' || text[idx].c == '\n' {
			idx++
		}

		t.highlight(text, idx, t.currentWordStyle, t.nextWordStyle)
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

			time.Sleep(time.Duration(5E8))
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
		t.highlight(text, idx, t.currentWordStyle, t.nextWordStyle)
		redraw()

		ev := t.Scr.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			rc = TyperResize
			return
		case *tcell.EventKey:
			if runtime.GOOS != "windows" && ev.Key() == tcell.KeyBackspace { //Control+backspace on unix terms
				deleteWord()
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
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if ev.Modifiers() == tcell.ModAlt || ev.Modifiers() == tcell.ModCtrl {
					deleteWord()
				} else {
					t.highlight(text, idx, t.backgroundStyle, t.backgroundStyle)

					if idx == 0 {
						break
					}

					idx--

					for idx > 0 && text[idx].c == '\n' {
						idx--
					}

					text[idx].style = t.backgroundStyle

					t.highlight(text, idx, t.currentWordStyle, t.nextWordStyle)
				}
			case tcell.KeyRune:
				if idx < len(text) {
					switch {
					case ev.Rune() == text[idx].c:
						text[idx].style = t.correctStyle
						idx++
					case ev.Rune() == ' ' && t.SkipWord:
						if idx > 0 && text[idx-1].c == ' ' && text[idx].c != ' ' { //Do nothing on word boundaries.
							break
						}

						for idx < len(text) && text[idx].c != ' ' && text[idx].c != '\n' {
							text[idx].style = t.incorrectStyle
							idx++
						}

						if idx < len(text) {
							text[idx].style = t.incorrectSpaceStyle
							idx++
						}
					default:
						if text[idx].c == ' ' {
							text[idx].style = t.incorrectSpaceStyle
						} else {
							text[idx].style = t.incorrectStyle
						}
						idx++
					}

					for idx < len(text) && text[idx].c == '\n' {
						idx++
					}

					t.highlight(text, idx, t.currentWordStyle, t.nextWordStyle)
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
