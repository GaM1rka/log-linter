package golangci

import (
	"github.com/GaM1rka/log-linter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}
