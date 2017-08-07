package tar

import (
	"archive/tar"
	"os"

	"github.com/sakemi/go-training/ch10/ex02/archive"
)

type reader struct {
	r *tar.Reader
}

func newReader(f *os.File) (archive.Reader, error) {
	r := tar.NewReader(f)
	return &reader{r}, nil
}

func init() {
	archive.RegisterFormat("tar", newReader)
}

func (r *reader) Next() error {
	_, err := r.r.Next()
	if err != nil {
		return err
	}
	return nil
}

func (r *reader) Read(b []byte) (int, error) {
	return r.r.Read(b)
}
