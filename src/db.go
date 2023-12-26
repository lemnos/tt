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
	var ok bool
	var data string
	var home string

	if home, ok = os.LookupEnv("HOME"); !ok {
		if home, ok = os.LookupEnv("USERPROFILE"); !ok {
			die("Could not resolve home directory.")
		}
	}

	if data, ok = os.LookupEnv("XDG_DATA_HOME"); ok {
		data = filepath.Join(data, "/tt")
	} else {
		data = filepath.Join(home, "/.local/share/tt")
	}

	os.MkdirAll(data, 0700)

	FILE_STATE_DB = filepath.Join(data, ".db")
	MISTAKE_DB = filepath.Join(data, ".errors")
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
