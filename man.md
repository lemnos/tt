% tt(1)

# NAME

tt - A terminal based typing test

# SYNOPSIS

usage: tt \[OPTION\]... \[FILE\]

# DESCRIPTION

  By default tt creates a test consisting of 50 randomly generated words from
  the top 1000 words in the English language. If provided with a path, tt will
  use the given file as input treating each paragraph as a separate segment of
  the test. The program will automatically keep track of your position in the
  file so subsequent invocations on the same path will place you at the most
  recent paragraph (-start 0 can be used to reset your position).  
  
  Arbitrary text can also be piped directly into the program to create a custom
  test. Each paragraph of the input is treated as a segment unless '-multi' is
  supplied in which case each paragraph is treated as a separate test. 

# OPTIONS

## Modes

-words  *WORDFILE*

: Specifies the file from which words are randomly generated (default: 1000en).

-quotes *QUOTEFILE*

: Starts quote mode in which quotes are randomly generated from the given file. The file should be JSON encoded and have the following form:

    [{"text": "foo", attribution: "bar"}]

## Word Mode

-n *GROUPSZ*

: Sets the number of words which constitute a group.

-g *NGROUPS*

: Sets the number of groups which constitute a test.

## File Mode
-start *PARAGRAPH*

: The offset of the starting paragraph, set this to 0 to reset progress on a given file.

## Aesthetics

-showwpm

: Display WPM whilst typing.

-theme *THEMEFILE*

: The theme to use. 

-w 

: The maximum line length in characters. This option is ignored if -raw is present.

## Test Parameters

-t *SECONDS*

: Terminate the test after the given number of seconds.

-noskip

: Disable word skipping when space is pressed.

-nohighlight

: Disable highlighting.

-highlight1

: Only highlight the current word.

-highlight2

: Only highlight the next word.

## Scripting

-oneshot

: Automatically exit after a single run.

-noreport

: Don't show a report at the end of a test.

-csv

: Print CSV formatted results.

	Tests have the form:

	```
	test,[wpm],[cpm],[accuracy],[timestamp].
	```

	Mistakes have the form:

	```
	mistake,[word],[typed]
	```

-json

: Print the test output in JSON.

-raw

: Don't reflow STDIN text or show one paragraph at a time.  Note that line breaks
are determined exclusively by the input.  

-multi 

: Treat each input paragraph as a self contained test.

## Misc

**-list** *TYPE*\

    Lists internal resources of the given type. TYPE=[themes|quotes|words].

**-v**\

    Print the current version.

# EXAMPLES

Creates a series of tests each consisting of a random quote drawn from the
builtin quote file 'en'.
```
tt -quotes en
```

Creates a series of tests each consisting of 10 random words drawn from
words.txt
```
tt -words words.txt -n 10
```

Starts a sequence of tests in which each test consists of a paragraph from war
and peace starting with paragraph 1.
```
tt ~/war_and_peace.txt -start 1
```

Produces a test consisting of 40 random words draw from 
the system dictionary (similar to 'tt -n 40').
```
shuf -n 40 /usr/share/dict/words|tt
```

Starts a test consisting of two randomly drawn quotes from api.quotable.io and
prints the output of each test to STDOUT in csv format.
```
curl https://api.quotable.io/quotes|\
    jq '[.results[]|.text=.content|.attribution=.author][:2]'|\
    tt -quotes - -norreport -csv
```

# PATHS

  Some options like **-words** and **-theme** accept a path. If the given path does
  not exist, the following directories are searched for a file with the given
  name before falling back to internal resources:

  ~/.tt/words\
  ~/.tt/themes\
  /etc/tt/words\
  /etc/tt/themes

# KEYS

  **esc: ** Restarts the test\
  **C-c: ** Terminates tt\
  **C-backspace: ** Deletes the previous word\
  **right** Move to the next test.\
  **left** Move to the previous test.

# AUTHOR

Aetnaeus (aetnaeus@protonmail.com)

# SEE ALSO

## Project Page

    https://github.com/lemnos/tt

# LICENSE

MIT
