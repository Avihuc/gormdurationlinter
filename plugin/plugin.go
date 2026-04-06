package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/Avihuc/gormdurationlinter"
	"golang.org/x/tools/go/analysis"
)

type gormdurationPlugin struct{}

func (p *gormdurationPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{gormdurationlinter.Analyzer}, nil
}

func (p *gormdurationPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func init() {
	register.Plugin("gormduration", New)
}

func New(conf any) (register.LinterPlugin, error) {
	return &gormdurationPlugin{}, nil
}
