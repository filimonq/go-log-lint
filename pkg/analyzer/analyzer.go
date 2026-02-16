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

// run — это функция запуска, которую дергает фреймворк.
func run(pass *analysis.Pass) (interface{}, error) {
	// 1. Инициализируем слой приложения (Application)
	linterApp := application.NewLinter()

	// 2. Получаем активные правила (Domain)
	activeRules := linterApp.EnabledRules()

	// 3. Создаем Инспектора (Infrastructure), скармливаем ему правила
	visitor := ast_visitor.NewLogVisitor(pass, activeRules)

	// 4. Запускаем обход
	visitor.Walk()

	return nil, nil
}
