package logger

import (
	"strings"

	"go.uber.org/zap"
)

// UserID field helper.
func UserID(id string) zap.Field {
	return zap.String("user_id", id)
}

// RequestID field helper.
func RequestID(id string) zap.Field {
	return zap.String("request_id", id)
}

// Operation field helper.
func Operation(name string) zap.Field {
	return zap.String("operation", name)
}

// RedactEmail redacts email addresses.
func RedactEmail(email string) string {
	if len(email) == 0 {
		return "[REDACTED]"
	}

	index := strings.Index(email, "@")
	if index == -1 {
		return "[REDACTED]"
	}

	if index > 0 {
		return string(email[0]) + "***@" + email[index+1:]
	}
	return "[REDACTED]"
}

// RedactSensitiveData redacts sensitive data.
func RedactSensitiveData(data string) string {
	if len(data) <= 4 {
		return "[REDACTED]"
	}
	return string(data[:2]) + "***" + string(data[len(data)-2])
}
