package storage

import (
	"fmt"
	"os"

	errors "github.com/apenella/go-common-utils/error"
)

type Storager interface {
	Stat(path string) (os.FileInfo, error)
}

type ProjectStorage struct {
	storage Storager
}

func New(fs Storager) ProjectStorage {
	return ProjectStorage{
		storage: fs,
	}
}

func (s ProjectStorage) DirExists(path string) (bool, error) {
	errContext := "storage::DirExists"

	stat, err := s.storage.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, errors.New(errContext, fmt.Sprintf("Stat for '%s' could not be achieved", path), err)
		}
	}

	return stat.IsDir(), nil
}
