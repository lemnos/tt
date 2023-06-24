package main

import (
	"bytes"
	"encoding/json"
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
var jsonMode bool

type result struct {
	Wpm       int       `json:"wpm"`
	Cpm       int       `json:"cpm"`
	Accuracy  float64   `json:"accuracy"`
	Timestamp int64     `json:"timestamp"`
	Mistakes  []mistake `json:"mistakes"`
}

func die(format string, args ...interface{}) {
	if scr != nil {
		scr.Fini()
	}
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

	if jsonMode {
		//Avoid null in serialized JSON.
		for i := range results {
			if results[i].Mistakes == nil {
				results[i].Mistakes = []mistake{}
			}
		}

		b, err := json.Marshal(results)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(b)
	}

	if csvMode {
		for _, r := range results {
			fmt.Printf("test,%d,%d,%.2f,%d\n", r.Wpm, r.Cpm, r.Accuracy, r.Timestamp)
			for _, m := range r.Mistakes {
				fmt.Printf("mistake,%s,%s\n", m.Word, m.Typed)
			}
		}
	}

	os.Exit(rc)
}

func showReport(scr tcell.Screen, cpm, wpm int, accuracy float64, attribution string, mistakes []mistake, titleStyle tcell.Style) {
	//consumes last key presses to avoid space press

	scr.Clear()
	scr.Show()
	hold := make(chan tcell.Event)
	go func(hold chan tcell.Event) {
		discard := make(chan<- tcell.Event)
		for {
			select {
			case discard <- scr.PollEvent():
				{

				}
			case <-time.After(300 * time.Millisecond):
				close(discard)
				close(hold)
				return

			}
		}
	}(hold)
	<-hold
	mistakeStr := ""
	if attribution != "" {
		attribution = "\n\nAttribution: " + attribution
	}

	if len(mistakes) > 0 {
		mistakeStr = "\nMistakes:    "
		for i, m := range mistakes {
			mistakeStr += m.Word
			if i != len(mistakes)-1 {
				mistakeStr += ", "
			}
		}
	}

	report := fmt.Sprintf("WPM:         %d\nCPM:         %d\nAccuracy:    %.2f%%%s%s", wpm, cpm, accuracy, mistakeStr, attribution)

	drawStringAtCenter(scr, report, tcell.StyleDefault)
	drawStringAsTitle(scr, "Press ESC, SPACE, or ENTER to continue.", titleStyle)
	scr.HideCursor()
	scr.Show()

	for {
		key, ok := scr.PollEvent().(*tcell.EventKey)

		if ok && (key.Key() == tcell.KeyEscape || key.Key() == tcell.KeyEnter || key.Rune() == 32) {
			return
		} else if ok && key.Key() == tcell.KeyCtrlC {
			exit(1)
		}
	}
}

func createDefaultTyper(scr tcell.Screen) *typer {
	return NewTyper(scr, true, tcell.ColorDefault,
		tcell.ColorDefault,
		tcell.ColorWhite,
		tcell.ColorGreen,
		tcell.ColorGreen,
		tcell.ColorMaroon)
}

func createTyper(scr tcell.Screen, bold bool, themeName string) *typer {
	var theme map[string]string

	if b := readResource("themes", themeName); b == nil {
		die("%s does not appear to be a valid theme, try '-list themes' for a list of built in thems.", themeName)
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

	return NewTyper(scr, bold, fgcol, bgcol, hicol, hicol2, hicol3, errcol)
}

var usage = `usage: tt [options] [file]

Modes
    -words  WORDFILE    Specifies the file from which words are randomly
                        drawn (default: 1000en).
    -quotes QUOTEFILE   Starts quote mode in which quotes are randomly drawn
                        from the given file. The file should be JSON encoded and
                        have the following form:

                        [{"text": "foo", attribution: "bar"}]

Word Mode
    -n GROUPSZ          Sets the number of words which constitute a group.
    -g NGROUPS          Sets the number of groups which constitute a test.

File Mode
    -start PARAGRAPH    The offset of the starting paragraph, set this to 0 to
                        reset progress on a given file.
Aesthetics
    -showwpm            Display WPM whilst typing.
    -theme THEMEFILE    The theme to use. 
    -w                  The maximum line length in characters. This option is 
    -notheme            Attempt to use the default terminal theme. 
                        This may produce odd results depending 
                        on the theme colours.
    -blockcursor        Use the default cursor style.
    -bold               Embolden typed text.
                        ignored if -raw is present.
Test Parameters
    -t SECONDS          Terminate the test after the given number of seconds.
    -noskip             Disable word skipping when space is pressed.
    -nobackspace        Disable the backspace key.
    -nohighlight        Disable current and next word highlighting.
    -highlight1         Only highlight the current word.
    -highlight2         Only highlight the next word.

Scripting
    -oneshot            Automatically exit after a single run.
    -noreport           Don't show a report at the end of a test.
    -csv                Print the test results to stdout in the form:
                        [type],[wpm],[cpm],[accuracy],[timestamp].
    -json               Print the test output in JSON.
    -raw                Don't reflow STDIN text or show one paragraph at a time.
                        Note that line breaks are determined exclusively by the
                        input.
    -multi              Treat each input paragraph as a self contained test.

Misc
    -list TYPE          Lists internal resources of the given type.
                        TYPE=[themes|quotes|words]

Version
    -v                  Print the current version.
`

func saveMistakes(mistakes []mistake) {
	var db []mistake

	if err := readValue(MISTAKE_DB, &db); err != nil {
		db = nil
	}

	db = append(db, mistakes...)
	writeValue(MISTAKE_DB, db)
}

func main() {
	var n int
	var g int

	var rawMode bool
	var oneShotMode bool
	var noHighlightCurrent bool
	var noHighlightNext bool
	var noHighlight bool
	var maxLineLen int
	var noSkip bool
	var noBackspace bool
	var noReport bool
	var noTheme bool
	var normalCursor bool
	var timeout int
	var startParagraph int

	var listFlag string
	var wordFile string
	var quoteFile string

	var themeName string
	var showWpm bool
	var multiMode bool
	var versionFlag bool
	var boldFlag bool

	var err error
	var testFn func() []segment

	flag.IntVar(&n, "n", 50, "")
	flag.IntVar(&g, "g", 1, "")
	flag.IntVar(&startParagraph, "start", -1, "")

	flag.IntVar(&maxLineLen, "w", 80, "")
	flag.IntVar(&timeout, "t", -1, "")

	flag.BoolVar(&versionFlag, "v", false, "")

	flag.StringVar(&wordFile, "words", "", "")
	flag.StringVar(&quoteFile, "quotes", "", "")

	flag.BoolVar(&showWpm, "showwpm", false, "")
	flag.BoolVar(&noSkip, "noskip", false, "")
	flag.BoolVar(&normalCursor, "blockcursor", false, "")
	flag.BoolVar(&noBackspace, "nobackspace", false, "")
	flag.BoolVar(&noTheme, "notheme", false, "")
	flag.BoolVar(&oneShotMode, "oneshot", false, "")
	flag.BoolVar(&noHighlight, "nohighlight", false, "")
	flag.BoolVar(&noHighlightCurrent, "highlight2", false, "")
	flag.BoolVar(&noHighlightNext, "highlight1", false, "")
	flag.BoolVar(&noReport, "noreport", false, "")
	flag.BoolVar(&boldFlag, "bold", false, "")
	flag.BoolVar(&csvMode, "csv", false, "")
	flag.BoolVar(&jsonMode, "json", false, "")
	flag.BoolVar(&rawMode, "raw", false, "")
	flag.BoolVar(&multiMode, "multi", false, "")
	flag.StringVar(&themeName, "theme", "default", "")
	flag.StringVar(&listFlag, "list", "", "")

	flag.Usage = func() { os.Stdout.Write([]byte(usage)) }
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
		fmt.Fprintf(os.Stderr, "tt version 0.4.2\n")
		os.Exit(1)
	}

	if noTheme {
		os.Setenv("TCELL_TRUECOLOR", "disable")
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

	switch {
	case wordFile != "":
		testFn = generateWordTest(wordFile, n, g)
	case quoteFile != "":
		testFn = generateQuoteTest(quoteFile)
	case !isatty.IsTerminal(os.Stdin.Fd()):
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		testFn = generateTestFromData(b, rawMode, multiMode)
	case len(flag.Args()) > 0:
		path := flag.Args()[0]
		testFn = generateTestFromFile(path, startParagraph)
	default:
		testFn = generateWordTest("1000en", n, g)
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

	var typer *typer
	if noTheme {
		typer = createDefaultTyper(scr)
	} else {
		typer = createTyper(scr, boldFlag, themeName)
	}

	if noHighlightNext || noHighlight {
		typer.currentWordStyle = typer.nextWordStyle
		typer.nextWordStyle = typer.defaultStyle
	}

	if noHighlightCurrent || noHighlight {
		typer.currentWordStyle = typer.defaultStyle
	}

	typer.SkipWord = !noSkip
	typer.DisableBackspace = noBackspace
	typer.BlockCursor = normalCursor
	typer.ShowWpm = showWpm

	if timeout != -1 {
		timeout *= 1e9
	}

	var tests [][]segment
	var idx = 0

	for {
		if idx >= len(tests) {
			tests = append(tests, testFn())
		}

		if tests[idx] == nil {
			exit(0)
		}

		if !rawMode {
			for i, _ := range tests[idx] {
				tests[idx][i].Text = reflow(tests[idx][i].Text)
			}
		}

		nerrs, ncorrect, t, rc, mistakes := typer.Start(tests[idx], time.Duration(timeout))
		saveMistakes(mistakes)

		switch rc {
		case TyperNext:
			idx++
		case TyperPrevious:
			if idx > 0 {
				idx--
			}
		case TyperComplete:
			cpm := int(float64(ncorrect) / (float64(t) / 60e9))
			wpm := cpm / 5
			accuracy := float64(ncorrect) / float64(nerrs+ncorrect) * 100

			results = append(results, result{wpm, cpm, accuracy, time.Now().Unix(), mistakes})
			if !noReport {
				attribution := ""
				if len(tests[idx]) == 1 {
					attribution = tests[idx][0].Attribution
				}
				showReport(scr, cpm, wpm, accuracy, attribution, mistakes, typer.nextWordStyle)
			}
			if oneShotMode {
				exit(0)
			}

			idx++
		case TyperSigInt:
			exit(1)

		case TyperResize:
			//Resize events restart the test, this shouldn't be a problem in the vast majority of cases
			//and allows us to avoid baking rewrapping logic into the typer.

			//TODO: implement state-preserving resize (maybe)
		}
	}
}
