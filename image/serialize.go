package image

import (
	"encoding/gob"
	"os"
)

func register() {
	gob.Register(dockerLayer{})
}

// Save encodes a generic object via Gob to the given file path
func Save(path string, object interface{}) error {
	register()
	file, err := os.Create(path)
	defer file.Close()
	if err == nil {
		encoder := gob.NewEncoder(file)
		return encoder.Encode(object)
	}
	return err
}

// Load decodes via Gob the contents of the given file to an object
func Load(path string, object interface{}) error {
	register()
	file, err := os.Open(path)
	defer file.Close()
	if err == nil {
		decoder := gob.NewDecoder(file)
		return decoder.Decode(object)
	}
	return err
}
