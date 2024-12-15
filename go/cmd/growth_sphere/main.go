package main

import (
	"GROWTHSPHERE/pkg/logger"
	"log/slog"
)

func main() {
	logger.InitializeLogger(slog.LevelInfo)
	logger.Info("Start GROWTHSPHERE")
}
