package file

import "os"
import "io/ioutil"

// Handler a simple type for file operations
type Handler struct{}

// Exists returns a bool saying if a file exists on disk or not
func (Handler) Exists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Read returns the bytes of a file or an error if the file cannot be read
func (Handler) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
