package gout

import (
	"encoding/json"
	"io/ioutil"
)

// File struct
type File struct {
	Path string
}

// ReadJSON method
func (f *File) ReadJSON(out interface{}) error {
	bytes, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, out)
}
