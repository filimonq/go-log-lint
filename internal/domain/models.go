package domain

import "go/token"

type LogEntry struct {
	Message  string    // The log message
	Package  string    // Package is the name of the package where the log entry is located
	Function string    // Info, Error, Debug, etc.
	Pos      token.Pos // Position in the source code where the log entry is located
}

type Issue struct {
	Message string    // Description of the issue
	Pos     token.Pos // Position in the source code where the issue is located
}
