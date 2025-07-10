// Package logger provides structured, context-aware, and configurable logging utilities
// for the IAM service. It wraps Uber's zap logging library with custom helpers for
// logging contextual data such as user ID, request ID, trace ID, and supports redacting
// sensitive information from logs.
package logger
