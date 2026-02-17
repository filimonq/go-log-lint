package rules

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/filimonq/go-log-lint/internal/domain"
)

type LowercaseRule struct{}

func NewLowercaseRule() *LowercaseRule {
	return &LowercaseRule{}
}

func (r *LowercaseRule) Name() string {
	return "LowercaseRule"
}

func (r *LowercaseRule) Check(entry domain.LogEntry) []domain.Issue {
	if entry.Message == "" {
		return nil
	}
	firstRune, _ := utf8.DecodeRuneInString(entry.Message)

	if unicode.IsLetter(firstRune) && unicode.IsUpper(firstRune) {
		runes := []rune(entry.Message)
		fixedMsg := string(unicode.ToLower(runes[0])) + string(runes[1:])

		return []domain.Issue{{
			Pos:         entry.Pos,
			Message:     fmt.Sprintf("message should start with lowercase letter: '%c'", firstRune),
			Replacement: fixedMsg,
		}}
	}
	return nil
}
