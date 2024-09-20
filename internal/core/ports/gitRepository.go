package ports

// GitRepository interface
type GitRepository interface {
	Clone(directory, url string) error
}
