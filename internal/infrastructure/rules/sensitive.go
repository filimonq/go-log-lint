package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/filimonq/go-log-lint/internal/domain"
)

type SensitiveDataRule struct {
	patterns []*regexp.Regexp
}

func NewSensitiveDataRule(customPatterns []string) *SensitiveDataRule {
	defaults := []string{
		`(?i)password`,
		`(?i)token`,
		`(?i)secret`,
		`(?i)api_key`,
		`(?i)bearer`,
	}

	allPatterns := append(defaults, customPatterns...)

	compiled := make([]*regexp.Regexp, 0, len(allPatterns))
	for _, p := range allPatterns {
		if re, err := regexp.Compile(p); err == nil {
			compiled = append(compiled, re)
		}
	}

	return &SensitiveDataRule{patterns: compiled}
}

func (r *SensitiveDataRule) Name() string {
	return "SensitiveDataRule"
}

func (r *SensitiveDataRule) Check(entry domain.LogEntry) []domain.Issue {
	for _, re := range r.patterns {
		if loc := re.FindStringIndex(entry.Message); loc != nil {
			matchedStr := entry.Message[loc[0]:loc[1]]

			sensitiveKey := strings.Split(matchedStr, " ")[0]
			if len(sensitiveKey) > 20 {
				sensitiveKey = sensitiveKey[:20] + "..."
			}

			return []domain.Issue{{
				Pos:     entry.Pos,
				Message: fmt.Sprintf("message contains potential sensitive data: match pattern '%s'", sensitiveKey),
			}}
		}
	}
	return nil
}
