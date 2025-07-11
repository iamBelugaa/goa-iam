package design

import (
	"goa.design/goa/v3/dsl"
)

// API defines the global settings and metadata for the IAM Platform service.
var _ = dsl.API("iam-platform", func() {
	dsl.Title("Identity and Access Management Platform")
	dsl.Description(
		`The IAM Platform provides core identity related functionalities such as:
		- User registration and login
		- JWT-based authentication
		- Role-based authorization
		- Token management (access/refresh)`,
	)
	dsl.Version("0.1.0")

	// Server block defines where the services are hosted.
	dsl.Server("iam", func() {
		dsl.Description("Primary server hosting the IAM service")

		// List of services exposed on this server.
		dsl.Services("auth")

		// Host represents a deployment environment (e.g., dev, staging, prod).
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
