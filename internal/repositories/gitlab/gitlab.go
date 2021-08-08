package gitlabrepo

import (
	errors "github.com/apenella/go-common-utils/error"
	gitlab "github.com/xanzy/go-gitlab"
)

const (
	PerPage int = 100
)

type GitlabRepository struct {
	client *gitlab.Client
}

func NewGitlabRepository(token, baseurl string) (GitlabRepository, error) {

	errContext := "gitlabrepo::NewGitlabRepository"

	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseurl))
	if err != nil {
		return GitlabRepository{}, errors.New(errContext, "Error creating gitlab client", err)
	}

	return GitlabRepository{
		client: client,
	}, nil
}
