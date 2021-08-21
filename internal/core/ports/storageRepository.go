package ports

type StorageRepository interface {
	DirExists(path string) (bool, error)
}
