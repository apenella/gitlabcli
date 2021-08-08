package listproject

import (
	"fmt"

	handler "github.com/apenella/gitlabcli/internal/handlers"
	"github.com/apenella/gitlabcli/pkg/command"
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

func RunEHandler(h handler.CliHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error

		if group != "" {
			err = h.ListProjectsFromGroup(group)
		} else {
			err = h.ListProjects()
		}

		if err != nil {
			return err
		}

		return nil
	}
}
