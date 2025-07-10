// Package config defines configuration types used across the IAM service.
package config

// Environment defines the type for representing different runtime environments.
type Environment string

// Supported environment constants.
var (
	EnvironmentProduction  Environment = "PRODUCTION"
	EnvironmentDevelopment Environment = "DEVELOPMENT"
)

// Logging holds settings for how logging should behave in different environments.
type Logging struct {
	Level           string `json:"level"`           // Log level (e.g., "debug", "info", "warn", "error").
	RequestLogging  bool   `json:"requestLogging"`  // Enable/disable request-level logging.
	RedactSensitive bool   `json:"redactSensitive"` // Enable/disable sensitive data redaction in logs.
}
