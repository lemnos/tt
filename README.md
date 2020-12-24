# What

A terminal based typing test.

![](demo.gif)

# Installation

```
git clone github.com/lemnos/tt
make && sudo make install
```

or

`go get github.com/lemnos/tt` if you have $GOPATH/bin in your $PATH

Binaries are also available:

## OSX

```
sudo curl https://raw.githubusercontent.com/lemnos/tt/master/binaries/tt-osx_amd64 -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
```

## Linux

```
sudo curl https://raw.githubusercontent.com/lemnos/tt/master/binaries/tt-linux_amd64 -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
```

Best served on a terminal with truecolor and cursor shape support (e.g urxvt, kitty, iterm)

# Usage

- `tt -n <num>` produces a test consisting of *num* randomly drawn english words
- `tt -csv -n <num>` outputs the csv formatted results to STDOUT

Custom text can be supplied by piping aribirary text to the program.

E.G

- `shuf -n 40 /etc/dictionaries-common/words|tr '\n' ' '|fold -s -w 80|tt` produces a test consisting of 40 random words drawn from `/etc/dictionaries-common/words`.

Note that line breaks are exclusively determined by the input.

- Pressing `escape` at any point restarts the test.
- `C-c` exits the test.

# Configuration

The theme can be configured by setting the following options in `~/.ttrc`:

 - bgcol:  The default background colour.
 - fgcol:  The default text colour.
 - hicol:  The colour used to highlight typed text.
 - hicol2  The colour used to highlight the current word.
 - hicol3: The colour used to highlight the next word.
 - errcol: The colour used to highlight errors.
