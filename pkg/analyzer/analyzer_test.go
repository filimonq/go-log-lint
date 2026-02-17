package analyzer_test

import (
	"testing"

	"github.com/filimonq/go-log-lint/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMyLinter(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, analyzer.Analyzer, "a")
}
