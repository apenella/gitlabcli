package listproject

import (
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/apenella/gitlabcli/internal/core/ports"
	listservice "github.com/apenella/gitlabcli/internal/core/services/list"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/list"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabgrouprepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/group"
	gitlabprojectrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/project"
	projectoutputrepo "github.com/apenella/gitlabcli/internal/repositories/output/project"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

const PerPage = 100

var groupName string

func NewCommand(conf *configuration.Configuration) *command.AppCommand {

	getProjectsCmd := &cobra.Command{
		Use:     "projects",
		Aliases: []string{"project"},
		Short:   "Get project information from Gitlab",
		Long:    `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listProjects(conf)
		},
	}

	getProjectsCmd.Flags().StringVarP(&groupName, "group", "g", "", "Gitlab group name to be consulted")

	return command.NewCommand(getProjectsCmd)
}

func listProjects(conf *configuration.Configuration) error {
	var gitlab gitlabrepo.GitlabRepository
	var err error
	var output ports.GitlabProjectOutputRepository
	var projectService listservice.ListProjectService
	var groupService listservice.ListGroupService
	var projectHandler handler.ListProjectCliHandler
	var groupHandler handler.ListGroupProjectCliHandler

	errContext := "listproject::listProjects"

	output = projectoutputrepo.NewProjectOutputRepository(os.Stdout)

	gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, PerPage)
	if err != nil {
		return errors.New(errContext, "Gitlab repository could not be created", err)
	}

	project := gitlabprojectrepo.NewGitlabProjectRepository(gitlab.Client.Projects, PerPage)

	group := gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage)

	projectService, err = listservice.NewListProjectService(project)
	if err != nil {
		return errors.New(errContext, "Gitlab group service could not be created", err)
	}

	projectHandler, err = handler.NewListProjectCliHandler(projectService, output)
	if err != nil {
		return errors.New(errContext, "Group handler cli could not be created", err)
	}

	groupService, err = listservice.NewListGroupService(group)
	if err != nil {
		return errors.New(errContext, "Gitlab group service could not be created", err)
	}

	groupHandler, err = handler.NewListGroupProjectCliHandler(groupService, output)
	if err != nil {
		return errors.New(errContext, "Group handler cli could not be created", err)
	}

	if groupName != "" {
		err = groupHandler.ListProjectsFromGroup(groupName)
	} else {
		err = projectHandler.ListProjects()
	}

	if err != nil {
		return errors.New(errContext, "Project could not be listed", err)
	}

	return nil
}

func RunEHandler(project ports.GitlabProjectRepository, group ports.GitlabGroupRepository, output ports.GitlabProjectOutputRepository) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		var projectService listservice.ListProjectService
		var groupService listservice.ListGroupService
		var projectHandler handler.ListProjectCliHandler
		var groupHandler handler.ListGroupProjectCliHandler

		errContext := "listproject::RunEHandler"

		projectService, err = listservice.NewListProjectService(project)
		if err != nil {
			return errors.New(errContext, "Gitlab group service could not be created", err)
		}

		projectHandler, err = handler.NewListProjectCliHandler(projectService, output)
		if err != nil {
			return errors.New(errContext, "Group handler cli could not be created", err)
		}

		groupService, err = listservice.NewListGroupService(group)
		if err != nil {
			return errors.New(errContext, "Gitlab group service could not be created", err)
		}

		groupHandler, err = handler.NewListGroupProjectCliHandler(groupService, output)
		if err != nil {
			return errors.New(errContext, "Group handler cli could not be created", err)
		}

		if groupName != "" {
			err = groupHandler.ListProjectsFromGroup(groupName)
		} else {
			err = projectHandler.ListProjects()
		}

		if err != nil {
			return errors.New(errContext, "Project could not be listed", err)
		}

		return nil
	}
}
