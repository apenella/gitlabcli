package save

import (
	"fmt"

	"github.com/spf13/afero"
)

// Configurer interface is an interface to configure a configuration
type Configurer interface {
	SetConfigFile(file string)
	MergeConfigMap(cfg map[string]interface{}) error
}

// ConfigurationWriter interface is an interface to write a configuration
type ConfigurationWriter interface {
	Configurer
	WriteConfig() error
	WriteConfigAs(file string) error
	SafeWriteConfig() error
	SafeWriteConfigAs(file string) error
}

// ConfigurationMapper interface is an interface to map a configuration
type ConfigurationMapper interface {
	ToMap() map[string]interface{}
}

// ConfigurationSaver interface is an interface to save a configuration
type ConfigurationSaver interface {
	Save(w ConfigurationWriter, fs afero.Fs, config ConfigurationMapper, configFile string) error
}

// SafeSave struct is a struct to save a configuration safely, withouth overwriting the configuration file
type SafeSave struct{}

// Save method is responsible to save a configuration safely
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

// Save struct is a struct to save a configuration
type Save struct{}

// Save method is responsible to save a configuration
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
