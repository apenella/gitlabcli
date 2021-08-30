package clone

import (
	"fmt"
	"os"

	"github.com/apenella/gitlabcli/internal/core/ports"
	cloneservice "github.com/apenella/gitlabcli/internal/core/services/clone"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/clone"
	gitrepo "github.com/apenella/gitlabcli/internal/repositories/git"
	storagerepo "github.com/apenella/gitlabcli/internal/repositories/storage"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/afero"
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

func RunEHandler(gitlab ports.GitlabCloneRepository, workingDir string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		var git gitrepo.GitRepository
		var service cloneservice.CloneService
		var storage storagerepo.ProjectStorage
		var h handler.CloneCliHandler

		errContext := "clone::RunEHandler"

		git, err = gitrepo.NewGitRepository()
		if err != nil {
			return errors.New(errContext, "Git repository could not be created", err)
		}

		storage = storagerepo.New(afero.NewOsFs())

		service, err = cloneservice.NewCloneService(
			gitlab,
			git,
			storage,
			cloneservice.WithUseNamespacePath(),
			cloneservice.WithBasePath(workingDir),
		)

		h, err = handler.NewCloneCliHandler(service, os.Stdout)
		if err != nil {
			return errors.New(errContext, "Handler cli could not be created", err)
		}

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
