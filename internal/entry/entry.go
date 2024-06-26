package entry

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Entry struct {
	Path string
	Info fs.FileInfo
}

func New(p string) (Entry, error) {
	path, err := filepath.Abs(p)

	if err != nil {
		return Entry{}, err
	}

	info, err := os.Lstat(path)

	if err != nil {
		return Entry{}, err
	}

	return Entry{Path: path, Info: info}, nil

}
