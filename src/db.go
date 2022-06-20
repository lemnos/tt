package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var FILE_STATE_DB string
var MISTAKE_DB string

func init() {
	if home, ok := os.LookupEnv("HOME"); !ok {
		die("Could not resolve home directory.")
	} else {
		FILE_STATE_DB = filepath.Join(home,  ".local/", "share/", "tt/", ".db")
		MISTAKE_DB = filepath.Join(home, ".local/", "share/", "tt/", ".errors")
	}
}

func readValue(path string, o interface{}) error {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, o)
}

func writeValue(path string, o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, b, 0600)
	if err != nil {
		panic(err)
	}
}
