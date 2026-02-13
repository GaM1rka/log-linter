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

func New(_ any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	// тебе нужен types info для zap (pass.TypesInfo), значит TypesInfo mode
	return register.LoadModeTypesInfo
}
