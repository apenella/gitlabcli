package storage

import (
	"fmt"
	"os"

	errors "github.com/apenella/go-common-utils/error"
)

// Storager interface defines the methods to stat a file
type Storager interface {
	Stat(path string) (os.FileInfo, error)
}

// ProjectStorage struct to store locally Gitlab repository
type ProjectStorage struct {
	storage Storager
}

// New returns a new ProjectStorage
func New(fs Storager) ProjectStorage {
	return ProjectStorage{
		storage: fs,
	}
}

// DirExists checks if a directory exists
func (s ProjectStorage) DirExists(path string) (bool, error) {
	errContext := "storage::DirExists"

	stat, err := s.storage.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, errors.New(errContext, fmt.Sprintf("Stat for '%s' could not be achieved", path), err)
	}

	return stat.IsDir(), nil
}
