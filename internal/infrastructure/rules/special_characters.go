package rules

import (
	"strings"

	"github.com/filimonq/go-log-lint/internal/domain"
)

type SpecialCharsRule struct{}

func NewSpecialCharsRule() *SpecialCharsRule {
	return &SpecialCharsRule{}
}

func (r *SpecialCharsRule) Name() string {
	return "SpecialCharsRule"
}

func (r *SpecialCharsRule) Check(entry domain.LogEntry) []domain.Issue {
	// emojis and other non-ASCII characters are handled by EnglishRule
	if strings.ContainsAny(entry.Message, "!?") {
		return []domain.Issue{{
			Pos:     entry.Pos,
			Message: "message should not contain '!' or '?' characters",
		}}
	}
	return nil
}
