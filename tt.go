package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-isatty"
)

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

func showReport(scr tcell.Screen, cpm, wpm int, accuracy float64) {
	report := fmt.Sprintf("CPM: %d\nWPM: %d\nAccuracy: %.2f%%\n", cpm, wpm, accuracy)

	scr.Clear()
	drawCellsAtCenter(scr, stringToCells(report), -1)
	scr.HideCursor()
	scr.Show()

	for {
		if key, ok := scr.PollEvent().(*tcell.EventKey); ok && key.Key() == tcell.KeyEscape {
			return
		} else if ok && key.Key() == tcell.KeyCtrlC {
			scr.Fini()
			os.Exit(0)
		}
	}
}

func main() {
	var n int
	var csvMode bool

	flag.IntVar(&n, "n", 50, "The number of random words which constitute the test.")
	flag.BoolVar(&csvMode, "csv", false, "Print the test results to stdout in the form <cpm>,<wpm>,<accuracy>.")
	flag.Usage = func() {
		fmt.Println(`Usage: tt [options]

  By default tt creates a test consisting of 50 random words. Arbitrary text can also be piped directly into the program to create a custom test.
  
  E.G
  
  shuf -n 40 /etc/dictionaries-common/words|tr '\n' ' '|fold -s -w 80|tt
  
  Note that linebreaks are determined exclusively by the input.
  
Keybindings:
  <esc> Restarts the test
  <C-c> Terminates tt
  <C-backspace> Deletes the previous word
  
Options:`)

		flag.PrintDefaults()
	}
	flag.Parse()

	contentFn := func() string {
		return randomText(n)
	}

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		contentFn = func() string {
			return string(b)
		}
	}

	scr, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := scr.Init(); err != nil {
		panic(err)
	}

	defer scr.Fini()

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
		nerrs, ncorrect, t, completed := typer.Start([]string{contentFn()})
		if completed {

			cpm := int(float64(ncorrect) / (float64(t) / 60E9))
			wpm := cpm / 5
			accuracy := float64(ncorrect) / float64(nerrs+ncorrect) * 100

			if csvMode {
				scr.Fini()
				fmt.Printf("%d,%d,%.2f\n", cpm, wpm, accuracy)
				return
			}

			showReport(scr, cpm, wpm, accuracy)
		}
	}
}
