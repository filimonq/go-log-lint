package rules

import (
	"fmt"
	"strings"

	"github.com/filimonq/go-log-lint/internal/domain"
)

type SensitiveDataRule struct {
	stopWords []string
}

func NewSensitiveDataRule() *SensitiveDataRule {
	return &SensitiveDataRule{
		stopWords: []string{
			"password", "pass", "token", "secret",
			"key", "auth", "credential", "bearer",
			"private",
		},
	}
}

func (r *SensitiveDataRule) Name() string {
	return "sensitive_data"
}

func (r *SensitiveDataRule) Check(entry domain.LogEntry) []domain.Issue {
	lowerMsg := strings.ToLower(entry.Message)

	for _, word := range r.stopWords {
		if strings.Contains(lowerMsg, word) {
			return []domain.Issue{{
				Pos:     entry.Pos,
				Message: fmt.Sprintf("message contains potential sensitive data: '%s'", word),
			}}
		}
	}
	return nil
}
