// Package logger provides a structured and configurable logging utility
// for the IAM service. It wraps the zap logger with additional configuration
// logic based on environment and service-level details.
package logger

import (
	"fmt"
	"os"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger wraps zap.SugaredLogger to provide structured and leveled logging,
// along with configuration from the IAM service.
type logger struct {
	*zap.SugaredLogger
	config *config.Logging
}

// NewWithConfig creates and returns a new Logger instance based on the provided configuration.
func NewWithConfig(service, version string, environment config.Environment, cfg *config.Logging) (*logger, error) {
	// Parse the log level from the config string.
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level : %w", err)
	}

	// Use development config and encoder as defaults.
	zapConfig := zap.NewDevelopmentConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	// For production, switch to production-safe configuration and encoding.
	if environment == config.EnvironmentProduction {
		zapConfig = zap.NewProductionConfig()
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		// Enable colored log levels for better readability in development.
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Customize the zap config.
	zapConfig.OutputPaths = []string{"stderr"}
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	// Customize the encoder keys and format.
	zapConfig.EncoderConfig = encoderConfig
	zapConfig.EncoderConfig.LevelKey = "level"
	zapConfig.EncoderConfig.CallerKey = "caller"
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.MessageKey = "message"
	zapConfig.EncoderConfig.StacktraceKey = "stacktrace"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Set initial fields that will be included with every log entry.
	zapConfig.InitialFields = map[string]any{
		"service": service,
		"version": version,
		"pid":     os.Getpid(),
	}

	return &logger{
		config:        cfg,
		SugaredLogger: zap.Must(zapConfig.Build()).Sugar(),
	}, nil
}

// LogRequest logs structured metadata about an HTTP request at the info level.
func (l *logger) LogRequest(msg, method, path, userID, clientIP, duration string, statusCode int) {
	l.Infow(msg,
		UserID(userID),
		zap.String("path", path),
		zap.String("method", method),
		zap.String("duration", duration),
		zap.String("client_ip", clientIP),
		zap.Int("status_code", statusCode),
	)
}

// Close flushes any buffered log entries to the output and releases resources.
func (l *logger) Close() error {
	return l.Sync()
}
