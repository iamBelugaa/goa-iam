// Package logger provides structured, context-aware, and configurable logging utilities
// for the IAM service.
package logger

import (
	"fmt"
	"os"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap.SugaredLogger and holds logging config information.
type Logger struct {
	*zap.SugaredLogger
	config *config.Logging
}

// NewWithConfig initializes a new logger using the provided service name,
// version, environment, and configuration. It sets up the encoder, log level,
// and output format based on whether the environment is development or production.
func NewWithConfig(service, version string, environment config.Environment, cfg *config.Logging) (*Logger, error) {
	// Parse the log level from the config string.
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level : %w", err)
	}

	// Use development config and encoder as defaults.
	zapConfig := zap.NewDevelopmentConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	// Switch to production settings if applicable.
	if environment == config.EnvironmentProduction {
		zapConfig = zap.NewProductionConfig()
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		// Add color for easier reading in development logs.
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Apply common configurations.
	zapConfig.Encoding = "json"
	zapConfig.DisableCaller = false
	zapConfig.DisableStacktrace = false
	zapConfig.Level = zap.NewAtomicLevelAt(level)

	zapConfig.OutputPaths = []string{"stderr"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	// Customize encoder keys and format.
	zapConfig.EncoderConfig = encoderConfig
	zapConfig.EncoderConfig.LevelKey = "level"
	zapConfig.EncoderConfig.CallerKey = "caller"
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.MessageKey = "message"
	zapConfig.EncoderConfig.StacktraceKey = "stacktrace"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Add service-level fields to every log entry.
	zapConfig.InitialFields = map[string]any{
		"service": service,
		"version": version,
		"pid":     os.Getpid(),
	}

	return &Logger{
		config: cfg,
		SugaredLogger: zap.Must(
			zapConfig.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)),
		).Sugar(),
	}, nil
}

// LogRequest logs an HTTP request with relevant metadata,
// including user ID, HTTP method, path, client IP, duration, and status code.
func (l *Logger) LogRequest(msg, method, path, userID, clientIP, duration string, statusCode int) {
	l.Infow(msg,
		UserID(userID),
		zap.String("path", path),
		zap.String("method", method),
		zap.String("duration", duration),
		zap.String("client_ip", clientIP),
		zap.Int("status_code", statusCode),
	)
}

// Close ensures that any buffered logs are flushed to the output.
func (l *Logger) Close() error {
	return l.Sync()
}
