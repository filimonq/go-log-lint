package analyzer

import (
	"golang.org/x/tools/go/analysis"

	"github.com/filimonq/go-log-lint/internal/application"
	ast_visitor "github.com/filimonq/go-log-lint/internal/infrastructure/ast"
	"github.com/filimonq/go-log-lint/internal/infrastructure/config"
)

var configPath string

var Analyzer = &analysis.Analyzer{
	Name: "gologlint",
	Doc:  "checks log messages for style guide compliance",
	Run:  run,
}

func init() {
	Analyzer.Flags.StringVar(&configPath, "config", ".gologlint.yaml", "path to configuration file")
}

func run(pass *analysis.Pass) (any, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}
	linterApp := application.NewLinter(cfg)

	activeRules := linterApp.EnabledRules()

	visitor := ast_visitor.NewLogVisitor(pass, activeRules)

	visitor.Walk()

	return nil, nil
}
