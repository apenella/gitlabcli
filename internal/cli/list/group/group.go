package listgroup

import (
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/apenella/gitlabcli/internal/core/ports"
	listservice "github.com/apenella/gitlabcli/internal/core/services/list"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/list"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabgrouprepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/group"
	groupoutputrepo "github.com/apenella/gitlabcli/internal/repositories/output/group"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

const PerPage = 100

func NewCommand(conf *configuration.Configuration) *command.AppCommand {

	listGroupCmd := &cobra.Command{
		Use:     "groups",
		Aliases: []string{"group"},
		Short:   "List gitlab groups",
		Long:    `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listGroups(conf)
		},
	}

	return command.NewCommand(listGroupCmd)
}

func listGroups(conf *configuration.Configuration) error {

	var service listservice.ListGroupService
	var h handler.ListGroupCliHandler
	var gitlab gitlabrepo.GitlabRepository
	var outputGroup ports.GitlabGroupOutputRepository
	var err error

	errContext := "listgroup::listGroups"

	outputGroup = groupoutputrepo.NewGroupOutputRepository(os.Stdout)

	gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, PerPage)
	if err != nil {
		return errors.New(errContext, "Gitlab repository could not be created", err)
	}

	gitlabgroup := gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage)

	service, err = listservice.NewListGroupService(gitlabgroup)
	if err != nil {
		return errors.New(errContext, "Gitlab service could not be created", err)
	}

	h, err = handler.NewListGroupCliHandler(service, outputGroup)
	if err != nil {
		return errors.New(errContext, "Handler cli could not be created", err)
	}

	err = h.ListGroups()
	if err != nil {
		return err
	}

	return nil

}
