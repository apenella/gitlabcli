package ports

import "os"

type FilesystemRepository interface {
	Mkdir(string, os.FileMode) error
	MkdirAll(string, os.FileMode) error
}
