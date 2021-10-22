package load

import (
	"fmt"

	"github.com/apenella/gitlabcli/internal/configuration"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type ConfigurationLoader interface {
	ReadInConfig() error
	SetConfigFile(file string)
	SetFs(afero.Fs)
}

type ConfigurationUnmarshaler interface {
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
}

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

func Unmarshal(unmarshaler ConfigurationUnmarshaler) (configuration.Configuration, error) {
	var err error
	var conf configuration.Configuration

	err = unmarshaler.Unmarshal(&conf)
	if err != nil {
		return conf, fmt.Errorf("error unmarshaling configuration. %s", err.Error())
	}

	return conf, nil
}
