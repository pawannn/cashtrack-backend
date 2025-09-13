package logger

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	slog *slog.Logger
	name string
}

func InitNewLogger(serviceName string) Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	return Logger{
		name: serviceName,
		slog: slog.New(handler),
	}
}

func (l *Logger) Debug(msg string, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelDebug, msg, append([]any{"service", l.name, "ts", time.Now().UTC()}, attrs...)...)
}

func (l *Logger) Error(msg string, err error, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelError, msg,
		append([]any{"service", l.name, "error", err, "ts", time.Now().UTC()}, attrs...)...)
}

func (l *Logger) Info(msg string, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelInfo, msg,
		append([]any{"service", l.name, "ts", time.Now().UTC()}, attrs...)...)
}

func (l *Logger) Http(method, path string, status int, duration time.Duration, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelInfo, "HTTP request",
		append([]any{
			"service", l.name,
			"method", method,
			"path", path,
			"status", status,
			"duration_ms", duration.Milliseconds(),
			"ts", time.Now().UTC(),
		}, attrs...)...)
}
