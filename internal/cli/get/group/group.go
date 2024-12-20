package getgroup

import (
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/apenella/gitlabcli/internal/core/ports"
	getservice "github.com/apenella/gitlabcli/internal/core/services/get"
	handler "github.com/apenella/gitlabcli/internal/handlers/cli/get"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	gitlabgrouprepo "github.com/apenella/gitlabcli/internal/repositories/gitlab/group"
	groupoutputrepo "github.com/apenella/gitlabcli/internal/repositories/output/group"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

// PerPage number of items to return on each Gitlab API page
const PerPage = 100

// NewCommand creates a get group command
func NewCommand(conf *configuration.Configuration) *command.AppCommand {

	getGroupCmd := &cobra.Command{
		Use:     "group [<group_name>]+",
		Aliases: []string{"groups", "grp", "g"},
		Short:   "Get group information from Gitlab",
		Long:    `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getGroup(conf, args)
		},
		Args: cobra.MinimumNArgs(1),
	}

	return command.NewCommand(getGroupCmd)
}

// getGroup function is responsible to get group information from Gitlab
func getGroup(conf *configuration.Configuration, groups []string) error {

	var err error
	var gitlab gitlabrepo.GitlabRepository
	var output ports.GitlabGroupOutputRepository
	var service getservice.GetGroupService
	var h handler.GetGroupCliHandler

	errContext := "getgroup::getGroup"

	output = groupoutputrepo.NewOutputRepository(os.Stdout)

	gitlab, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL, PerPage)
	if err != nil {
		return errors.New(errContext, "Gitlab repository could not be created", err)
	}

	group := gitlabgrouprepo.NewGitlabGroupRepository(gitlab.Client.Groups, PerPage)

	service, err = getservice.NewGetGroupService(group)
	if err != nil {
		return errors.New(errContext, "Gitlab service could not be created", err)
	}

	h, err = handler.NewGetGroupCliHandler(service, output)
	if err != nil {
		return errors.New(errContext, "Handler cli could not be created", err)
	}

	err = h.GetGroup(groups...)
	if err != nil {
		return errors.New(errContext, "Group detail could not be achieved", err)
	}

	return nil
}
