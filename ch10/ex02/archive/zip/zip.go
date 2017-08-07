package zip

import (
	"archive/zip"
	"io"
	"os"

	"github.com/sakemi/go-training/ch10/ex02/archive"
)

type reader struct {
	r *zip.Reader
	i int
}

func newReader(f *os.File) (archive.Reader, error) {
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	r, err := zip.NewReader(f, info.Size())
	if err != nil {
		return err
	}
	return &reader{r}, nil
}

func init() {
	archive.RegisterFormat("zip", newReader)
}

func (r *reader) Next() error {
	r.i++
	if len(r.r.File) <= i {
		return io.EOF
	}
	return nil
}

func (r *reader) Read(b []byte) (int, error) {
	f := r.r.File[r.i]
	rc, err := f.Open()
	if err != nil {
		return 0, err
	}
	defer rc.Close()
	return rc.Read(b)
}
