# go-log-lint

Линтер для проверки стиля логов в Go (slog, zap).

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
./gologlint ./...
```

Вариант Б: Через golangci-lint
- 1. Добавьте в .golangci.yml:
```YAML
linters-settings:
  custom:
    gologlint:
      path: ./gologlint.so
      description: checks log messages style
      original-url: [github.com/filimonq/go-log-lint](https://github.com/filimonq/go-log-lint)

linters:
  enable:
    - gologlint
```
- 2. Запустите:
```bash
golangci-lint run
```