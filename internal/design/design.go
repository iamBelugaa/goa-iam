package design

import (
	"goa.design/goa/v3/dsl"
)

// API defines the global settings and metadata for the IAM Platform service.
var _ = dsl.API("iam-platform", func() {
	dsl.Title("Identity and Access Management Platform")
	dsl.Description("The IAM Platform provides core identity related functionalities.")
	dsl.Version("0.1.0")

	// License information
	dsl.License(func() {
		dsl.Name("MIT")
		dsl.URL("https://opensource.org/licenses/MIT")
	})

	// Server block defines where the services are hosted.
	dsl.Server("development", func() {
		dsl.Description("Primary server hosting the IAM service")
		dsl.Host("localhost", func() {
			dsl.Description("Local development environment")
			dsl.URI("http://localhost:8080")
		})
	})

	// Global HTTP configuration.
	dsl.HTTP(func() {
		// All endpoints will be prefixed with /api/v1
		dsl.Path("/api/v1")
	})
})

// JWTAuth defines the JWT authentication scheme used throughout the API.
var JWTAuth = dsl.JWTSecurity("JWTAuth", func() {
	dsl.Description("JWT token authentication")
})

// SuccessResponse defines a standard structure for a successful API response.
var SuccessResponse = dsl.Type("SuccessResponse", func() {
	dsl.Description("A standard API response returned when an operation completes successfully.")

	dsl.Attribute("success", dsl.Boolean, "Indicates if the operation was successful", func() {
		dsl.Example(true)
	})

	dsl.Attribute("message", dsl.String, "Human readable message", func() {
		dsl.Example("Operation completed successfully")
	})

	dsl.Attribute("data", dsl.Any, "Optional payload returned on success")

	dsl.Required("success", "message")
})

// CommonErrors defines error types that are shared across multiple services.
func CommonErrors() {
	// 400 Bad Request
	dsl.Error("bad_request", ValidationError, "Invalid request data or parameters")

	// 401 Unauthorized
	dsl.Error("unauthorized", UnauthorizedError, "Authentication required or invalid credentials")

	// 404 Not Found
	dsl.Error("not_found", NotFoundError, "Requested resource not found")

	// 409 Conflict
	dsl.Error("conflict", ConflictError, "Resource conflict or already exists")

	// 422 Unprocessable Entity
	dsl.Error("validation_failed", ValidationError, "Request validation failed")

	// 500 Internal Server Error
	dsl.Error("internal_server_error", InternalServerError, "Internal server error occurred")
}
