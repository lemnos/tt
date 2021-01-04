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

func die(format string, args ...interface{}) {
	scr.Fini()
	fmt.Fprintf(os.Stderr, "ERROR: ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

var results []result

func parseConfig(b []byte) map[string]string {
	if b == nil {
		return nil
	}

	cfg := map[string]string{}
	for _, ln := range bytes.Split(b, []byte("\n")) {
		a := strings.SplitN(string(ln), ":", 2)
		if len(a) == 2 {
			cfg[a[0]] = strings.Trim(a[1], " ")
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

func createTyper(scr tcell.Screen, themeName string) *typer {
	var theme map[string]string

	if b := readResource("themes", themeName); b == nil {
		die("%s does not appear to be a valid theme, try '-list themes' for a list of built in themes.", themeName)
	} else {
		theme = parseConfig(b)
	}

	var bgcol, fgcol, hicol, hicol2, hicol3, errcol tcell.Color
	var err error

	if bgcol, err = newTcellColor(theme["bgcol"]); err != nil {
		die("bgcol is not defined and/or a valid hex colour.")
	}
	if fgcol, err = newTcellColor(theme["fgcol"]); err != nil {
		die("fgcol is not defined and/or a valid hex colour.")
	}
	if hicol, err = newTcellColor(theme["hicol"]); err != nil {
		die("hicol is not defined and/or a valid hex colour.")
	}
	if hicol2, err = newTcellColor(theme["hicol2"]); err != nil {
		die("hicol2 is not defined and/or a valid hex colour.")
	}
	if hicol3, err = newTcellColor(theme["hicol3"]); err != nil {
		die("hicol3 is not defined and/or a valid hex colour.")
	}
	if errcol, err = newTcellColor(theme["errcol"]); err != nil {
		die("errcol is not defined and/or a valid hex colour.")
	}

	return NewTyper(scr, fgcol, bgcol, hicol, hicol2, hicol3, errcol)
}

func main() {
	var n int
	var ngroups int
	var testFn func() []string
	var rawMode bool
	var oneShotMode bool
	var maxLineLen int
	var noSkip bool
	var noReport bool
	var timeout int
	var listFlag string
	var wordList string
	var err error
	var themeName string
	var showWpm bool
	var multiMode bool
	var versionFlag bool

	flag.IntVar(&n, "n", 50, "The number of words which constitute a group.")
	flag.IntVar(&ngroups, "g", 1, "The number of groups which constitute a generated test.")

	flag.IntVar(&maxLineLen, "w", 80, "The maximum line length in characters. (ignored if -raw is present).")
	flag.IntVar(&timeout, "t", -1, "Terminate the test after the given number of seconds.")

	flag.BoolVar(&versionFlag, "v", false, "Print the current version.")

	flag.StringVar(&wordList, "words", "1000en", "The name of the word list used to generate random text.")
	flag.BoolVar(&showWpm, "showwpm", false, "Display WPM whilst typing.")
	flag.BoolVar(&noSkip, "noskip", false, "Disable word skipping when space is pressed.")
	flag.BoolVar(&oneShotMode, "oneshot", false, "Automatically exit after a single run (useful for scripts).")
	flag.BoolVar(&noReport, "noreport", false, "Don't show a report at the end of the test (useful in conjunction with -o).")
	flag.BoolVar(&csvMode, "csv", false, "Print the test results to stdout in the form wpm,cpm,accuracy,time.")
	flag.BoolVar(&rawMode, "raw", false, "Don't reflow text or show one paragraph at a time. (note that linebreaks are determined exclusively by the input)")
	flag.BoolVar(&multiMode, "multi", false, "Treat each input paragraph as a self contained test.")
	flag.StringVar(&themeName, "theme", "default", "The theme to use.")
	flag.StringVar(&listFlag, "list", "", "Lists internal resources (e.g -list themes yields a list of builtin themes)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: tt [options]

  By default tt creates a test consisting of 50 randomly generated words from the
  top 1000 words in the English language. Arbitrary text can also be piped
  directly into the program to create a custom test. Each paragraph of the
  input is treated as a segment of the test. 
  
Examples:

  # Equivalent to 'tt -n 40 -words /usr/share/dict/words'
  shuf -n 40 /usr/share/dict/words|tt

  # Starts a test consisting of a random quote.
  curl https://api.quotable.io/random|jq -r .content|tt

  # Starts single a test consisting of multiple random quotes.
  curl https://api.quotable.io/quotes|jq -r .results[].content|sort -R|sed -e 's/$/\n/'|tt

  # Starts multiple tests each consisting of a random quote
  curl https://api.quotable.io/quotes|jq -r .results[].content|sort -R|sed -e 's/$/\n/'|tt -multi


Paths:

  Some options like '-words' and '-theme' accept a path. If the given path does
  not exist, the following directories are searched for a file with the given
  name before falling back to the internal resource (if one exists):
  
  -words (See -list words):

  ~/.tt/words/
  /etc/tt/words/

  -theme (See -list themes):

  ~/.tt/themes/
  /etc/tt/themes/
  
Keybindings:
  <esc> Restarts the test
  <C-c> Terminates tt
  <C-backspace> Deletes the previous word
  
Options:
`)

		flag.PrintDefaults()
	}
	flag.Parse()

	if listFlag != "" {
		prefix := listFlag + "/"
		for path, _ := range packedFiles {
			if strings.Index(path, prefix) == 0 {
				_, f := filepath.Split(path)
				fmt.Println(f)
			}
		}

		os.Exit(0)
	}

	if versionFlag {
		fmt.Fprintf(os.Stderr, "tt version 0.3.0\n")
		os.Exit(1)
	}

	reflow := func(s string) string {
		sw, _ := scr.Size()

		wsz := maxLineLen
		if wsz > sw {
			wsz = sw - 8
		}

		s = regexp.MustCompile("\\s+").ReplaceAllString(s, " ")
		return strings.Replace(
			wordWrap(strings.Trim(s, " "), wsz),
			"\n", " \n", -1)
	}

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		getParagraphs := func(s string) []string {
			s = strings.Replace(s, "\r", "", -1)
			s = regexp.MustCompile("\n\n+").ReplaceAllString(s, "\n\n")
			return strings.Split(strings.Trim(s, "\n"), "\n\n")
		}

		if rawMode {
			testFn = func() []string { return []string{string(b)} }
		} else if multiMode {
			paragraphs := getParagraphs(string(b))
			i := 0

			testFn = func() []string {
				if i < len(paragraphs) {
					p := paragraphs[i]
					i++
					return []string{p}
				} else {
					return nil
				}
			}
		} else {
			testFn = func() []string {
				return getParagraphs(string(b))
			}
		}
	} else {
		testFn = func() []string {
			var b []byte

			if b = readResource("words", wordList); b == nil {
				die("%s does not appear to be a valid word list. See '-list words' for a list of builtin word lists.", wordList)
			}

			words := regexp.MustCompile("\\s+").Split(string(b), -1)

			r := make([]string, ngroups)
			for i := 0; i < ngroups; i++ {
				r[i] = randomText(n, words)
			}

			return r
		}
	}

	scr, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := scr.Init(); err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			scr.Fini()
			panic(r)
		}
	}()

	typer := createTyper(scr, themeName)
	typer.SkipWord = !noSkip
	typer.ShowWpm = showWpm

	if timeout != -1 {
		timeout *= 1E9
	}

	var showNext = true
	var paragraphs []string

	for {
		if showNext {
			paragraphs = testFn()

			if paragraphs == nil {
				exit(0)
			}
		}

		if !rawMode {
			for i, _ := range paragraphs {
				paragraphs[i] = reflow(paragraphs[i])
			}
		}

		nerrs, ncorrect, t, rc := typer.Start(paragraphs, time.Duration(timeout))

		showNext = false
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
			showNext = true
		case TyperSigInt:
			exit(1)

		case TyperResize:
			//Resize events restart the test, this shouldn't be a problem in the vast majority of cases
			//and allows us to avoid baking rewrapping logic into the typer.

			//TODO: implement state-preserving resize (maybe)
		}
	}
}
