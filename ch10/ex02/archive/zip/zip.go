package zip

import (
	"archive/zip"

	"github.com/sakemi/go-training/ch10/ex02/archive"
)

type Reader struct {
	reader *zip.Reader
}

func NewReader(name string) (*archive.Reader, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	reader := Reader{r}
	return reader, nil
}

func init() {
	archive.RegisterFormat("zip", NewReader)
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
	r.Close()
}
