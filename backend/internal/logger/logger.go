package logger

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ContextKey is the type for context keys
type ContextKey string

const (
	// TraceIDKey is the context key for trace ID
	TraceIDKey ContextKey = "trace_id"
)

// Logger interface defines logging methods
type Logger interface {
	Infof(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
}

// DefaultLogger is the default implementation of Logger
type DefaultLogger struct {
	level LogLevel
}

// LogLevel represents the logging level
type LogLevel int

const (
	// DebugLevel for debug messages
	DebugLevel LogLevel = iota
	// InfoLevel for info messages
	InfoLevel
	// WarnLevel for warning messages
	WarnLevel
	// ErrorLevel for error messages
	ErrorLevel
)

var defaultLogger *DefaultLogger

func init() {
	defaultLogger = &DefaultLogger{
		level: InfoLevel,
	}
}

// New creates a new logger instance
func New() *DefaultLogger {
	return &DefaultLogger{
		level: InfoLevel,
	}
}

// SetLevel sets the logging level
func (l *DefaultLogger) SetLevel(level LogLevel) {
	l.level = level
}

// getCallerInfo returns the function name and line number of the caller
func getCallerInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown:0"
	}

	// Get function name
	fn := runtime.FuncForPC(pc)
	fnName := "unknown"
	if fn != nil {
		fnName = fn.Name()
		// Extract just the function name without the package path
		parts := strings.Split(fnName, "/")
		fnName = parts[len(parts)-1]
	}

	// Extract just the filename without the full path
	fileParts := strings.Split(file, "/")
	fileName := fileParts[len(fileParts)-1]

	return fmt.Sprintf("%s:%d [%s]", fileName, line, fnName)
}

// GetTraceID extracts or generates a trace ID from context
func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return "no-trace-id"
	}

	if traceID, ok := ctx.Value(TraceIDKey).(string); ok && traceID != "" {
		return traceID
	}

	return "no-trace-id"
}

// WithTraceID adds a trace ID to the context
func WithTraceID(ctx context.Context) context.Context {
	traceID := uuid.New().String()
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// WithTraceIDValue adds a specific trace ID to the context
func WithTraceIDValue(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// formatLog formats a log message with trace ID, caller info, and timestamp
func formatLog(ctx context.Context, level string, msg string) string {
	traceID := GetTraceID(ctx)
	callerInfo := getCallerInfo(3) // skip 3 levels: formatLog -> log method -> caller
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")

	return fmt.Sprintf("[%s] [%s] [%s] %s | %s", timestamp, level, traceID, callerInfo, msg)
}

// Infof logs an info message with formatting
func (l *DefaultLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	if l.level > InfoLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	log.Println(formatLog(ctx, "INFO", msg))
}

// Info logs an info message
func (l *DefaultLogger) Info(ctx context.Context, msg string) {
	if l.level > InfoLevel {
		return
	}
	log.Println(formatLog(ctx, "INFO", msg))
}

// Errorf logs an error message with formatting
func (l *DefaultLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	if l.level > ErrorLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	log.Println(formatLog(ctx, "ERROR", msg))
}

// Error logs an error message
func (l *DefaultLogger) Error(ctx context.Context, msg string) {
	if l.level > ErrorLevel {
		return
	}
	log.Println(formatLog(ctx, "ERROR", msg))
}

// Warnf logs a warning message with formatting
func (l *DefaultLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	if l.level > WarnLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	log.Println(formatLog(ctx, "WARN", msg))
}

// Warn logs a warning message
func (l *DefaultLogger) Warn(ctx context.Context, msg string) {
	if l.level > WarnLevel {
		return
	}
	log.Println(formatLog(ctx, "WARN", msg))
}

// Debugf logs a debug message with formatting
func (l *DefaultLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	if l.level > DebugLevel {
		return
	}
	msg := fmt.Sprintf(format, args...)
	log.Println(formatLog(ctx, "DEBUG", msg))
}

// Debug logs a debug message
func (l *DefaultLogger) Debug(ctx context.Context, msg string) {
	if l.level > DebugLevel {
		return
	}
	log.Println(formatLog(ctx, "DEBUG", msg))
}

// Package-level convenience functions using the default logger

// Infof logs an info message with formatting using the default logger
func Infof(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Infof(ctx, format, args...)
}

// Info logs an info message using the default logger
func Info(ctx context.Context, msg string) {
	defaultLogger.Info(ctx, msg)
}

// Errorf logs an error message with formatting using the default logger
func Errorf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Errorf(ctx, format, args...)
}

// Error logs an error message using the default logger
func Error(ctx context.Context, msg string) {
	defaultLogger.Error(ctx, msg)
}

// Warnf logs a warning message with formatting using the default logger
func Warnf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Warnf(ctx, format, args...)
}

// Warn logs a warning message using the default logger
func Warn(ctx context.Context, msg string) {
	defaultLogger.Warn(ctx, msg)
}

// Debugf logs a debug message with formatting using the default logger
func Debugf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Debugf(ctx, format, args...)
}

// Debug logs a debug message using the default logger
func Debug(ctx context.Context, msg string) {
	defaultLogger.Debug(ctx, msg)
}
