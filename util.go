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

	words := []string{
		"the",
		"be",
		"of",
		"and",
		"a",
		"to",
		"in",
		"he",
		"have",
		"it",
		"that",
		"for",
		"they",
		"I",
		"with",
		"as",
		"not",
		"on",
		"she",
		"at",
		"by",
		"this",
		"we",
		"you",
		"do",
		"but",
		"from",
		"or",
		"which",
		"one",
		"would",
		"all",
		"will",
		"there",
		"say",
		"who",
		"make",
		"when",
		"can",
		"more",
		"if",
		"no",
		"man",
		"out",
		"other",
		"so",
		"what",
		"time",
		"up",
		"go",
		"about",
		"than",
		"into",
		"could",
		"state",
		"only",
		"new",
		"year",
		"some",
		"take",
		"come",
		"these",
		"know",
		"see",
		"use",
		"get",
		"like",
		"then",
		"first",
		"any",
		"work",
		"now",
		"may",
		"such",
		"give",
		"over",
		"think",
		"most",
		"even",
		"find",
		"day",
		"also",
		"after",
		"way",
		"many",
		"must",
		"look",
		"before",
		"great",
		"back",
		"through",
		"long",
		"where",
		"much",
		"should",
		"well",
		"people",
		"down",
		"own",
		"just",
		"because",
		"good",
		"each",
		"those",
		"feel",
		"seem",
		"how",
		"high",
		"too",
		"place",
		"little",
		"world",
		"very",
		"still",
		"nation",
		"hand",
		"old",
		"life",
		"tell",
		"write",
		"become",
		"here",
		"show",
		"house",
		"both",
		"between",
		"need",
		"mean",
		"call",
		"develop",
		"under",
		"last",
		"right",
		"move",
		"thing",
		"general",
		"school",
		"never",
		"same",
		"another",
		"begin",
		"while",
		"number",
		"part",
		"turn",
		"real",
		"leave",
		"might",
		"want",
		"point",
		"form",
		"off",
		"child",
		"few",
		"small",
		"since",
		"against",
		"ask",
		"late",
		"home",
		"interest",
		"large",
		"person",
		"end",
		"open",
		"public",
		"follow",
		"during",
		"present",
		"without",
		"again",
		"hold",
		"govern",
		"around",
		"possible",
		"head",
		"consider",
		"word",
		"program",
		"problem",
		"however",
		"lead",
		"system",
		"set",
		"order",
		"eye",
		"plan",
		"run",
		"keep",
		"face",
		"fact",
		"group",
		"play",
		"stand",
		"increase",
		"early",
		"course",
		"change",
		"help",
		"line",
	}

	for i := 0; i < n; i++ {
		r += words[rand.Int()%len(words)]
		if i != n-1 {
			r += " "
		}
	}

	return strings.Replace(wordWrap(r, 80), "\n", " \n", -1)
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

func drawCellsAtCenter(scr tcell.Screen, s []cell, cursorIdx int) {
	rows := 0
	cols := 0
	c := 0

	for _, x := range s {
		if x.c == '\n' {
			rows++
			if c > cols {
				cols = c
			}
			c = 0
		} else {
			c++
		}
	}

	rows++
	if c > cols {
		cols = c
	}

	w, h := scr.Size()
	x := (w - cols) / 2
	y := (h - rows) / 2

	drawCells(scr, x, y, s, cursorIdx)
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
