package git

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/go-git/go-git/v5"
)

// OptionsFunc to set options to GitRepository
type OptionsFunc func(*Repository)

// Repository is a struct to manage git operations
type Repository struct {
	//Auth
}

// NewRepository returns a new Repository
func NewRepository(opts ...OptionsFunc) *Repository {
	g := &Repository{}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

// Clone clones a git repository
func (g *Repository) Clone(directory, url string) error {
	var err error

	errContext := "git::Clone"

	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return errors.New(errContext, fmt.Sprintf("Project '%s' could not be cloned", url), err)
	}

	return nil
}
