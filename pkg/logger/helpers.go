package logger

import (
	"go.uber.org/zap"
)

// UserID returns a zap.Field representing the user ID.
func UserID(id string) zap.Field {
	return zap.String("user_id", id)
}

// RequestID returns a zap.Field representing the request ID.
func RequestID(id string) zap.Field {
	return zap.String("request_id", id)
}

// Operation returns a zap.Field for naming the current operation or handler.
func Operation(name string) zap.Field {
	return zap.String("operation", name)
}
