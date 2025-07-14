package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// AuthService defines the authentication and authorization service interface.
var _ = dsl.Service("auth", func() {
	dsl.Description("The auth service handles user registration, authentication, and token issuance.")

	// Common domain level error types.
	rootdesign.CommonErrors()

	// Specific service level errors.
	dsl.Error("invalid_credentials", rootdesign.UnauthorizedError, "Invalid email or password")
	dsl.Error("password_mismatch", rootdesign.ValidationError, "Password and confirmation password do not match")
	dsl.Error("email_exists", rootdesign.ConflictError, "Email address is already registered")
	dsl.Error("user_not_found", rootdesign.NotFoundError, "User account not found")
	dsl.Error("invalid_token", rootdesign.UnauthorizedError, "Invalid or expired token")
	dsl.Error("session_expired", rootdesign.UnauthorizedError, "Session has expired")

	// Base path for the auth service.
	dsl.HTTP(func() {
		dsl.Path("/auth")
	})

	// --- Method: signup ---
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

	// --- Method: signin ---
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

	// --- Method: signout ---
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
