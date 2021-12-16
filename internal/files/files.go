package files

import (
	"io/ioutil"
	"os"
)

// ReadFromFile reads from file
func ReadFromFile(path string) ([]byte, error) {
	var file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bytes []byte
	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
