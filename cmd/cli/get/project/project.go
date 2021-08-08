package getproject

import (
	"fmt"

	handler "github.com/apenella/gitlabcli/internal/handlers"
	"github.com/apenella/gitlabcli/pkg/command"
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

func RunEHandler(h handler.CliHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		for _, project := range args {
			err := h.GetProject(project)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
