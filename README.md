# What

A terminal based typing test.

![](demo.gif)

# Installation

## Linux

```
sudo curl -L https://github.com/lemnos/tt/releases/download/v0.4.2/tt-linux -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
sudo curl -o /usr/share/man/man1/tt.1.gz -L https://github.com/lemnos/tt/releases/download/v0.4.2/tt.1.gz
```

## OSX

```
mkdir -p /usr/local/bin /usr/local/share/man/man1 # Usually created by brew

sudo curl -L https://github.com/lemnos/tt/releases/download/v0.4.2/tt-osx -o /usr/local/bin/tt && sudo chmod +x /usr/local/bin/tt
sudo curl -o /usr/local/share/man/man1/tt.1.gz -L https://github.com/lemnos/tt/releases/download/v0.4.2/tt.1.gz
```

## Uninstall

```
sudo rm /usr/local/bin/tt /usr/share/man/man1/tt.1.gz
```

## From source

```
# debian dependencies
sudo apt install golang

# clone and make
git clone https://github.com/lemnos/tt
cd tt
make && sudo make install
```

Best served on a terminal with truecolor and cursor shape support (e.g kitty, iterm)

# Usage

By default 50 words from the top 1000 words in the English language are used to
constitute the test. Custom text can be supplied by piping arbitrary text to the
program. Each paragraph in the input is shown as a separate segment of the text.
See `man tt` or `man.md` for a complete description and a comprehensive set of
options.

## Keys

- Pressing `escape` at any point restarts the test.
- `C-c` exits the test.
- `right` moves to the next test.
- `left` moves to the previous test.

## Examples

 - `tt -quotes en` Starts quote mode with the builtin quote list 'en'.
 - `tt -n 10 -g 5` produces a test consisting of 50 randomly drawn words in 5 groups of 10 words each.
 - `tt -t 10` starts a timed test lasting 10 seconds.
 - `tt -theme gruvbox` Starts tt with the gruvbox theme.

`tt` is designed to be easily scriptable and integrate nicely with
other *nix tools. With a little shell scripting most features the user can
conceive of should be possible to implement. Below are some simple examples of
what can be achieved.

 - `shuf -n 40 /usr/share/dict/words|tt`  Produces a test consisting of 40 random words drawn from your system's dictionary.
 - `curl http://api.quotable.io/random|jq '[.text=.content|.attribution=.author]'|tt -quotes -` Produces a test consisting of a random quote.
 - `alias ttd='tt -csv >> ~/wpm.csv'` Creates an alias called ttd which keeps a log of progress in your home directory`.

The default behaviour is equivalent to `tt -n 50`.

See `-help` for an exhaustive list of options.

## Configuration

Custom themes and word lists can be defined in `~/.tt/themes` and `~/.tt/words`
and used in conjunction with the `-theme` and `-words` flags. A list of
preloaded themes and word lists can be found in `words/` and `themes/` and are
accessible by default using the respective flags.
