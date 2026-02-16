package rules

import (
	"unicode"

	"github.com/filimonq/go-log-lint/internal/domain"
)

type EnglishRule struct{}

func NewEnglishRule() *EnglishRule {
	return &EnglishRule{}
}

func (r *EnglishRule) Name() string {
	return "EnglishRule"
}

func (r *EnglishRule) Check(entry domain.LogEntry) []domain.Issue {
	for _, char := range entry.Message {
		if char > unicode.MaxASCII {
			return []domain.Issue{{
				Pos:     entry.Pos,
				Message: "message should contain only English (ASCII) characters",
			}}
		}
	}
	return nil
}
