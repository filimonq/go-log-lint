package domain

type Rule interface {
	Name() string
	Check(entry LogEntry) []Issue
}
