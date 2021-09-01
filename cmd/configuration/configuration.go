package configuration

const (
	BaseUrlKey    = "base_url"
	TokenKey      = "gitlab_token"
	WorkingDirKey = "working_dir"
)

type ConfigValidator interface {
	Struct(s interface{}) error
}

// Configuration for gitlabcli
type Configuration struct {
	// Gitlab server url
	BaseURL string `mapstructure:"base_url" validate:"required,url"`
	// Token to authenticate
	Token string `mapstructure:"gitlab_token" validate:"required"`
	// workingDir
	WorkingDir string `mapstructure:"working_dir,omitempty" validate:"required"`
}

func New(base, token, workingDir string) *Configuration {
	return &Configuration{
		BaseURL:    base,
		Token:      token,
		WorkingDir: workingDir,
	}
}

func (c *Configuration) ToMap() map[string]interface{} {
	return map[string]interface{}{
		BaseUrlKey:    c.BaseURL,
		TokenKey:      c.Token,
		WorkingDirKey: c.WorkingDir,
	}
}

func (c *Configuration) Validate(validator ConfigValidator) error {
	return validator.Struct(c)
}
