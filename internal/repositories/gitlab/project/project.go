package project

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	"github.com/xanzy/go-gitlab"
)

type GitlabProjectRepository struct {
	gitlabrepo.GitlabRepository
}

func (g GitlabProjectRepository) Find(name string) ([]domain.Project, error) {
	projects := []domain.Project{}

	list := []*gitlab.Project{}

	listProjectsOptions := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: g.PerPage,
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
			PerPage: g.PerPage,
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

	local_list, _, err := g.Client.Projects.ListProjects(options)
	if err != nil {
		return nil, err
	}

	list = append(list, local_list...)

	if len(local_list) < g.PerPage {
		return list, nil
	} else {
		options.ListOptions.Page++
		return g.list(options, list)
	}
}

func (g GitlabProjectRepository) ListFromGroup(group string) ([]domain.Project, error) {
	projects := []domain.Project{}

	group_list, err := g.Find(group)
	if err != nil {
		return projects, err
	}

	for _, list := range group_list {
		listGroupProjectsOptions := &gitlab.ListGroupProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    1,
				PerPage: g.PerPage,
			},
		}
		list, err := g.listFromGroup(list.Id, listGroupProjectsOptions, []*gitlab.Project{})
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
	}

	return projects, nil
}

func (g GitlabProjectRepository) listFromGroup(id int, options *gitlab.ListGroupProjectsOptions, list []*gitlab.Project) ([]*gitlab.Project, error) {

	local_list, _, err := g.Client.Groups.ListGroupProjects(id, options)
	if err != nil {
		return nil, err
	}

	list = append(list, local_list...)

	if len(local_list) < g.PerPage {
		return list, nil
	} else {
		options.ListOptions.Page++
		return g.listFromGroup(id, options, list)
	}
}
