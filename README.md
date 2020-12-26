# What

A terminal based typing test.

![](demo.gif)

# Installation

## Linux

```
sudo curl https://raw.githubusercontent.com/lemnos/tt/master/binaries/tt-linux_amd64 -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
```

## OSX

```
sudo curl https://raw.githubusercontent.com/lemnos/tt/master/binaries/tt-osx_amd64 -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
```

## From source

```
git clone github.com/lemnos/tt
make && sudo make install
```

or

`go get github.com/lemnos/tt` if you have `$GOPATH/bin` in your `$PATH`.

Best served on a terminal with truecolor and cursor shape support (e.g kitty, iterm)

# Usage

By default 50 words from the top 200 words in the English language are used to
constitute the test. Custom text can be supplied by piping aribirary text to
the program. Each paragraph in the input is shown as a separate segment of the
text.

E.G

- `shuf -n 40 /etc/dictionaries-common/words|tt` produces a test consisting of 40 random words drawn from `/etc/dictionaries-common/words`.
- `curl http://api.quotable.io/random|jq -r .content|tt` produces a test consisting of a random quote.

- `tt -n <num>` produces a test consisting of *num* randomly drawn english words
- `tt -t 10` starts a test which timesout after 10 seconds.
- `tt -theme <theme>` Starts tt with the provided theme (see `-list themes` for a list of options)
- `tt -csv` outputs the csv formatted results to STDOUT

The default behaviour is equivalent to 'tt -n 50'

See `-help` for an exhaustive and up to date list.

## Keys

- Pressing `escape` at any point restarts the test.
- `C-c` exits the test.

# Configuration

The theme can be configured by setting the following options in `~/.ttrc`:

 - `bgcol`:  The default background colour.
 - `fgcol`:  The default text colour.
 - `hicol`:  The colour used to highlight typed text.
 - `hicol2`  The colour used to highlight the current word.
 - `hicol3`: The colour used to highlight the next word.
 - `errcol`: The colour used to highlight errors.
 - `theme`: The theme from which default colors are drawn, a list of builtin themes can be obtained via `-list themes`.
