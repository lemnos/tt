package main

import "regexp"

func generateWordTest(name string, n int, g int) func() []segment {
	var b []byte

	if b = readResource("words", name); b == nil {
		die("%s does not appear to be a valid word list. See '-list words' for a list of builtin word lists.", name)
	}

	words := regexp.MustCompile("\\s+").Split(string(b), -1)

	return func() []segment {
		segments := make([]segment, g)
		for i := 0; i < g; i++ {
			segments[i] = segment{randomText(n, words), ""}
		}

		return segments
	}
}
