package lib

import (
	"io/ioutil"
	"os"
)

// Key is used for safe reading of public/private keys.
//
// Tests: todo
type Key struct {
	data []byte
	filePath string
	password string
}

func NewKey(filePath string, password string) *Key {
	return &Key{
		data: nil,
		filePath: filePath,
		password: password,
	}
}

// Data returns key-data, it was read already,
// otherwise - it reads the data from the file and than returns it.
// Should be always covered by WipeOut().
//
// Tests: todo
func (k *Key) Data() ([]byte, error) {
	// It seems, that key is already read into memory.
	// No need to read it once more nad to trace it through OS caches.
	if k.data != nil && len(k.data) != 0 {
		return k.data, nil
	}

	// Seems that key was not read yet.
	keyFile, err := os.OpenFile(k.filePath, os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}

	// Todo: avoid os caches on reading, try using direct IO.
	data, err := ioutil.ReadAll(keyFile)
	if err != nil {
		return nil, err
	}

	k.data = data
	return k.data, nil
}

// WipeOut clears internal memory from sensitive data.
//
// Tests: todo
func (k* Key) WipeOut() {
	for i := 0; i < len(k.data); i++ {
		k.data[i] = 0
	}

	k.password = ""
}

// --------------------------------------------------------------------------------------------------------------------

func (e *EthereumClient) PrivateKey(path, password string) *Key {
	return NewKey(path, password)
}
