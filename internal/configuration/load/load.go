package load

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// ConfigurationLoader interface is an interface to load configuration
type ConfigurationLoader interface {
	ReadInConfig() error
	SetConfigFile(file string)
	SetFs(afero.Fs)
}

// ConfigurationUnmarshaler interface is an interface to unmarshal configuration
type ConfigurationUnmarshaler interface {
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
}

// Load function is responsible to load configuration from a file
func Load(loader ConfigurationLoader, fs afero.Fs, configFile string) error {
	var err error

	if fs != nil {
		loader.SetFs(fs)
	}

	if configFile != "" {
		loader.SetConfigFile(configFile)
	}

	err = loader.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading configuration from '%s'. %s", configFile, err.Error())
	}

	return nil
}

// Unmarshal function is responsible to unmarshal configuration
func Unmarshal(unmarshaler ConfigurationUnmarshaler) (configuration.Configuration, error) {
	var err error
	var conf configuration.Configuration

	err = unmarshaler.Unmarshal(&conf)
	if err != nil {
		return conf, fmt.Errorf("error unmarshaling configuration. %s", err.Error())
	}

	return conf, nil
}
