package spoke

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Spoke"
)

type Settings struct {
	common *cfg.Common

	apiKey   string
	username string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("ZENDESK_API"))),
		username: ymlConfig.UString("username"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
