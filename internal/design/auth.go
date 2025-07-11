package design

import (
	"goa.design/goa/v3/dsl"
)

var SignupRequest = dsl.Type("SignupRequest", func() {
	dsl.Description("Signup request payload structure")

	dsl.Attribute("firstName", dsl.String, "First Name", func() {
		dsl.MinLength(3)
		dsl.Example("john")
	})

	dsl.Attribute("lastName", dsl.String, "Last Name", func() {
		dsl.MinLength(2)
		dsl.Example("doe")
	})

	dsl.Attribute("email", dsl.String, "Email Address", func() {
		dsl.Format(dsl.FormatEmail)
		dsl.Example("work", "john@work.com")
		dsl.Example("personal", "john@gmail.com")
	})

	dsl.Attribute("password", dsl.String, "User Password", func() {
		dsl.MinLength(8)
		dsl.MaxLength(32)
		dsl.Example("secure-password")
	})

	dsl.Required("firstName", "lastName", "email", "password")
})

var SignupResponse = dsl.Type("SignupResponse", func() {
	dsl.Description("Signup response payload structure")

	dsl.Attribute("success", dsl.Boolean, "Indicates success or false")

	dsl.Attribute("email", dsl.String, "User Email Address", func() {
		dsl.Format(dsl.FormatEmail)
		dsl.Example("work", "john@work.com")
		dsl.Example("personal", "john@gmail.com")
	})

	dsl.Required("success", "email")
})

// AuthService defines the authentication and authorization endpoints.
var _ = dsl.Service("auth", func() {
	dsl.Description("Authentication and Authorization service")

	// Define common error responses that can occur across all auth endpoints.
	dsl.Error("conflict", dsl.ErrorResult, "Email Conflict")
	dsl.Error("unauthorized", dsl.ErrorResult, "Unauthorized")
	dsl.Error("invalid_credentials", dsl.ErrorResult, "Invalid Credentials")

	// HTTP transport configuration for the auth service.
	dsl.HTTP(func() {
		dsl.Path("/auth")
	})

	dsl.Method("signup", func() {
		dsl.Description("Signup user endpoint")

		dsl.Payload(SignupRequest)
		dsl.Result(SignupResponse)
		dsl.Error("conflict")

		dsl.HTTP(func() {
			dsl.POST("/signup")

			dsl.Body(func() {
				dsl.Attribute("firstName")
				dsl.Attribute("lastName")
				dsl.Attribute("email")
				dsl.Attribute("password")
			})

			dsl.Response("conflict", dsl.StatusConflict)
			dsl.Response(dsl.StatusCreated, func() {
				dsl.Body(SignupResponse)
			})
		})
	})
})
