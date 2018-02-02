package util

import (
	"encoding/hex"
	"encoding/json"
	"github.com/satori/go.uuid"
	"os"
)

func ReadJSONFile(name string, data interface{}) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}

func EncodeHex(data string) string {
	return hex.EncodeToString([]byte(data))
}

func DecodeHex(data string) string {
	bin, err := hex.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(bin)
}

func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
