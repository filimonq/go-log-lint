package rules

import (
	"testing"

	"github.com/filimonq/go-log-lint/internal/domain"
)

func TestLowercaseRule(t *testing.T) {
	rule := NewLowercaseRule()

	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{"Valid lowercase", "success user login", false},
		{"Invalid uppercase", "Success user login", true},
		{"Valid with number", "123 record found", false},
		{"Invalid russian uppercase", "Ошибка доступа", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(domain.LogEntry{Message: tt.message})
			if (len(issues) > 0) != tt.wantErr {
				t.Errorf("Check('%s') issues = %d, wantErr %v", tt.message, len(issues), tt.wantErr)
			}
		})
	}
}

func TestEnglishRule(t *testing.T) {
	rule := NewEnglishRule()

	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{"Pure english", "all ok", false},
		{"Russian text", "ошибка тут", true},
		{"Mixed text", "error здесь", true},
		{"Special chars ok", "error: #123!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(domain.LogEntry{Message: tt.message})
			if (len(issues) > 0) != tt.wantErr {
				t.Errorf("Check('%s') issues = %d, wantErr %v", tt.message, len(issues), tt.wantErr)
			}
		})
	}
}

func TestSpecialCharsRule(t *testing.T) {
	rule := NewSpecialCharsRule()

	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{"Normal message", "process finished", false},
		{"With exclamation", "attention!", true},
		{"With question", "why?", true},
		{"Multiple issues", "what?!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(domain.LogEntry{Message: tt.message})
			if (len(issues) > 0) != tt.wantErr {
				t.Errorf("Check('%s') issues = %d, wantErr %v", tt.message, len(issues), tt.wantErr)
			}
		})
	}
}

func TestSensitiveDataRule(t *testing.T) {
	rule := NewSensitiveDataRule(nil)

	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{"Safe message", "user logged in", false},
		{"Password leak", "user password is 123", true},
		{"Token leak", "auth token invalid", true},
		{"Secret leak", "my secret key", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := rule.Check(domain.LogEntry{Message: tt.message})
			if (len(issues) > 0) != tt.wantErr {
				t.Errorf("Check('%s') issues = %d, wantErr %v", tt.message, len(issues), tt.wantErr)
			}
		})
	}
}
