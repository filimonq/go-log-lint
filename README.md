# go-log-lint

Линтер для проверки стиля логов в Go.

## 📦 Поддерживаемые логгеры
Линтер автоматически определяет и проверяет вызовы из следующих библиотек:
* **Standard Library:** `log/slog`
* **Uber Zap:** `go.uber.org/zap`

## ✅ Что проверяем (Rules)
1. **LowercaseRule** — Сообщение с маленькой буквы.
2. **NoSpecialCharsRule** — Нет знаков `!` `?`.
3. **SensitiveDataRule** — Нет паролей, токенов и т.д.
4. **EnglishOnlyRule** — Только английский язык.

## 🔨 Сборка

1. Обычная версия (Standalone):
```bash
go build -o gologlint cmd/go-log-lint/main.go
```

2. Плагин для golangci-lint:
```Bash
go build -buildmode=plugin -o gologlint.so plugin/main.go
```

3. Запуск:

Вариант А: Ручной запуск
```bash
./gologlint ./demo/...
```

С авто-исправлением (на данный момент фиксит 2 ошибки):
```bash
./gologlint -fix ./demo/...
```

Вариант Б: Через golangci-lint
- 1. Добавьте в .golangci.yml:
```YAML
linters-settings:
  custom:
    gologlint:
      path: ./gologlint.so
      description: checks log messages style
      original-url: https://github.com/filimonq/go-log-lint

linters:
  enable:
    - gologlint
```
- 2. Запустите:
```bash
golangci-lint run
```

С авто-исправлением (на данный момент фиксит 2 ошибки):
```bash
golangci-lint run --fix
```

## ⚙️ Конфигурация

По умолчанию включены все правила. Вы можете отключить ненужные или добавить свои паттерны для поиска секретов через YAML-файл.

**1. Создайте файл `config.yaml`:**

```yaml
rules:
  SpecialCharsRule:
    enabled: false

  SensitiveDataRule:
    enabled: true
    patterns:
      - "(?i)my_secret_key"
      - secret
      - api_key
      - token

  LowercaseRule:
    enabled: true
  EnglishOnlyRule:
    enabled: true
```

2. Запустите линтер с флагом -config:
```bash
./gologlint -config config.yaml ./...
```