package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *slog.Logger
)

// InitializeLogger initializes the global logger with the given level.
func InitializeLogger(logLevel slog.Level) {
	once.Do(func() { // 保証: 初期化は一度だけ
		levelVar := new(slog.LevelVar)
		levelVar.Set(logLevel)

		opts := slog.HandlerOptions{
			Level: levelVar,
		}

		handler := slog.NewJSONHandler(os.Stdout, &opts)
		logger = slog.New(handler)
	})
}

// Info logs an informational message.
func Info(msg string, keysAndValues ...any) {
	if logger == nil {
		panic("Logger not initialized. Call InitializeLogger first.")
	}
	logger.Info(msg, keysAndValues...)
}

// Debug logs a debug message.
func Debug(msg string, keysAndValues ...any) {
	if logger == nil {
		panic("Logger not initialized. Call InitializeLogger first.")
	}
	logger.Debug(msg, keysAndValues...)
}

// Error logs an error message.
func Error(msg string, err error, keysAndValues ...any) {
	if logger == nil {
		panic("Logger not initialized. Call InitializeLogger first.")
	}
	logger.Error(msg, append(keysAndValues, "error", err)...)
}