package context

import "context"

// contextKey defines a typed string for storing and retrieving log metadata from context.
type contextKey string

// Predefined context keys for structured request and user metadata.
var (
	UserIDContextKey    contextKey = "user_id"
	TracingIDContextKey contextKey = "tracing_id"
	RequestIDContextKey contextKey = "request_id"
)

// GetRequestID retrieves the request ID from the given context.
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDContextKey).(string); ok {
		return id
	}
	return ""
}

// SetRequestID sets the request ID in the given context.
func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDContextKey, requestID)
}

// GetUserID extracts the user ID from the context.
func GetUserID(ctx context.Context) string {
	if id, ok := ctx.Value(UserIDContextKey).(string); ok {
		return id
	}
	return ""
}

// SetUserID sets the user ID in the context.
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDContextKey, userID)
}

// GetTraceID extracts the trace ID from the context for distributed tracing.
func GetTraceID(ctx context.Context) string {
	if id, ok := ctx.Value(TracingIDContextKey).(string); ok {
		return id
	}
	return ""
}

// SetTraceID sets the trace ID in the context for distributed tracing.
func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TracingIDContextKey, traceID)
}
