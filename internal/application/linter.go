package application

import (
	"github.com/filimonq/go-log-lint/internal/domain"
	"github.com/filimonq/go-log-lint/internal/infrastructure/rules"
)

type Linter struct {
	config *domain.Config
}

func NewLinter(cfg *domain.Config) *Linter {
	return &Linter{config: cfg}
}

func (l *Linter) EnabledRules() []domain.Rule {
	var activeRules []domain.Rule

	isEnabled := func(name string) bool {
		if val, ok := l.config.Rules[name]; ok {
			return val.Enabled
		}
		return true
	}

	if isEnabled("LowercaseRule") {
		activeRules = append(activeRules, rules.NewLowercaseRule())
	}
	if isEnabled("SpecialCharsRule") {
		activeRules = append(activeRules, rules.NewSpecialCharsRule())
	}
	if isEnabled("SensitiveDataRule") {
		var patterns []string
		if cfg, ok := l.config.Rules["SensitiveDataRule"]; ok {
			patterns = cfg.Patterns
		}
		activeRules = append(activeRules, rules.NewSensitiveDataRule(patterns))
	}
	if isEnabled("EnglishOnlyRule") {
		activeRules = append(activeRules, rules.NewEnglishRule())
	}

	return activeRules
}
