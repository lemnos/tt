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
	"time"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-isatty"
)

var scr tcell.Screen
var csvMode bool

type result struct {
	wpm       int
	cpm       int
	accuracy  float64
	timestamp int64
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

func exit(rc int) {
	scr.Fini()

	if csvMode {
		for _, r := range results {
			fmt.Printf("%d,%d,%.2f,%d\n", r.wpm, r.cpm, r.accuracy, r.timestamp)
		}
	}

	os.Exit(rc)
}

func showReport(scr tcell.Screen, cpm, wpm int, accuracy float64) {
	report := fmt.Sprintf("WPM: %d\nCPM: %d\nAccuracy: %.2f%%", wpm, cpm, accuracy)

	scr.Clear()
	drawStringAtCenter(scr, report, tcell.StyleDefault)
	scr.HideCursor()
	scr.Show()

	for {
		if key, ok := scr.PollEvent().(*tcell.EventKey); ok && key.Key() == tcell.KeyEscape {
			return
		} else if ok && key.Key() == tcell.KeyCtrlC {
			exit(1)
		}
	}
}

func main() {
	var n int
	var ngroups int
	var contentFn func(sw, sh int) []string
	var rawMode bool
	var oneShotMode bool
	var maxLineLen int
	var noSkip bool
	var noReport bool
	var timeout int
	var listFlag string
	var err error
	var themeName string
	var showWpm bool

	flag.IntVar(&n, "n", 50, "The number of random words which constitute the test.")
	flag.IntVar(&ngroups, "g", 1, "The number of groups into which the generated test is split.")
	flag.IntVar(&maxLineLen, "w", 80, "The maximum line length in characters. (ignored if -raw is present).")
	flag.IntVar(&timeout, "t", -1, "Terminate the test after the given number of seconds.")

	flag.BoolVar(&showWpm, "showwpm", false, "Display WPM whilst typing.")
	flag.BoolVar(&noSkip, "noskip", false, "Disable word skipping when space is pressed.")
	flag.BoolVar(&oneShotMode, "oneshot", false, "Automatically exit after a single run (useful for scripts).")
	flag.BoolVar(&noReport, "noreport", false, "Don't show a report at the end of the test (useful in conjunction with -o).")
	flag.BoolVar(&csvMode, "csv", false, "Print the test results to stdout in the form wpm,cpm,accuracy,time.")
	flag.BoolVar(&rawMode, "raw", false, "Don't reflow text or show one paragraph at a time.")
	flag.StringVar(&themeName, "theme", "", "The theme to use (overrides ~/.ttrc).")
	flag.StringVar(&listFlag, "list", "", "-list themes prints a list of available themes.")

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

	if listFlag == "themes" {
		for t, _ := range themes {
			fmt.Println(t)
		}
		os.Exit(0)
	}

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		if rawMode {
			contentFn = func(sw, sh int) []string { return []string{string(b)} }
		} else {
			contentFn = func(sw, sh int) []string {
				wsz := maxLineLen
				if wsz > sw {
					wsz = sw - 8
				}

				s := strings.Replace(string(b), "\r", "", -1)
				s = regexp.MustCompile("\n\n+").ReplaceAllString(s, "\n\n")
				content := strings.Split(strings.Trim(s, "\n"), "\n\n")

				for i, _ := range content {
					content[i] = strings.Replace(wordWrap(strings.Trim(content[i], " "), wsz), "\n", " \n", -1)
				}

				return content
			}
		}
	} else {
		contentFn = func(sw, sh int) []string {
			wsz := maxLineLen
			if wsz > sw {
				wsz = sw - 8
			}

			if ngroups > n {
				ngroups = n
			}

			r := make([]string, ngroups)
			sz := n / ngroups
			for i := 0; i < ngroups-1; i++ {
				r[i] = randomText(sz, wsz)
			}

			r[ngroups-1] = randomText(sz+n%ngroups, wsz)

			return r
		}
	}

	cfg := readConfig()

	var bgcol, fgcol, hicol, hicol2, hicol3, errcol tcell.Color

	//If theme is explicitly specified as a flag
	if themeName != "" {
		if theme, ok := themes[themeName]; !ok {
			fmt.Fprintf(os.Stderr, "ERROR: %s is not a valid theme (see -list themes for a list of valid options).\n", themeName)
			os.Exit(1)
		} else {
			bgcol = newTcellColor(theme["bgcol"])
			fgcol = newTcellColor(theme["fgcol"])
			hicol = newTcellColor(theme["hicol"])
			hicol2 = newTcellColor(theme["hicol2"])
			hicol3 = newTcellColor(theme["hicol3"])
			errcol = newTcellColor(theme["errcol"])
		}
	} else {
		//Use the theme as a base
		theme := themes["default"]
		if c, ok := cfg["theme"]; ok {
			if v, ok := themes[c]; ok {
				theme = v
			}
		}

		bgcol = newTcellColor(theme["bgcol"])
		fgcol = newTcellColor(theme["fgcol"])
		hicol = newTcellColor(theme["hicol"])
		hicol2 = newTcellColor(theme["hicol2"])
		hicol3 = newTcellColor(theme["hicol3"])
		errcol = newTcellColor(theme["errcol"])

		//Allow individual colours to be overriden
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
	}

	scr, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := scr.Init(); err != nil {
		panic(err)
	}

	typer := NewTyper(scr, fgcol, bgcol, hicol, hicol2, hicol3, errcol)
	typer.SkipWord = !noSkip
	typer.ShowWpm = showWpm

	if timeout != -1 {
		timeout *= 1E9
	}

	for {
		sw, sh := scr.Size()
		nerrs, ncorrect, t, rc := typer.Start(contentFn(sw, sh), time.Duration(timeout))

		switch rc {
		case TyperComplete:
			cpm := int(float64(ncorrect) / (float64(t) / 60E9))
			wpm := cpm / 5
			accuracy := float64(ncorrect) / float64(nerrs+ncorrect) * 100

			results = append(results, result{wpm, cpm, accuracy, time.Now().Unix()})
			if !noReport {
				showReport(scr, cpm, wpm, accuracy)
			}
			if oneShotMode {
				exit(0)
			}
		case TyperSigInt:
			exit(1)

		case TyperResize:
			//Resize events restart the test, this shouldn't be a problem in the vast majority of cases
			//and allows us to avoid baking rewrapping logic into the typer.

			//TODO: implement state-preserving resize (maybe)
		}
	}
}
