package archive

import (
	"fmt"
	"strings"
)

type Reader interface {
	Next() error
	Read(b []byte) (int, error)
	Close()
}

var format map[string]func(string) (*Reader, error)

func RegisterFormat(extention string, newReader func(string) (*Reader, error)) {
	if format == nil {
		format = map[string]func(string) (*Reader, error){}
	}
	format[extention] = newReader
}

func NewReader(name string) (*Reader, error) {
	n := strings.Split(name, ".")
	ext := n[len(n)-1]
	newReader, ok := format[ext]
	if !ok {
		return nil, fmt.Errorf("Unsupported file format: %s", ext)
	}
	return newReader(name)
}
