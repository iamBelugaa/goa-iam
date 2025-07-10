package logger

import (
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
