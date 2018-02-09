package util

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"os"
	"path/filepath"
)

func ReadJSONFile(name string, data interface{}) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}

func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func ExeDirJoin(elem ...string) string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	elem = append([]string{filepath.Dir(exe)}, elem...)
	return filepath.Join(elem...)
}
