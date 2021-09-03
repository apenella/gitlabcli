package listproject

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/core/ports"
	listservice "github.com/apenella/gitlabcli/internal/core/services/list"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/list"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

var groupName string

func NewCommand() *command.AppCommand {

	getProjectsCmd := &cobra.Command{
		Use:     "projects",
		Aliases: []string{"project"},
		Short:   "Get project information from Gitlab",
		Long:    `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("unregitered handler for RunE on 'project' subcommand")
		},
	}

	getProjectsCmd.Flags().StringVarP(&groupName, "group", "g", "", "Gitlab group name to be consulted")

	return command.NewCommand(getProjectsCmd)
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
