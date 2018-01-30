package main

import (
	"encoding/json"
	"os"
)

// ReadJSONFile fills object with data from given JSON-file.
func readJSONFile(name string, data interface{}) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}
