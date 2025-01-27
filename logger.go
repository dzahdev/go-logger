package logger

import (
	"github.com/dzahdev/go-logger/prettylog"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

var globalLogger *slog.Logger

func init() {
	env := getEnvLevel()
	switch env {
	case envDev:
		globalLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProd:
		globalLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case envLocal:
		prettyHandler := prettylog.NewHandler(&slog.HandlerOptions{
			Level:       slog.LevelDebug,
			AddSource:   false,
			ReplaceAttr: nil,
		})
		globalLogger = slog.New(prettyHandler)
	}
}

func getEnvLevel() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = envLocal
	}
	return env
}

func Debug(msg string, args ...interface{}) {
	globalLogger.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	globalLogger.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	globalLogger.Error(msg, args...)
}
