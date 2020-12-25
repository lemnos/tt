package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-isatty"
)

var scr tcell.Screen
var csvMode bool
var rawMode bool

type result struct {
	wpm      int
	cpm      int
	accuracy float64
}

var results []result

func readConfig() map[string]string {
	cfg := map[string]string{}

	home, _ := os.LookupEnv("HOME")
	path := filepath.Join(home, ".ttrc")

	if b, err := ioutil.ReadFile(path); err == nil {
		for _, ln := range bytes.Split(b, []byte("\n")) {
			a := strings.SplitN(string(ln), ":", 2)
			if len(a) == 2 {
				cfg[a[0]] = strings.Trim(a[1], " ")
			}
		}
	}

	return cfg
}

func exit() {
	scr.Fini()

	if csvMode {
		for _, r := range results {
			fmt.Printf("%d,%d,%.2f\n", r.wpm, r.cpm, r.accuracy)
		}
	}

	os.Exit(0)
}

func showReport(scr tcell.Screen, cpm, wpm int, accuracy float64) {
	report := fmt.Sprintf("WPM: %d\nCPM: %d\nAccuracy: %.2f%%\n", wpm, cpm, accuracy)

	scr.Clear()
	drawCellsAtCenter(scr, stringToCells(report), -1)
	scr.HideCursor()
	scr.Show()

	for {
		if key, ok := scr.PollEvent().(*tcell.EventKey); ok && key.Key() == tcell.KeyEscape {
			return
		} else if ok && key.Key() == tcell.KeyCtrlC {
			exit()
		}
	}
}

func main() {
	var n int
	var contentFn func() []string
	var oneShotMode bool
	var wrapSz int
	var err error

	flag.IntVar(&n, "n", 50, "The number of random words which constitute the test.")
	flag.IntVar(&wrapSz, "w", 80, "Wraps the input text at the given number of columns (ignored if -raw is present)")

	flag.BoolVar(&csvMode, "csv", false, "Print the test results to stdout in the form <wpm>,<cpm>,<accuracy>.")
	flag.BoolVar(&rawMode, "raw", false, "Don't reflow text or show one paragraph at a time.")
	flag.BoolVar(&oneShotMode, "o", false, "Automatically exit after a single run.")

	flag.Usage = func() {
		fmt.Println(`Usage: tt [options]

  By default tt creates a test consisting of 50 random words. Arbitrary text
  can also be piped directly into the program to create a custom test. Each
  paragraph of the input is treated as a segment of the test.
  
  E.G
  
  shuf -n 40 /etc/dictionaries-common/words|tt
  
  Note that linebreaks are determined exclusively by the input if -raw is specified.
  
Keybindings:
  <esc> Restarts the test
  <C-c> Terminates tt
  <C-backspace> Deletes the previous word
  
Options:`)

		flag.PrintDefaults()
	}
	flag.Parse()

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		if rawMode {
			contentFn = func() []string { return []string{string(b)} }
		} else {
			s := strings.Replace(string(b), "\r", "", -1)
			s = regexp.MustCompile("\n\n+").ReplaceAllString(s, "\n\n")
			content := strings.Split(strings.Trim(s, "\n"), "\n\n")

			for i, _ := range content {
				content[i] = strings.Replace(wordWrap(strings.Trim(content[i], " "), wrapSz), "\n", " \n", -1)
			}

			contentFn = func() []string { return content }
		}
	} else {
		contentFn = func() []string { return []string{randomText(n, wrapSz)} }
	}

	scr, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := scr.Init(); err != nil {
		panic(err)
	}

	fgcol := newTcellColor("#8C8C8C")
	bgcol := newTcellColor("#282828")

	hicol2 := newTcellColor("#805b13")
	hicol3 := newTcellColor("#b4801b")
	hicol := newTcellColor("#ffffff")
	errcol := newTcellColor("#a10705")

	cfg := readConfig()
	if c, ok := cfg["bgcol"]; ok {
		bgcol = newTcellColor(c)
	}
	if c, ok := cfg["fgcol"]; ok {
		fgcol = newTcellColor(c)
	}
	if c, ok := cfg["hicol"]; ok {
		hicol = newTcellColor(c)
	}
	if c, ok := cfg["hicol2"]; ok {
		hicol2 = newTcellColor(c)
	}
	if c, ok := cfg["hicol3"]; ok {
		hicol3 = newTcellColor(c)
	}
	if c, ok := cfg["errcol"]; ok {
		errcol = newTcellColor(c)
	}

	typer := NewTyper(scr, fgcol, bgcol, hicol, hicol2, hicol3, errcol)

	for {
		scr.Clear()
		nerrs, ncorrect, t, exitKey := typer.Start(contentFn())

		switch exitKey {
		case 0:
			cpm := int(float64(ncorrect) / (float64(t) / 60E9))
			wpm := cpm / 5
			accuracy := float64(ncorrect) / float64(nerrs+ncorrect) * 100

			results = append(results, result{wpm, cpm, accuracy})
			if oneShotMode {
				exit()
			}
			showReport(scr, cpm, wpm, accuracy)
		case tcell.KeyCtrlC:
			exit()
		}
	}
}
