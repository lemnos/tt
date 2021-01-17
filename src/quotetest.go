package main

import (
	"encoding/json"
	"math/rand"
)

func generateQuoteTest(name string) func() []segment {
	var quotes []segment

	if b := readResource("quotes", name); b == nil {
		die("%s does not appear to be a valid quote file. See '-list quotes' for a list of builtin quotes.", name)
	} else {
		err := json.Unmarshal(b, &quotes)
		if err != nil {
			die("Improperly formatted quote file: %v", err)
		}
	}

	return func() []segment {
		idx := rand.Int() % len(quotes)
		return quotes[idx : idx+1]
	}
}
