package clone

import (
	"fmt"
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/apenella/gitlabcli/internal/core/ports"
	cloneservice "github.com/apenella/gitlabcli/internal/core/services/clone"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/clone"
	gitrepo "github.com/apenella/gitlabcli/internal/repositories/git"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabgrouprepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/group"
	gitlabprojectrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/project"
	storagerepo "github.com/apenella/gitlabcli/internal/repositories/storage"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const PerPage = 100

var cloneAll bool
var groupName string

//var dir, group string

func NewCommand(conf *configuration.Configuration) *command.AppCommand {
	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone repositories from Gitlab to localhost",
		Long:  `Clone repositories from Gitlab to localhost`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return clone(conf, args)
		},
	}

	cloneCmd.Flags().BoolVar(&cloneAll, "all", false, "Clone all projects from all groups")
	cloneCmd.Flags().StringVarP(&groupName, "group", "g", "", "Group which its projects have to be cloned")
	//cloneCmd.Flags().StringVarP(&dir, "directory", "d", "", "Directory to clone the project")

	return command.NewCommand(cloneCmd)
}

func clone(conf *configuration.Configuration, projects []string) error {
	var err error
	var gitlab gitlabrepo.GitlabRepository
	var git gitrepo.GitRepository
	var service cloneservice.CloneService
	var storage storagerepo.ProjectStorage
	var h handler.CloneCliHandler

	errContext := "clone::clone"

	gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, PerPage)
	if err != nil {
		return errors.New(errContext, "Gitlab repository could not be created", err)
	}

	project := gitlabprojectrepo.NewGitlabProjectRepository(gitlab.Client.Projects, PerPage)
	group := gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage)
	storage = storagerepo.New(afero.NewOsFs())

	git, err = gitrepo.NewGitRepository()
	if err != nil {
		return errors.New(errContext, "Git repository could not be created", err)
	}

	service, err = cloneservice.NewCloneService(
		project,
		group,
		git,
		storage,
		cloneservice.WithUseNamespacePath(),
		cloneservice.WithBasePath(conf.WorkingDir),
	)
	if err != nil {
		return errors.New(errContext, "Clone service could not be created", err)
	}

	h, err = handler.NewCloneCliHandler(service, os.Stdout)
	if err != nil {
		return errors.New(errContext, "Handler cli could not be created", err)
	}

	if groupName != "" {
		err = h.CloneProjectFromGroup(groupName)
		if err != nil {
			return errors.New(errContext, fmt.Sprintf("Error cloning projects from group '%s'", groupName), err)
		}
	}

	if cloneAll {
		err = h.CloneAll()
		if err != nil {
			return errors.New(errContext, "Error cloning all projects", err)
		}
	}

	for _, project := range projects {
		err = h.CloneProject(project)
		if err != nil {
			return errors.New(errContext, fmt.Sprintf("Error cloning '%s'", project), err)
		}
	}

	return nil
}

func RunEHandler(project ports.GitlabProjectRepository, group ports.GitlabGroupRepository, workingDir string) func(cmd *cobra.Command, args []string) error {
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
			project,
			group,
			git,
			storage,
			cloneservice.WithUseNamespacePath(),
			cloneservice.WithBasePath(workingDir),
		)
		if err != nil {
			return errors.New(errContext, "Clone service could not be created", err)
		}

		h, err = handler.NewCloneCliHandler(service, os.Stdout)
		if err != nil {
			return errors.New(errContext, "Handler cli could not be created", err)
		}

		if groupName != "" {
			err = h.CloneProjectFromGroup(groupName)
			if err != nil {
				return errors.New(errContext, fmt.Sprintf("Error cloning projects from group '%s'", groupName), err)
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
