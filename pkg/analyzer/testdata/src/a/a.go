package a

import (
	"log/slog"
)

func checkLogs() {
	// 1. Ошибка: заглавная буква
	slog.Info("Hello world") // want `\[LowercaseRule\] message should start with lowercase letter: 'H'`

	// 2. Ошибка: спецсимволы
	slog.Warn("danger!") // want `\[no_special_chars\] message should not contain '!' or '\?' characters`

	// 3. Ошибка: секретные данные
	slog.Debug("user password is 123") // want `\[sensitive_data\] message contains potential sensitive data: 'password'`

	// 4. Ошибка: русский текст
	slog.Error("ошибка") // want `\[EnglishRule\] message should contain only English \(ASCII\) characters`

	// 5. Правильный лог — тут линтер должен молчать
	slog.Info("all systems nominal")
}
