package apiserver

import "log/slog"

// NewLogger creates a new logger instance.
// It uses the default logger configuration.
func NewLogger() *slog.Logger {
	return slog.Default()
}
