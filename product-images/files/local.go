package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local is the implementation of the Storage interface which works
// with local disk on the current machine

type Local struct {
	basePath    string
	maxFileSize int
}

func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p, maxFileSize: maxSize}, nil
}

// Returns the absolute path
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}

func (l *Local) Save(path string, contents io.Reader) error {
	fp := l.fullPath(path)

	// Get the directory
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}
	// if the file exists, delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}
	// Create a new file at the path
	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer f.Close()

	// write contents to the new file
	// Ensure that we are not writing contents greater than max bytes

	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}

// Get the file at the given path and return a Reader
// The calling function is responsible for closing the reader
func (l *Local) Get(path string) (*os.File, error) {
	fp := l.fullPath(path)
	r, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return r, nil
}
