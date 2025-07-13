package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// SignupRequest defines the structure of the payload sent to the signup endpoint.
var SignupRequest = dsl.Type("SignupRequest", func() {
	dsl.Description("Payload for user signup. Includes user identity and authentication credentials.")

	dsl.Attribute("firstName", dsl.String, "User's first name", func() {
		dsl.MinLength(1)
		dsl.MaxLength(50)
		dsl.Pattern("^[a-zA-Z\\s]+$")
		dsl.Example("John")
	})

	dsl.Attribute("lastName", dsl.String, "User's last name", func() {
		dsl.MinLength(1)
		dsl.MaxLength(50)
		dsl.Pattern("^[a-zA-Z\\s]+$")
		dsl.Example("Doe")
	})

	dsl.Attribute("email", dsl.String, "User's email address", func() {
		dsl.Format(dsl.FormatEmail)
		dsl.Example("work", "john@work.com")
		dsl.Example("personal", "john@gmail.com")
	})

	dsl.Attribute("password", dsl.String, "User's password (8-32 characters)", func() {
		dsl.MinLength(8)
		dsl.MaxLength(128)
		dsl.Example("secure-password")
	})

	dsl.Attribute("confirmPassword", dsl.String, "Password confirmation (must match password)", func() {
		dsl.MinLength(8)
		dsl.MaxLength(128)
		dsl.Example("secure-password")
	})

	dsl.Required("firstName", "lastName", "email", "password", "confirmPassword")
})

// SignupResponse defines the structure of the response returned after successful signup.
var SignupResponse = dsl.Type("SignupResponse", func() {
	dsl.Description("Response returned after user signup indicating success and the associated email.")
	dsl.Extend(rootdesign.SuccessResponse)
})

// SigninRequest defines the structure of the payload sent to the signin endpoint.
var SigninRequest = dsl.Type("SigninRequest", func() {
	dsl.Description("Payload for user signin. Contains email and password.")

	dsl.Attribute("email", dsl.String, "User's email address", func() {
		dsl.Format(dsl.FormatEmail)
		dsl.Example("user@example.com")
	})

	dsl.Attribute("password", dsl.String, "User's password", func() {
		dsl.MinLength(8)
		dsl.Example("secure-password")
	})

	dsl.Required("email", "password")
})

// TokenPayload defines the structure of the JWT tokens.
var TokenPayload = dsl.Type("TokenPayload", func() {
	dsl.Description("JWT access and refresh tokens issued after successful authentication.")

	dsl.Attribute("accessToken", dsl.String, "JWT access token", func() {
		dsl.Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
	})

	dsl.Attribute("refreshToken", dsl.String, "JWT refresh token", func() {
		dsl.Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
	})

	dsl.Attribute("tokenType", dsl.String, "Type of token issued", func() {
		dsl.Enum("Bearer")
		dsl.Example("Bearer")
	})

	dsl.Attribute("expiresIn", dsl.Int, "Expiration time in seconds for the access token", func() {
		dsl.Minimum(60)
		dsl.Example(3600)
	})

	dsl.Required("accessToken", "refreshToken", "tokenType", "expiresIn")
})

// TokenResponse defines the JWT tokens returned after successful authentication.
var TokenResponse = dsl.Type("TokenResponse", func() {
	dsl.Description("Response containing access and refresh tokens after successful signin.")
	dsl.Reference(rootdesign.SuccessResponse)

	dsl.Attribute("success")
	dsl.Attribute("message")
	dsl.Attribute("data", TokenPayload)
})

// SignoutRequest defines the payload for the signout endpoint.
var SignoutRequest = dsl.Type("SignoutRequest", func() {
	dsl.Description("Payload for user signout.")

	dsl.Token("token", dsl.String, "JWT token from Authorization header", func() {
		dsl.Description("JWT access token extracted from Authorization: Bearer <token>")
		dsl.Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
	})

	dsl.Required("token")
})

// SignoutResponse defines the response returned after a successful user signout operation.
var SignoutResponse = dsl.Type("SignoutResponse", func() {
	dsl.Description("Response indicating that the user has been signed out successfully.")
	dsl.Extend(rootdesign.SuccessResponse)
})
