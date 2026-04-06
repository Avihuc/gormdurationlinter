package gormdurationlinter_test

import (
	"testing"

	"github.com/lema-ai/gormdurationlinter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gormdurationlinter.Analyzer, "example")
}
