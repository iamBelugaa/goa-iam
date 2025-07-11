package design

import (
	"goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = dsl.API("iam-platform", func() {
	dsl.Title("Identity and Access Management Platform")
	dsl.Description("A IAM service for authentication, authorization and user management")
	dsl.Version("0.1.0")

	// Define server configuration.
	dsl.Server("iam", func() {
		dsl.Description("IAM service server")

		// Services define the different functional areas of our API.
		dsl.Services("auth")

		// Host defines where the server runs.
		dsl.Host("localhost", func() {
			dsl.Description("Development server")
			dsl.URI("http://localhost:8080")
		})
	})

	// Global prefix for all services.
	dsl.HTTP(func() {
		dsl.Path("/api/v1")
	})
})
