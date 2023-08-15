package dependencies

import (
	"io/fs"
	"os"

	"github.com/google/renameio/v2"
)

var OS = NewOsImpl()

// Os is an interface that wraps os package APIs
//
//go:generate mockery --name Os
type Os interface {
	// Lstat returns a FileInfo describing the named file.
	// If the file is a symbolic link, the returned FileInfo
	// describes the symbolic link. Lstat makes no attempt to follow the link.
	// If there is an error, it will be of type *PathError.
	Lstat(name string) (fs.FileInfo, error)
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	Stat(name string) (fs.FileInfo, error)
	// Remove removes the named file or (empty) directory.
	// If there is an error, it will be of type *PathError.
	Remove(name string) error
	// ReadFile reads the named file and returns the contents.
	// A successful call returns err == nil, not err == EOF.
	// Because ReadFile reads the whole file, it does not treat an EOF from Read
	// as an error to be reported.
	ReadFile(name string) ([]byte, error)
	// WriteFile mirrors ioutil.WriteFile, replacing an existing file with the same
	// name atomically.
	WriteFile(name string, data []byte, perm fs.FileMode, opts ...renameio.Option) error
}

type osImpl struct{}

func NewOsImpl() Os {
	return &osImpl{}
}

// Lstat returns a FileInfo describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
func (o *osImpl) Lstat(name string) (fs.FileInfo, error) {
	return os.Lstat(name)
}

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError
func (o *osImpl) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

// Remove removes the named file or (empty) directory.
// If there is an error, it will be of type *PathError.
func (o *osImpl) Remove(name string) error {
	return os.Remove(name)
}

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
func (o *osImpl) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// WriteFile mirrors ioutil.WriteFile, replacing an existing file with the same
// name atomically.
func (o *osImpl) WriteFile(name string, data []byte, perm fs.FileMode, opts ...renameio.Option) error {
	return renameio.WriteFile(name, data, perm, opts...)
}
