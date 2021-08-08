package configuration

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

const (
	DefaultConfigFile = "config"

	BaseUrlKey    = "base_url"
	TokenKey      = "gitlab_token"
	WorkingDirKey = "working_dir"
)

// Configuration for gitlabcli
type Configuration struct {
	// Gitlab server url
	BaseURL string `mapstructure:"base_url" validate:"required"`
	// Token to authenticate
	Token string `mapstructure:"gitlab_token" validate:"required"`
	// workingDir
	WorkingDir string `mapstructure:"working_dir,omitempty" validate:"required"`
}

func New(configFile string) (Configuration, error) {
	var err error
	var conf Configuration

	user, err := user.Current()
	if err != nil {
		return conf, fmt.Errorf("current user information can not be achieved. %s", err.Error())
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("gitlabcli")
	viper.SetConfigName(DefaultConfigFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(user.HomeDir, ".config", "gitlabcli"))

	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	err = viper.ReadInConfig()
	if err != nil {
		_, isConfigFileNotFoundError := err.(viper.ConfigFileNotFoundError)
		if !isConfigFileNotFoundError {
			return conf, fmt.Errorf("error reading configuration from '%s'. %s", configFile, err.Error())
		}
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return conf, fmt.Errorf("error unmarshaling configuration. %s", err.Error())
	}

	return conf, nil
}

func (c Configuration) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
