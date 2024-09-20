package clone

import (
	"fmt"
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	service "github.com/apenella/gitlabcli/internal/core/services/clone"
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

// PerPage number of items to return on each Gitlab API page
const PerPage = 100

var cloneAll bool
var groupName string
var workingDir string

// NewCommand creates a clone command
func NewCommand(conf *configuration.Configuration) *command.AppCommand {
	cloneCmd := &cobra.Command{
		Use:   "clone [project_name]",
		Short: "Clone repositories from Gitlab to localhost",
		Long:  `Clone repositories from Gitlab to localhost`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return clone(conf, args)
		},
	}

	cloneCmd.Flags().BoolVar(&cloneAll, "all", false, "Clone all projects from all groups")
	cloneCmd.Flags().StringVarP(&groupName, "group", "g", "", "Gitlab group whose projects will be cloned")
	cloneCmd.Flags().StringVarP(&workingDir, "working-dir", "d", "", "Directory base path where projects are cloned to")

	return command.NewCommand(cloneCmd)
}

// clone functions is responsible to choose the clone strategy to be used: clone all projects, clone projects from a group or clone a list of projects
func clone(conf *configuration.Configuration, projects []string) error {
	var err error
	var gitlab gitlabrepo.GitlabRepository
	var git *gitrepo.Repository
	var cloneService service.Service
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

	git = gitrepo.NewRepository()

	if workingDir == "" {
		workingDir = conf.WorkingDir
	}

	cloneService, err = service.NewService(
		project,
		group,
		git,
		storage,
		service.WithUseNamespacePath(),
		service.WithBasePath(workingDir),
	)
	if err != nil {
		return errors.New(errContext, "Clone service could not be created", err)
	}

	h, err = handler.NewCloneCliHandler(cloneService, os.Stdout)
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

	err = h.CloneProject(projects...)
	if err != nil {
		return errors.New(errContext, "Clone could not be performed properly", err)
	}

	return nil
}
