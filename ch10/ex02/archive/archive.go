package archive

import (
	"fmt"
	"os"
	"strings"
)

type Reader interface {
	Next() error
	Read(b []byte) (int, error)
}

var format map[string]func(*os.File) (Reader, error)

func RegisterFormat(extention string, newReader func(*os.File) (Reader, error)) {
	if format == nil {
		format = map[string]func(string) (Reader, error){}
	}
	format[extention] = newReader
}

func NewReader(f *os.File) (Reader, error) {
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	n := strings.Split(info.Name(), ".")
	ext := n[len(n)-1]
	newReader, ok := format[ext]
	if !ok {
		return nil, fmt.Errorf("Unsupported file format: %s", ext)
	}
	return newReader(f)
}
