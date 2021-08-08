package clone

import (
	"fmt"

	handler "github.com/apenella/gitlabcli/internal/handlers"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

var cloneAll bool
var group string

//var dir, group string

func NewCommand() *command.AppCommand {
	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone repositories from Gitlab to localhost",
		Long:  `Clone repositories from Gitlab to localhost`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("unregitered handler for RunE on 'clone' subcommand")
		},
	}

	cloneCmd.Flags().BoolVar(&cloneAll, "all", false, "Clone all projects from all groups")
	cloneCmd.Flags().StringVarP(&group, "group", "g", "", "Group which its projects have to be cloned")
	//cloneCmd.Flags().StringVarP(&dir, "directory", "d", "", "Directory to clone the project")

	return command.NewCommand(cloneCmd)
}

func RunEHandler(h handler.CliHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		errContext := "clone::RunEHandler"

		if group != "" {
			err = h.CloneProjectFromGroup(group)
			if err != nil {
				return errors.New(errContext, fmt.Sprintf("Error cloning projects from group '%s'", group), err)
			}
		}

		if cloneAll {
			err = h.CloneAll()
			if err != nil {
				return errors.New(errContext, "Error cloning all projects", err)
			}
		}

		for _, project := range args {
			err = h.CloneProject(project)
			if err != nil {
				return errors.New(errContext, fmt.Sprintf("Error cloning '%s'", project), err)
			}
		}

		return nil
	}
}
