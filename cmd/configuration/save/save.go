package save

import (
	"fmt"

	"github.com/spf13/afero"
)

type Configurer interface {
	SetConfigFile(file string)
	MergeConfigMap(cfg map[string]interface{}) error
}

type ConfigurationWriter interface {
	Configurer
	WriteConfig() error
	WriteConfigAs(file string) error
	SafeWriteConfig() error
	SafeWriteConfigAs(file string) error
}

type ConfigurationMapper interface {
	ToMap() map[string]interface{}
}

type ConfigurationSaver interface {
	Save(w ConfigurationWriter, fs afero.Fs, config ConfigurationMapper, configFile string) error
}

type SafeSave struct{}

func (s *SafeSave) Save(w ConfigurationWriter, fs afero.Fs, config ConfigurationMapper, configFile string) error {

	var err error

	err = w.MergeConfigMap(config.ToMap())
	if err != nil {
		return err
	}

	if configFile != "" {
		err = w.SafeWriteConfigAs(configFile)
	} else {
		err = w.SafeWriteConfig()
	}

	if err != nil {
		return fmt.Errorf("error saving configuration to file. %s", err.Error())
	}

	return nil
}

type Save struct{}

func (s *Save) Save(w ConfigurationWriter, fs afero.Fs, config ConfigurationMapper, configFile string) error {

	var err error

	err = w.MergeConfigMap(config.ToMap())
	if err != nil {
		return err
	}

	if configFile != "" {
		err = w.WriteConfigAs(configFile)
	} else {
		err = w.WriteConfig()
	}

	if err != nil {
		return fmt.Errorf("error saving configuration to file. %s", err.Error())
	}

	return nil
}
