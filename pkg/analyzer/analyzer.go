package analyzer

import (
	"golang.org/x/tools/go/analysis"

	"github.com/filimonq/go-log-lint/internal/application"
	ast_visitor "github.com/filimonq/go-log-lint/internal/infrastructure/ast"
)

var Analyzer = &analysis.Analyzer{
	Name: "gologlint",
	Doc:  "checks log messages for style guide compliance",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	linterApp := application.NewLinter()

	activeRules := linterApp.EnabledRules()

	visitor := ast_visitor.NewLogVisitor(pass, activeRules)

	visitor.Walk()

	return nil, nil
}
