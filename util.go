package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gdamore/tcell"
)

type cell struct {
	c     rune
	style tcell.Style
}

func dbgPrintf(scr tcell.Screen, format string, args ...interface{}) {
	for i := 0; i < 80; i++ {
		for j := 0; j < 80; j++ {
			scr.SetContent(i, j, ' ', nil, tcell.StyleDefault)
		}
	}
	drawString(scr, 0, 0, fmt.Sprintf(format, args...), -1, tcell.StyleDefault)
}

func wordWrapBytes(s []byte, n int) {
	sp := 0
	sz := 0

	for i := 0; i < len(s); i++ {
		sz++

		if s[i] == '\n' {
			s[i] = ' '
		}

		if s[i] == ' ' {
			sp = i
		}

		if sz > n {
			if sp != 0 {
				s[sp] = '\n'
			}

			sz = i - sp
		}
	}

}

func wordWrap(s string, n int) string {
	r := []byte(s)
	wordWrapBytes(r, n)
	return string(r)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randomText(n int) string {
	r := ""

	for i := 0; i < n; i++ {
		r += defaultWordList[rand.Int()%len(defaultWordList)]
		if i != n-1 {
			r += " "
		}
	}

	return strings.Replace(r, "\n", " \n", -1)
}

func stringToCells(s string) []cell {
	a := make([]cell, len(s))
	s = strings.TrimRight(s, "\n ")

	len := 0
	for _, r := range s {
		a[len].c = r
		a[len].style = tcell.StyleDefault
		len++
	}

	return a[:len]
}

func drawString(scr tcell.Screen, x, y int, s string, cursorIdx int, style tcell.Style) {
	sx := x

	for i, c := range s {
		if c == '\n' {
			y++
			x = sx
		} else {
			scr.SetContent(x, y, c, nil, style)
			if i == cursorIdx {
				scr.ShowCursor(x, y)
			}

			x++
		}
	}

	if cursorIdx == len(s) {
		scr.ShowCursor(x, y)
	}
}

func drawCells(scr tcell.Screen, x, y int, s []cell, cursorIdx int) {
	sx := x

	for i, c := range s {
		if c.c == '\n' {
			y++
			x = sx
		} else {
			scr.SetContent(x, y, c.c, nil, c.style)
			if i == cursorIdx {
				scr.ShowCursor(x, y)
			}

			x++
		}
	}

	if cursorIdx == len(s) {
		scr.ShowCursor(x, y)
	}
}

func drawStringAtCenter(scr tcell.Screen, s string, style tcell.Style) {
	nc, nr := calcStringDimensions(s)
	sw, sh := scr.Size()

	x := (sw - nc) / 2
	y := (sh - nr) / 2

	drawString(scr, x, y, s, -1, style)
}

func calcStringDimensions(s string) (nc, nr int) {
	c := 0

	for _, x := range s {
		if x == '\n' {
			nr++
			if c > nc {
				nc = c
			}
			c = 0
		} else {
			c++
		}
	}

	nr++
	if c > nc {
		nc = c
	}

	return
}

func newTcellColor(s string) tcell.Color {
	if len(s) != 7 || s[0] != '#' {
		panic(fmt.Errorf("%s is not a valid hex color", s))
	}

	tonum := func(c byte) int32 {
		if c > '9' {
			if c >= 'a' {
				return (int32)(c - 'a' + 10)
			} else {
				return (int32)(c - 'A' + 10)
			}
		} else {
			return (int32)(c - '0')
		}
	}

	r := tonum(s[1])<<4 | tonum(s[2])
	g := tonum(s[3])<<4 | tonum(s[4])
	b := tonum(s[5])<<4 | tonum(s[6])

	return tcell.NewRGBColor(r, g, b)
}
