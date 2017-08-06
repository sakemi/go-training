package archive

import (
	"fmt"
	"strings"
)

var format map[string]func(string) (*ReadCloser, error)

func RegisterFormat(extention string, openReader func(string) (*ReadCloser, error)) {
	if format == nil {
		format = map[string]func(string) (*ReadCloser, error){}
	}
	format[extention] = openReader
}

func OpenReader(name string) (*ReadCloser, error) {
	n := strings.Split(name, ".")
	ext := n[len(n)-1]
	openReader, ok := format[ext]
	if !ok {
		return nil, fmt.Errorf("Unsupported file format: %s", ext)
	}
	return openReader(name)
}
