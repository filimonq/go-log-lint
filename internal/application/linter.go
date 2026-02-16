package application

import (
	"github.com/filimonq/go-log-lint/internal/domain"
	"github.com/filimonq/go-log-lint/internal/infrastructure/rules"
)

type Linter struct{}

func NewLinter() *Linter {
	return &Linter{}
}

func (l *Linter) EnabledRules() []domain.Rule {
	return []domain.Rule{
		rules.NewLowercaseRule(),
		rules.NewEnglishRule(),
		rules.NewSpecialCharsRule(),
		rules.NewSensitiveDataRule(),
	}
}
