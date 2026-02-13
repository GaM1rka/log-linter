package golangci

import (
	"github.com/GaM1rka/log-linter/pkg/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type Plugin struct{}

func init() {
	register.Plugin("loglinter", New)
}

func New(settings any) (register.LinterPlugin, error) {
	cfg := analyzer.Config{
		EnableLowercase:   true,
		EnableEnglishOnly: true,
		EnableNoSpecial:   true,
		EnableNoSensitive: true,
	}

	if settingsMap, ok := settings.(map[string]any); ok {
		if v, ok := settingsMap["enableLowercase"].(bool); ok {
			cfg.EnableLowercase = v
		}
		if v, ok := settingsMap["enableEnglishOnly"].(bool); ok {
			cfg.EnableEnglishOnly = v
		}
		if v, ok := settingsMap["enableNoSpecial"].(bool); ok {
			cfg.EnableNoSpecial = v
		}
		if v, ok := settingsMap["enableNoSensitive"].(bool); ok {
			cfg.EnableNoSensitive = v
		}
	}

	analyzer.SetConfig(cfg)
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
