package main

import (
	"os"
	"path/filepath"
)

func generateTestFromFile(path string, startParagraph int) func() []segment {
	var paragraphs []string
	var db map[string]int
	var err error

	if path, err = filepath.Abs(path); err != nil {
		panic(err)
	}

	if err := readValue(FILE_STATE_DB, &db); err != nil {
		db = map[string]int{}
	}

	if startParagraph != -1 {
		db[path] = startParagraph
		writeValue(FILE_STATE_DB, db)
	}

	idx := db[path] - 1

	if b, err := os.ReadFile(path); err != nil {
		die("Failed to read %s.", path)
	} else {
		paragraphs = getParagraphs(string(b))
	}

	return func() []segment {
		idx++
		db[path] = idx
		writeValue(FILE_STATE_DB, db)

		if idx >= len(paragraphs) {
			return nil
		}

		return []segment{{paragraphs[idx], "", ""}}
	}
}
