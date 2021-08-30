package listproject

import (
	"fmt"
	"os"

	"github.com/apenella/gitlabcli/internal/core/ports"
	listservice "github.com/apenella/gitlabcli/internal/core/services/list"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/list"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

var group string

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

	getProjectsCmd.Flags().StringVarP(&group, "group", "g", "", "Gitlab group name to be consulted")

	return command.NewCommand(getProjectsCmd)
}

func RunEHandler(gitlab ports.GitlabProjectRepository) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		var service listservice.ListProjectService
		var h handler.ListProjectCliHandler

		errContext := "listproject::RunEHandler"

		service, err = listservice.NewListProjectService(gitlab)
		if err != nil {
			return errors.New(errContext, "Gitlab service could not be created", err)
		}

		h, err = handler.NewListProjectCliHandler(service, os.Stdout)
		if err != nil {
			return errors.New(errContext, "Handler cli could not be created", err)
		}

		if group != "" {
			err = h.ListProjectsFromGroup(group)
		} else {
			err = h.ListProjects()
		}

		if err != nil {
			return errors.New(errContext, "Project could not be listed", err)
		}

		return nil
	}
}
