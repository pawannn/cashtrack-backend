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

// Debug logs a debug message with reqID and optional attributes
func (l *Logger) Debug(reqID string, msg string, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelDebug, msg, append([]any{
		"service", l.name,
		"reqID", reqID,
		"ts", time.Now().UTC(),
	}, attrs...)...)
}

// Error logs an error message with reqID and optional attributes
func (l *Logger) Error(reqID string, msg string, err error, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelError, msg, append([]any{
		"service", l.name,
		"reqID", reqID,
		"error", err,
		"ts", time.Now().UTC(),
	}, attrs...)...)
}

// Info logs an info message with reqID and optional attributes
func (l *Logger) Info(reqID string, msg string, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelInfo, msg, append([]any{
		"service", l.name,
		"reqID", reqID,
		"ts", time.Now().UTC(),
	}, attrs...)...)
}

// Http logs an HTTP request with reqID and optional attributes
func (l *Logger) Http(reqID string, method, path string, status int, duration time.Duration, attrs ...any) {
	l.slog.Log(context.Background(), slog.LevelInfo, "HTTP request", append([]any{
		"service", l.name,
		"reqID", reqID,
		"method", method,
		"path", path,
		"status", status,
		"duration_ms", duration.Milliseconds(),
		"ts", time.Now().UTC(),
	}, attrs...)...)
}
