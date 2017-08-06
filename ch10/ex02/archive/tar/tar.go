package tar

import (
	"archive/tar"
	"os"

	"github.com/sakemi/go-training/ch10/ex02/archive"
)

type Reader struct {
	reader *tar.Reader
	file   *io.File
}

func NewReader(name string) (*archive.Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	reader := Reader{tar.NewReader(file)}
	return reader, nil
}

func init() {
	archive.RegisterFormat("tar", NewReader)
}

func (r *Reader) Next() error {
	//TODO
	return nil
}

func (r *Reader) Read(b []byte) (int, error) {
	//TODO
	return 0, nil
}

func (r *Reader) Close() {
	r.file.Close()
}
