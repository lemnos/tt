# What

A terminal based typing test.

![](demo.gif)

# Installation

## Linux

```
sudo curl -L https://github.com/lemnos/tt/releases/download/0.0.1/tt-linux -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
```

## OSX

```
sudo curl -L https://github.com/lemnos/tt/releases/download/0.0.1/tt-osx -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
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

By default 50 words from the top 1000 words in the English language are used to
constitute the test. Custom text can be supplied by piping aribirary text to
the program. Each paragraph in the input is shown as a separate segment of the
text.

## Keys

- Pressing `escape` at any point restarts the test.
- `C-c` exits the test.

## Examples

 - `tt -n 10` produces a test consisting of 10 randomly drawn english words
 - `tt -n 50 -g 5` produces a test consisting of 50 randomly drawn words in 5 groups of 10 words each.
 - `tt -t 10` starts a timed test consisting of 50 words
 - `tt -theme gruvbox` Starts tt with the gruvbox theme

The default behaviour is equivalent to `tt -n 50`.

See `-help` for an exhaustive list of options.

## Configuration

The theme can be configured by setting the following options in `~/.ttrc`:

 - `bgcol`:  The default background colour.
 - `fgcol`:  The default text colour.
 - `hicol`:  The colour used to highlight typed text.
 - `hicol2`  The colour used to highlight the current word.
 - `hicol3`: The colour used to highlight the next word.
 - `errcol`: The colour used to highlight errors.
 - `theme`: The theme from which default colors are drawn, a list of builtin themes can be obtained via `-list themes`.

## Recipes

`tt` is designed to be easily scriptable and integrate nicely with other with
other *nix tools. With a little shell scripting most features the user can
conceive of should be possible to implement. Below are some simple examples of
what can be achieved.

 - `shuf -n 40 /usr/share/dict/words|tt` Produces a test consisting of 40 random words drawn from your system's dictionary.
 - `curl http://api.quotable.io/random|jq -r .content|tt` Produces a test consisting of a random quote.
 - `alias ttd='tt -csv >> ~/wpm.csv'` Creates an alias called ttd which keeps a log of your progress in your home directory`.

