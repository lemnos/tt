package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	hue "github.com/gerow/go-color"
)

var CONFIG_DIRS []string

func init() {
	home, _ := os.LookupEnv("HOME")

	CONFIG_DIRS = []string{
		filepath.Join(home, ".tt"),
		"/etc/tt",
	}
}

type cell struct {
	c     rune //rendered character
	style tcell.Style

	ScreenPos        [2]int
	WaitingStyle     tcell.Style
	ComparedStyle    tcell.Style
	WrongStyle       tcell.Style
	CursorStyle      tcell.Style
	AfterCursorStyle tcell.Style
	Format           int // 0 = normal, 1 = mask
	RubricInd        int
}

func getParagraphs(s string) []string {
	s = strings.Replace(s, "\r", "", -1)
	s = regexp.MustCompile("\n\n+").ReplaceAllString(s, "\n\n")
	return strings.Split(strings.Trim(s, "\n"), "\n\n")
}

var rawReplace map[rune][]rune = map[rune][]rune{
	'\t': []rune("›··"),
	'\n': []rune("↩\n"),
}

// wordWrapBytes wraps a byte slice to a given width marking the end of lines with a newline or vertical tab when raw is true
func wordWrapBytes(s []rune, n int, isRaw bool) {
	sp := 0  // last space
	sz := 0  // current size of line
	lsp := 0 // last space in line
	lineEnd := rune('\n')
	r := make([]rune, len(s))
	copy(r, s)
	s = s[:0]
	if isRaw {
		lineEnd = rune('\r') //group separator
	}
	for i := 0; i < len(r); i++ {
		//save last space
		if r[i] == '\n' || r[i] == '\t' || r[i] == ' ' || r[i] == '\r' {
			sp = len(s)
			lsp = sz
		}
		// add replacement if needed
		if isRaw {
			if r[i] == 0 {
				continue
			}
			new, ok := rawReplace[r[i]]
			if !ok {
				s = append(s, r[i])
				sz++
			} else {
				s = append(s, new...)
				if new[len(new)-1] == '\n' {
					sp = len(s) - 1
					lsp = sz
				}
				sz += len(new)
			}
		} else {
			s = append(s, r[i])
			sz++
		}
		//check if we need to wrap
		if sz > n && sp != 0 {
			if isRaw && s[sz-1] != '\n'{
				s = append(s[:sp+1], append([]rune{lineEnd}, s[sp+1:]...)...)
			} else {
				s[sp] = lineEnd
			}

			sz = sz - lsp
		}
	}
}

// this function changes a slice by reslicing it up to the last non \x00 character
func sliceTrimmer(s []rune) []rune {
	sz := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '\x00' {
			sz = i + 1
			break
		}
	}
	return s[:sz]
}

func wordWrap(s string, wsz int) string {
	var notRaw bool
	r := []rune(s)
	wordWrapBytes(r, wsz, notRaw)
	return string(r)
}

func softWrap(s string, wsz, hsz int) string {
	raw := true
	var r = make([]rune, len(s)*4)
	copy(r, []rune(s))
	wordWrapBytes(r, wsz, raw)
	x := string(sliceTrimmer(r))
	return x
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randomText(n int, words []string) string {
	r := ""

	var last string
	for i := 0; i < n; i++ {
		w := words[rand.Int()%len(words)]
		for last == w {
			w = words[rand.Int()%len(words)]
		}

		r += w
		if i != n-1 {
			r += " "
		}

		last = w
	}

	return strings.Replace(r, "\n", " \n", -1)
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

func drawStringAtCenter(scr tcell.Screen, s string, style tcell.Style) {
	nc, nr := calcStringDimensions(s)
	sw, sh := scr.Size()

	x := (sw - nc) / 2
	y := (sh - nr) / 2

	drawString(scr, x, y, s, -1, style)
}

// gives boundarys of a textbox by finding max number of columns and rows in a string when rendered
func calcStringDimensions(s string) (nc, nr int) {
	if s == "" {
		return 0, 0
	}

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

func calcStringDimensions_raw(s string) (nc, nr int) {
	if s == "" {
		return 0, 0
	}

	c := 0

	for _, x := range s {
		if x == '\r' || x == '\n' {
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

func newTcellColor(s string) (tcell.Color, error) {
	if len(s) != 7 || s[0] != '#' {
		return 0, fmt.Errorf("%s is not a valid hex color", s)
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

	return tcell.NewRGBColor(r, g, b), nil
}

func readResource(typ, name string) []byte {
	if name == "-" {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		return b
	}

	if b, err := ioutil.ReadFile(name); err == nil {
		return b
	}

	for _, d := range CONFIG_DIRS {
		if b, err := ioutil.ReadFile(filepath.Join(d, typ, name)); err == nil {
			return b
		}
	}

	return readPackedFile(filepath.Join(typ, name))
}

func DimColor(color tcell.Color) tcell.Color {

	r, g, b := color.RGB()
	hsl := hue.RGB{R: float64(r), G: float64(g), B: float64(b)}.ToHSL()
	hsl.L *= 0.5
	hsl.S *= 0.8
	new := hsl.ToRGB()
	return tcell.NewRGBColor(int32(new.R), int32(new.G), int32(new.B))
}