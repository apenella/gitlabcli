package configuration

import (
	"github.com/fatih/color"
)

const (
	// BaseURLKey is the key used to store the Gitlab server url
	BaseURLKey = "gitlab_api_url"
	// TokenKey is the key used to store the token to authenticate
	TokenKey = "gitlab_token"
	// WorkingDirKey is the key used to store the working directory
	WorkingDirKey = "working_dir"
)

// ConfigValidator interface used to validate configuration
type ConfigValidator interface {
	Struct(s interface{}) error
}

// Configuration for gitlabcli
type Configuration struct {
	// DEPRECATED_BaseURL
	DEPRECATEDBaseURL string `mapstructure:"base_url"`
	// Gitlab server url
	BaseURL string `mapstructure:"gitlab_api_url" validate:"required,url"`
	// Token to authenticate
	Token string `mapstructure:"gitlab_token" validate:"required"`
	// workingDir
	WorkingDir string `mapstructure:"working_dir,omitempty" validate:"required"`
}

// New return a configuration struct
func New(url, token, workingDir string) *Configuration {
	return &Configuration{
		BaseURL:    url,
		Token:      token,
		WorkingDir: workingDir,
	}
}

// ToMap returns a map with configuration
func (c *Configuration) ToMap() map[string]interface{} {
	return map[string]interface{}{
		BaseURLKey:    c.BaseURL,
		TokenKey:      c.Token,
		WorkingDirKey: c.WorkingDir,
	}
}

// FixCompatibility prepares configuration to valid attributes instead of deprecated ones
func (c *Configuration) FixCompatibility() {
	if c.DEPRECATEDBaseURL != "" {
		color.Cyan("[DEPRECATED] Use 'gitlab_api_url' configuration parameter instead of 'base_url'")
		if c.BaseURL == "" {
			c.BaseURL = c.DEPRECATEDBaseURL
		}
	}
}

// Validate checks configuration
func (c *Configuration) Validate(validator ConfigValidator) error {
	return validator.Struct(c)
}
