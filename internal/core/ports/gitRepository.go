package ports

type GitRepository interface {
	Clone(directory, url string) error
}
