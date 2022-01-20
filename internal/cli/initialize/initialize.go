package initialize

import (
	"os"

	"github.com/apenella/gitlabcli/internal/configuration"
	saveconfiguration "github.com/apenella/gitlabcli/internal/configuration/save"
	"github.com/apenella/gitlabcli/pkg/command"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
)

// NewCommand creates a initialize command
func NewCommand(viperconfig *viper.Viper) *command.AppCommand {
	var baseURL, workingDir, token, configFile string
	var force bool
	var err error
	var saver saveconfiguration.ConfigurationSaver

	initCmd := &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   "Initializes gitlabcli",
		Long:    "Initializes gitlabcli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		// Errors silenced because they are propagated
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			errContext := "initialize::NewCommand::RunE"

			ui := &input.UI{
				Reader: os.Stdin,
			}

			if baseURL == "" {
				baseURL, err = ui.Ask("Which is your Gitlab base url?",
					&input.Options{
						Required:  true,
						Loop:      true,
						HideOrder: true,
					})

				if err != nil {
					return errors.New(errContext, "Base URL could not be read", err)
				}
			}

			if workingDir == "" {
				workingDir, err = ui.Ask("Which is your gitlabcli working dir?",
					&input.Options{
						Required:  true,
						Loop:      true,
						HideOrder: true,
					})

				if err != nil {
					return errors.New(errContext, "Working directory could not be read", err)
				}
			}

			token, err = ui.Ask("Which is your Gitlab user token?",
				&input.Options{
					Required:  true,
					Mask:      true,
					HideOrder: true,
					Loop:      true,
				})

			if err != nil {
				return errors.New(errContext, "Gitlab token could not be read", err)
			}

			conf := configuration.New(baseURL, token, workingDir)

			v := validator.New()
			err = conf.Validate(v)
			if err != nil {
				return errors.New(errContext, "Invalid configuration", err)
			}

			if force {
				saver = &saveconfiguration.Save{}
			} else {
				saver = &saveconfiguration.SafeSave{}
			}

			err = saver.Save(viperconfig, afero.NewOsFs(), conf, configFile)
			if err != nil {
				return errors.New(errContext, "Configuration could no be saved", err)
			}

			return nil
		},
	}

	initCmd.PersistentFlags().StringVar(&configFile, "config", "", "Configuration file")
	initCmd.Flags().StringVarP(&baseURL, "gitlab-api-url", "u", "", "Gitlab API URL base. Check it on https://docs.gitlab.com/ee/api/#how-to-use-the-api")
	initCmd.Flags().StringVarP(&workingDir, "working-dir", "d", "", "Location to store cloned gitlabcli repositories")
	initCmd.Flags().BoolVar(&force, "force", false, "Force ot override current configuration")

	return command.NewCommand(initCmd)
}
