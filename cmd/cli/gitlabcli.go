package gitlabcli

import (
	"os"

	"github.com/apenella/gitlabcli/cmd/cli/clone"
	"github.com/apenella/gitlabcli/cmd/cli/get"
	getgroup "github.com/apenella/gitlabcli/cmd/cli/get/group"
	getproject "github.com/apenella/gitlabcli/cmd/cli/get/project"
	"github.com/apenella/gitlabcli/cmd/cli/list"
	listgroup "github.com/apenella/gitlabcli/cmd/cli/list/group"
	listproject "github.com/apenella/gitlabcli/cmd/cli/list/project"
	"github.com/apenella/gitlabcli/cmd/cli/version"
	"github.com/apenella/gitlabcli/cmd/configuration"
	"github.com/apenella/gitlabcli/internal/core/ports"
	service "github.com/apenella/gitlabcli/internal/core/services"
	handler "github.com/apenella/gitlabcli/internal/handlers"
	gitrepo "github.com/apenella/gitlabcli/internal/repositories/git"
	gitlabrepo "github.com/apenella/gitlabcli/internal/repositories/gitlab"
	storagerepo "github.com/apenella/gitlabcli/internal/repositories/storage"
	"github.com/apenella/gitlabcli/pkg/command"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var configFile string

func NewCommand() *command.AppCommand {
	var err error
	var conf configuration.Configuration
	var getGroupSubcommand,
		getProjectSubcommand,
		getSubcommand,
		versionSubcommand,
		listGroupsSubcommand,
		listProjectsSubcommand,
		listSubcommand,
		cloneSubcommand *command.AppCommand
	var gitRepo gitrepo.GitRepository
	var glRepo gitlabrepo.GitlabRepository
	var glSrv ports.Service
	var glStorage storagerepo.ProjectStorage
	var cliHandler handler.CliHandler

	gitlabCmd := &cobra.Command{
		Use:   "gitlabcli",
		Short: "gitlab cli",
		Long:  `Set of utils to manage Gitlab repositories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			conf, err = configuration.New(configFile)
			if err != nil {
				return err
			}
			err = conf.Validate()
			if err != nil {
				return err
			}

			gitRepo, err = gitrepo.NewGitRepository()
			if err != nil {
				return err
			}

			glRepo, err = gitlabrepo.NewGitlabRepository(conf.Token, conf.BaseURL)
			if err != nil {
				return err
			}

			glStorage = storagerepo.New(afero.NewOsFs())

			glSrv, err = service.New(glRepo, gitRepo, glStorage,
				service.WithUseNamespacePath(),
				service.WithBasePath(conf.WorkingDir),
			)

			if err != nil {
				return err
			}

			cliHandler, err = handler.NewCliHandler(glSrv, os.Stdout)
			if err != nil {
				return err
			}

			listGroupsSubcommand.Options(command.WithRunE(listgroup.RunEHandler(cliHandler)))
			listProjectsSubcommand.Options(command.WithRunE(listproject.RunEHandler(cliHandler)))
			getGroupSubcommand.Options(command.WithRunE(getgroup.RunEHandler(cliHandler)))
			getProjectSubcommand.Options(command.WithRunE(getproject.RunEHandler(cliHandler)))
			cloneSubcommand.Options(command.WithRunE(clone.RunEHandler(cliHandler)))

			return nil
		},
	}
	gitlabCmd.PersistentFlags().StringVar(&configFile, "config", "", "Configuration file")

	getGroupSubcommand = getgroup.NewCommand()
	getProjectSubcommand = getproject.NewCommand()
	getSubcommand = get.NewCommand()
	versionSubcommand = version.NewCommand()

	listGroupsSubcommand = listgroup.NewCommand()
	listProjectsSubcommand = listproject.NewCommand()
	listSubcommand = list.NewCommand()

	cloneSubcommand = clone.NewCommand()

	gitlabCommand := &command.AppCommand{
		Command: gitlabCmd,
	}

	getSubcommand.AddCommands(getGroupSubcommand, getProjectSubcommand)
	listSubcommand.AddCommands(listGroupsSubcommand, listProjectsSubcommand)
	gitlabCommand.AddCommands(getSubcommand, listSubcommand, cloneSubcommand, versionSubcommand)

	return gitlabCommand
}
