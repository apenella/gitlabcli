package git

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/go-git/go-git/v5"
)

type OptionsFunc func(*GitRepository)

type GitRepository struct {
	//Auth
}

func NewGitRepository(opts ...OptionsFunc) (GitRepository, error) {
	g := &GitRepository{}

	for _, opt := range opts {
		opt(g)
	}

	return *g, nil
}

func (g GitRepository) Clone(directory, url string) error {
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
