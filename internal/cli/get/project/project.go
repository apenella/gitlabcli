package getproject

import (
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/apenella/gitlabcli/internal/core/ports"
	getservice "github.com/apenella/gitlabcli/internal/core/services/get"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/get"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabprojectrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/project"
	projectoutputrepo "github.com/apenella/gitlabcli/internal/repositories/output/project"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

var project string

const PerPage = 100

func NewCommand(conf *configuration.Configuration) *command.AppCommand {

	getProjectsCmd := &cobra.Command{
		Use:     "project [<project_name>]+",
		Aliases: []string{"projects", "prj", "p"},
		Short:   "Get project information from Gitlab",
		Long:    `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getProject(conf, args)
		},
		Args: cobra.MinimumNArgs(1),
	}

	getProjectsCmd.Flags().StringVarP(&project, "project", "p", "", "Gitlab project name to be consulted")

	return command.NewCommand(getProjectsCmd)
}

func getProject(conf *configuration.Configuration, projects []string) error {

	var gitlab gitlabrepo.GitlabRepository
	var err error
	var output ports.GitlabProjectOutputRepository
	var service getservice.GetProjectService
	var h handler.GetProjectCliHandler

	errContext := "getproject::getProject"

	output = projectoutputrepo.NewProjectOutputRepository(os.Stdout)

	gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, PerPage)
	if err != nil {
		return errors.New(errContext, "Gitlab repository could not be created", err)
	}

	project := gitlabprojectrepo.NewGitlabProjectRepository(gitlab.Client.Projects, PerPage)

	service, err = getservice.NewGetProjectService(project)
	if err != nil {
		return errors.New(errContext, "Gitlab service could not be created", err)
	}

	h, err = handler.NewGetProjectCliHandler(service, output)
	if err != nil {
		return errors.New(errContext, "Handler cli could not be created", err)
	}

	err = h.GetProject(projects...)
	if err != nil {
		return errors.New(errContext, "Project detail could not be achieved", err)
	}

	return nil
}
