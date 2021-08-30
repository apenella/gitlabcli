package gitlabrepo

// import (
// 	"github.com/apenella/gitlabcli/internal/core/domain"
// 	errors "github.com/apenella/go-common-utils/error"
// 	gitlab "github.com/xanzy/go-gitlab"
// )

// func (g GitlabRepository) FindGroup(name string) ([]domain.Group, error) {
// 	groups := []domain.Group{}
// 	list := []*gitlab.Group{}

// 	errContext := "gitlabrepo::FindGroup"

// 	listGroupsOptions := &gitlab.ListGroupsOptions{
// 		ListOptions: gitlab.ListOptions{
// 			Page:    1,
// 			PerPage: PerPage,
// 		},
// 		Search: &name,
// 	}

// 	list, err := g.listGroups(listGroupsOptions, list)
// 	if err != nil {
// 		return groups, errors.New(errContext, "Error listing groups from gitlab repository", err)
// 	}

// 	for _, item := range list {
// 		g := domain.NewGroup(item.ID, item.Name)
// 		groups = append(groups, g)
// 	}

// 	return groups, nil
// }

// func (g GitlabRepository) ListGroups() ([]domain.Group, error) {
// 	groups := []domain.Group{}
// 	list := []*gitlab.Group{}

// 	errContext := "gitlabrepo::ListGroups"

// 	listGroupsOptions := &gitlab.ListGroupsOptions{
// 		ListOptions: gitlab.ListOptions{
// 			Page:    1,
// 			PerPage: PerPage,
// 		},
// 	}

// 	list, err := g.listGroups(listGroupsOptions, list)
// 	if err != nil {
// 		return groups, errors.New(errContext, "Error listing groups from gitlab repository", err)
// 	}

// 	for _, item := range list {
// 		g := domain.NewGroup(item.ID, item.Name)
// 		groups = append(groups, g)
// 	}

// 	return groups, nil
// }

// func (g GitlabRepository) listGroups(options *gitlab.ListGroupsOptions, list []*gitlab.Group) ([]*gitlab.Group, error) {

// 	errContext := "gitlabrepo::listGroups"

// 	local_list, _, err := g.Client.Groups.ListGroups(options)
// 	if err != nil {
// 		return nil, errors.New(errContext, "Gitlab client could not list groups", err)
// 	}

// 	list = append(list, local_list...)

// 	if len(local_list) < PerPage {
// 		return list, nil
// 	} else {
// 		options.ListOptions.Page++
// 		return g.listGroups(options, list)
// 	}
// }
