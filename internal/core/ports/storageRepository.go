package ports

// StorageRepository interface
type StorageRepository interface {
	DirExists(path string) (bool, error)
}
