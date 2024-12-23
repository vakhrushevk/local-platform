package logger

import (
	"log/slog"
	"os"

	"github.com/vakhrushevk/local-platform/pkg/logger/slogpretty"
)

// TODO: мб удалить
const (
	EnvDebug      = "Debug"
	EnvProduction = "Info"
)

var globalSlogger *slog.Logger

func Init() {
	globalSlogger = setupPrettySlog()
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}

func Debug(msg string, args ...interface{}) {
	globalSlogger.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	globalSlogger.Info(msg, args...)
}

func Error(msg string, args ...interface{}) {
	globalSlogger.Error(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	globalSlogger.Warn(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	globalSlogger.Error(msg, args...)
	panic(msg)
}
