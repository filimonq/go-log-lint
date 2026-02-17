package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("Hello world?")
	slog.Warn("Warning!!!")
	slog.Info("user api_key is 123")
	logger, _ := zap.NewProduction()
	logger.Info("Error!")
}
