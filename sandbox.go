// sandbox.go
package main

// Имитируем логгер, чтобы код был валидным
type Logger struct{}

func (l *Logger) Info(msg string)  {}
func (l *Logger) Debug(msg string) {}
func (l *Logger) Error(msg string) {}

func main() {
	log := &Logger{}

	log.Info("Hello world")        // Fail: Uppercase
	log.Debug("Ошибка")            // Fail: Russian
	log.Error("Stop!")             // Fail: Special char
	log.Info("my password is 123") // Fail: Sensitive
	log.Info("all good")           // OK
}
