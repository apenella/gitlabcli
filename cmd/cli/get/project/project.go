package getproject

import (
	"fmt"
	"os"

	"github.com/apenella/gitlabcli/internal/core/ports"
	getservice "github.com/apenella/gitlabcli/internal/core/services/get"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/get"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

var project string

func NewCommand() *command.AppCommand {

	getProjectsCmd := &cobra.Command{
		Use:   "project [<project_name>]+",
		Short: "Get project information from Gitlab",
		Long:  `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("unregitered handler for RunE on 'project' subcommand")
		},
		Args: cobra.MinimumNArgs(1),
	}

	getProjectsCmd.Flags().StringVarP(&project, "project", "p", "", "Gitlab project name to be consulted")

	return command.NewCommand(getProjectsCmd)
}

func RunEHandler(gitlab ports.GitlabProjectRepository) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		var service getservice.GetProjectService
		var h handler.GetProjectCliHandler

		errContext := "getproject::RunEHandler"

		service, err = getservice.NewGetProjectService(gitlab)
		if err != nil {
			return errors.New(errContext, "Gitlab service could not be created", err)
		}

		h, err = handler.NewGetProjectCliHandler(service, os.Stdout)
		if err != nil {
			return errors.New(errContext, "Handler cli could not be created", err)
		}

		for _, project := range args {
			err = h.GetProject(project)
			if err != nil {
				return errors.New(errContext, "Project detail could not be achieved", err)
			}
		}

		return nil
	}
}
