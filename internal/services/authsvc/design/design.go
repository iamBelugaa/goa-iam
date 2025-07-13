package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// AuthService defines the authentication and authorization service interface.
var _ = dsl.Service("auth", func() {
	dsl.Description("The auth service handles user registration, authentication, and token issuance.")

	// Common errors across the auth endpoints.
	rootdesign.CommonErrors()
	AuthErrors()

	// Base path for the auth service.
	dsl.HTTP(func() {
		dsl.Path("/auth")
	})

	// Signup endpoint.
	dsl.Method("signup", func() {
		dsl.Description("Registers a new user with provided credentials and personal info.")

		dsl.Payload(SignupRequest)
		dsl.Result(SignupResponse)

		// Possible errors
		dsl.Error("email_exists")
		dsl.Error("validation_failed")
		dsl.Error("password_mismatch")

		dsl.HTTP(func() {
			dsl.POST("/signup")
			dsl.Body(func() {
				dsl.Attribute("firstName")
				dsl.Attribute("lastName")
				dsl.Attribute("email")
				dsl.Attribute("password")
			})

			dsl.Response(dsl.StatusCreated, func() {
				dsl.Body(SignupResponse)
				dsl.Description("User registered successfully")
			})
		})
	})

	// Signin endpoint.
	dsl.Method("signin", func() {
		dsl.Description("Authenticates a user and returns a JWT access and refresh token.")

		dsl.Payload(SigninRequest)
		dsl.Result(TokenResponse)

		dsl.Error("not_found")
		dsl.Error("invalid_credentials")

		dsl.HTTP(func() {
			dsl.POST("/signin")
			dsl.Body(func() {
				dsl.Attribute("email")
				dsl.Attribute("password")
			})

			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(TokenResponse)
			})
		})
	})

	// Signout endpoint.
	dsl.Method("signout", func() {
		dsl.Description("Logs out an authenticated user by invalidating the access or refresh token.")
		dsl.Security(rootdesign.JWTAuth)

		dsl.Payload(SignoutRequest)
		dsl.Result(SignoutResponse)

		dsl.Error("not_found")
		dsl.Error("unauthorized")
		dsl.Error("invalid_token")

		dsl.HTTP(func() {
			dsl.POST("/signout")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(SignoutResponse)
			})
		})
	})
})
