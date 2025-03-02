package logger

import (
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel defines the log level
type LogLevel string

const (
	// DebugLevel logs debug or higher logs
	DebugLevel LogLevel = "debug"
	// InfoLevel logs info or higher logs
	InfoLevel LogLevel = "info"
	// WarnLevel logs warn or higher logs
	WarnLevel LogLevel = "warn"
	// ErrorLevel logs error or higher logs
	ErrorLevel LogLevel = "error"
)

// Logger defines the interface for logging
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	With(fields ...Field) Logger
	Sync() error
}

// Field represents a log field
type Field = zapcore.Field

// String creates a string field
func String(key, value string) Field {
	return zap.String(key, value)
}

// Int creates an int field
func Int(key string, value int) Field {
	return zap.Int(key, value)
}

// Error creates an error field
func Error(err error) Field {
	return zap.Error(err)
}

// Bool creates a bool field
func Bool(key string, value bool) Field {
	return zap.Bool(key, value)
}

var (
	defaultLogger     Logger
	defaultLoggerOnce sync.Once
)

// ZapLogger is a Zap implementation of the Logger interface
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger creates a new Zap logger
func NewZapLogger(level LogLevel, isProduction bool) Logger {
	// Parse the log level
	var zapLevel zapcore.Level
	switch strings.ToLower(string(level)) {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// Set up the encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	var output zapcore.WriteSyncer

	if isProduction {
		// Production setup
		encoder = zapcore.NewJSONEncoder(encoderConfig)
		output = zapcore.AddSync(os.Stdout)
	} else {
		// Development setup with colored output
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		output = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, output, zapLevel)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	return &ZapLogger{logger: zapLogger}
}

// Debug logs a debug message
func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info message
func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning message
func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error message
func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

// Fatal logs a fatal message and exits
func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

// With returns a logger with attached fields
func (l *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{logger: l.logger.With(fields...)}
}

// Sync flushes the log buffer
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// GetLogger returns the default logger
func GetLogger() Logger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewZapLogger(InfoLevel, false)
	})
	return defaultLogger
}

// InitLogger initializes the default logger
func InitLogger(level LogLevel, isProduction bool) {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewZapLogger(level, isProduction)
	})
}
