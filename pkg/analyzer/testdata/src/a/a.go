package a

import (
	"log/slog"
)

func checkLogs() {
	// 1. Ошибка: заглавная буква
	slog.Info("Hello world") // want `\[LowercaseRule\] message should start with lowercase letter: 'H'`

	slog.Warn("danger!") // want `\[SpecialCharsRule\] message should not contain '!' or '\?' characters`

	// 3. Ошибка: секретные данные
	slog.Debug("user password is 123") // want `\[SensitiveDataRule\] message contains potential sensitive data: match pattern 'password'`

	// 4. Ошибка: русский текст
	slog.Error("ошибка") // want `\[EnglishOnlyRule\] message should contain only English \(ASCII\) characters`

	// 5. Правильный лог — тут линтер должен молчать
	slog.Info("all systems nominal")
}
