package project

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	"github.com/xanzy/go-gitlab"
)

type GitlabProjectLister interface {
	ListProjects(opt *gitlab.ListProjectsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
}

type GitlabProjectRepository struct {
	perPage int
	project GitlabProjectLister
}

func NewGitlabProjectRepository(project GitlabProjectLister, perPage int) *GitlabProjectRepository {
	return &GitlabProjectRepository{
		perPage: perPage,
		project: project,
	}
}

func (g GitlabProjectRepository) Find(name string) ([]domain.Project, error) {
	projects := []domain.Project{}

	list := []*gitlab.Project{}

	listProjectsOptions := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: g.perPage,
		},
		Search: &name,
	}

	list, err := g.list(listProjectsOptions, list)
	if err != nil {
		return projects, err
	}

	for _, item := range list {
		p := domain.NewProject(
			item.ID,
			item.Name,
			item.DefaultBranch,
			item.PathWithNamespace,
			item.SSHURLToRepo,
			item.HTTPURLToRepo)

		projects = append(projects, p)
	}

	return projects, nil
}

func (g GitlabProjectRepository) List() ([]domain.Project, error) {
	projects := []domain.Project{}

	listProjectsOptions := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: g.perPage,
		},
	}

	list, err := g.list(listProjectsOptions, []*gitlab.Project{})
	if err != nil {
		return projects, err
	}

	for _, item := range list {
		p := domain.NewProject(
			item.ID,
			item.Name,
			item.DefaultBranch,
			item.PathWithNamespace,
			item.SSHURLToRepo,
			item.HTTPURLToRepo)
		projects = append(projects, p)
	}

	return projects, nil
}

func (g GitlabProjectRepository) list(options *gitlab.ListProjectsOptions, list []*gitlab.Project) ([]*gitlab.Project, error) {

	localList, _, err := g.project.ListProjects(options)
	if err != nil {
		return nil, err
	}

	list = append(list, localList...)

	if len(localList) < g.perPage {
		return list, nil
	} else {
		options.ListOptions.Page++
		return g.list(options, list)
	}
}
