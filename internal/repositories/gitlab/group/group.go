package group

import (
	"github.com/apenella/gitlabcli/internal/core/domain"
	errors "github.com/apenella/go-common-utils/error"
	gitlab "github.com/xanzy/go-gitlab"
)

type GitlabGroupLister interface {
	ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error)
	ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
}

type GitlabGroupRepository struct {
	group   GitlabGroupLister
	perPage int
}

func NewGitlabGroupRepository(group GitlabGroupLister, perPage int) *GitlabGroupRepository {
	return &GitlabGroupRepository{
		group:   group,
		perPage: perPage,
	}
}

func (g GitlabGroupRepository) Find(name string) ([]domain.Group, error) {
	groups := []domain.Group{}
	list := []*gitlab.Group{}

	errContext := "gitlabrepo::FindGroup"

	listGroupsOptions := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: g.perPage,
		},
		Search: &name,
	}

	list, err := g.list(listGroupsOptions, list)
	if err != nil {
		return groups, errors.New(errContext, "Error listing groups from gitlab repository", err)
	}

	for _, item := range list {
		g := domain.NewGroup(item.ID, item.Name)
		groups = append(groups, g)
	}

	return groups, nil
}

func (g GitlabGroupRepository) List() ([]domain.Group, error) {
	groups := []domain.Group{}
	list := []*gitlab.Group{}

	errContext := "gitlabrepo::ListGroups"

	listGroupsOptions := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: g.perPage,
		},
	}

	list, err := g.list(listGroupsOptions, list)
	if err != nil {
		return groups, errors.New(errContext, "Error listing groups from gitlab repository", err)
	}

	for _, item := range list {
		g := domain.NewGroup(item.ID, item.Name)
		groups = append(groups, g)
	}

	return groups, nil
}

func (g GitlabGroupRepository) list(options *gitlab.ListGroupsOptions, list []*gitlab.Group) ([]*gitlab.Group, error) {

	errContext := "gitlabrepo::listGroups"

	local_list, _, err := g.group.ListGroups(options)
	if err != nil {
		return nil, errors.New(errContext, "Gitlab client could not list groups", err)
	}

	list = append(list, local_list...)

	if len(local_list) < g.perPage {
		return list, nil
	} else {
		options.ListOptions.Page++
		return g.list(options, list)
	}
}

func (g GitlabGroupRepository) ListProjects(group string) ([]domain.Project, error) {
	projects := []domain.Project{}

	group_list, err := g.Find(group)
	if err != nil {
		return projects, err
	}

	for _, list := range group_list {
		listGroupProjectsOptions := &gitlab.ListGroupProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    1,
				PerPage: g.perPage,
			},
		}
		list, err := g.listProjects(list.Id, listGroupProjectsOptions, []*gitlab.Project{})
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

func (g GitlabGroupRepository) listProjects(id int, options *gitlab.ListGroupProjectsOptions, list []*gitlab.Project) ([]*gitlab.Project, error) {

	localList, _, err := g.group.ListGroupProjects(id, options)
	if err != nil {
		return nil, err
	}

	list = append(list, localList...)

	if len(localList) < g.perPage {
		return list, nil
	} else {
		options.ListOptions.Page++
		return g.listProjects(id, options, list)
	}
}
