package gitlabrepo

// import (
// 	"github.com/apenella/gitlabcli/internal/core/domain"
// 	"github.com/xanzy/go-gitlab"
// )

// func (g GitlabRepository) FindProject(name string) ([]domain.Project, error) {
// 	projects := []domain.Project{}

// 	list := []*gitlab.Project{}

// 	listProjectsOptions := &gitlab.ListProjectsOptions{
// 		ListOptions: gitlab.ListOptions{
// 			Page:    1,
// 			PerPage: PerPage,
// 		},
// 		Search: &name,
// 	}

// 	list, err := g.listProjects(listProjectsOptions, list)
// 	if err != nil {
// 		return projects, err
// 	}

// 	for _, item := range list {
// 		p := domain.NewProject(
// 			item.ID,
// 			item.Name,
// 			item.DefaultBranch,
// 			item.PathWithNamespace,
// 			item.SSHURLToRepo,
// 			item.HTTPURLToRepo)

// 		projects = append(projects, p)
// 	}

// 	return projects, nil
// }

// func (g GitlabRepository) ListProjects() ([]domain.Project, error) {
// 	projects := []domain.Project{}

// 	listProjectsOptions := &gitlab.ListProjectsOptions{
// 		ListOptions: gitlab.ListOptions{
// 			Page:    1,
// 			PerPage: PerPage,
// 		},
// 	}

// 	list, err := g.listProjects(listProjectsOptions, []*gitlab.Project{})
// 	if err != nil {
// 		return projects, err
// 	}

// 	for _, item := range list {
// 		p := domain.NewProject(
// 			item.ID,
// 			item.Name,
// 			item.DefaultBranch,
// 			item.PathWithNamespace,
// 			item.SSHURLToRepo,
// 			item.HTTPURLToRepo)
// 		projects = append(projects, p)
// 	}

// 	return projects, nil
// }

// func (g GitlabRepository) listProjects(options *gitlab.ListProjectsOptions, list []*gitlab.Project) ([]*gitlab.Project, error) {

// 	local_list, _, err := g.Client.Projects.ListProjects(options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	list = append(list, local_list...)

// 	if len(local_list) < PerPage {
// 		return list, nil
// 	} else {
// 		options.ListOptions.Page++
// 		return g.listProjects(options, list)
// 	}
// }

// func (g GitlabRepository) ListProjectsFromGroup(group string) ([]domain.Project, error) {
// 	projects := []domain.Project{}

// 	group_list, err := g.FindGroup(group)
// 	if err != nil {
// 		return projects, err
// 	}

// 	for _, list := range group_list {
// 		listGroupProjectsOptions := &gitlab.ListGroupProjectsOptions{
// 			ListOptions: gitlab.ListOptions{
// 				Page:    1,
// 				PerPage: PerPage,
// 			},
// 		}
// 		list, err := g.listProjectsFromGroup(list.Id, listGroupProjectsOptions, []*gitlab.Project{})
// 		if err != nil {
// 			return projects, err
// 		}

// 		for _, item := range list {
// 			p := domain.NewProject(
// 				item.ID,
// 				item.Name,
// 				item.DefaultBranch,
// 				item.PathWithNamespace,
// 				item.SSHURLToRepo,
// 				item.HTTPURLToRepo)
// 			projects = append(projects, p)
// 		}
// 	}

// 	return projects, nil
// }

// func (g GitlabRepository) listProjectsFromGroup(id int, options *gitlab.ListGroupProjectsOptions, list []*gitlab.Project) ([]*gitlab.Project, error) {

// 	local_list, _, err := g.Client.Groups.ListGroupProjects(id, options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	list = append(list, local_list...)

// 	if len(local_list) < PerPage {
// 		return list, nil
// 	} else {
// 		options.ListOptions.Page++
// 		return g.listProjectsFromGroup(id, options, list)
// 	}
// }
