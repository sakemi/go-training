package tar

import "archive/zip"

func OpenReader(name string) (*ReadCloser, error) {
	return zip.OpenReader(name)
}

func init() {

}
