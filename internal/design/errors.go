package design

import (
	"github.com/iamBelugaa/goa-iam/internal/domain/codes"
	"goa.design/goa/v3/dsl"
)

// ErrorCode is a string enum representing error categories.
var ErrorCode = dsl.Type("ErrorCode", dsl.String, func() {
	dsl.Description("A string enum representing API error codes for standardized error handling.")
	dsl.Enum(
		codes.NotFoundErrCode,
		codes.ValidationErrCode,
		codes.InternalServerErrCode,
		codes.UnauthorizedErrCode,
	)
	dsl.Example(codes.ValidationErrCode)
})

// ErrorDetail provides detailed error information.
var ErrorDetail = dsl.Type("ErrorDetail", func() {
	dsl.Description("Detailed error information")

	dsl.Attribute("field", dsl.String, "Field that caused the error", func() {
		dsl.Example("email")
	})

	dsl.Attribute("message", dsl.String, "Detailed error message", func() {
		dsl.Example("Email format is invalid")
	})

	dsl.Attribute("code", ErrorCode, "Error code for programmatic handling", func() {
		dsl.Example("INTERNAL_SERVER_ERROR")
	})

	dsl.Required("message", "code")
})

// ValidationError represents validation error details.
var ValidationError = dsl.Type("ValidationError", func() {
	dsl.Description("Validation error response")

	dsl.Attribute("message", dsl.String, "Error message", func() {
		dsl.Example("Validation failed")
	})

	dsl.Attribute("details", dsl.ArrayOf(ErrorDetail), "Detailed validation error")

	dsl.Attribute("code", ErrorCode, "Error code", func() {
		dsl.Example("VALIDATION_ERROR")
	})

	dsl.Required("message", "details")
})

// ConflictError represents a conflict error (e.g., duplicate resource).
var ConflictError = dsl.Type("ConflictError", func() {
	dsl.Description("Conflict error response")

	dsl.Attribute("message", dsl.String, "Error message", func() {
		dsl.Example("Resource already exists")
	})

	dsl.Attribute("code", ErrorCode, "Error code", func() {
		dsl.Example("RESOURCE_CONFLICT")
	})

	dsl.Required("message", "code")
})

// UnauthorizedError represents an authentication error.
var UnauthorizedError = dsl.Type("UnauthorizedError", func() {
	dsl.Description("Unauthorized error response")

	dsl.Attribute("message", dsl.String, "Error message", func() {
		dsl.Example("Authentication required")
	})

	dsl.Attribute("code", ErrorCode, "Error code", func() {
		dsl.Example("UNAUTHORIZED")
	})

	dsl.Required("message", "code")
})

// NotFoundError represents a resource not found error.
var NotFoundError = dsl.Type("NotFoundError", func() {
	dsl.Description("Not found error response")

	dsl.Attribute("message", dsl.String, "Error message", func() {
		dsl.Example("Resource not found")
	})

	dsl.Attribute("code", ErrorCode, "Error code", func() {
		dsl.Example("NOT_FOUND")
	})

	dsl.Required("message", "code")
})

// InternalServerError represents an internal server error.
var InternalServerError = dsl.Type("InternalServerError", func() {
	dsl.Description("Internal server error response")

	dsl.Attribute("message", dsl.String, "Error message", func() {
		dsl.Example("Internal server error")
	})

	dsl.Attribute("code", ErrorCode, "Error code", func() {
		dsl.Example("INTERNAL_SERVER_ERROR")
	})

	dsl.Required("message", "code")
})
