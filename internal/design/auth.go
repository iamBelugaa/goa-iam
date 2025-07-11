package design

import (
	"goa.design/goa/v3/dsl"
)

// SignupRequest defines the structure of the payload sent to the signup endpoint.
var SignupRequest = dsl.Type("SignupRequest", func() {
	dsl.Description("Payload for user signup. Includes user identity and authentication credentials.")

	dsl.Attribute("firstName", dsl.String, "User's first name", func() {
		dsl.MinLength(3)
		dsl.Example("John")
	})

	dsl.Attribute("lastName", dsl.String, "User's last name", func() {
		dsl.MinLength(2)
		dsl.Example("Doe")
	})

	dsl.Attribute("email", dsl.String, "User's email address", func() {
		dsl.Format(dsl.FormatEmail)
		dsl.Example("work", "john@work.com")
		dsl.Example("personal", "john@gmail.com")
	})

	dsl.Attribute("password", dsl.String, "User's password (8-32 characters)", func() {
		dsl.MinLength(8)
		dsl.MaxLength(32)
		dsl.Example("secure-password")
	})

	dsl.Required("firstName", "lastName", "email", "password")
})

// SignupResponse defines the structure of the response returned after successful signup.
var SignupResponse = dsl.Type("SignupResponse", func() {
	dsl.Description("Response returned after user signup indicating success and the associated email.")

	dsl.Attribute("success", dsl.Boolean, "Indicates if the signup was successful")
	dsl.Attribute("email", dsl.String, "Email address of the signed-up user", func() {
		dsl.Format(dsl.FormatEmail)
		dsl.Example("john@work.com")
	})

	dsl.Required("success", "email")
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

// TokenResponse defines the JWT token structure returned upon successful signin.
var TokenResponse = dsl.Type("TokenResponse", func() {
	dsl.Description("JWT tokens and metadata issued upon successful authentication.")

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

	dsl.Attribute("expiresIn", dsl.Int, "Expiration time (in seconds) for the access token", func() {
		dsl.Minimum(60)
		dsl.Example(3600)
	})

	dsl.Required("accessToken", "refreshToken", "tokenType", "expiresIn")
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

// AuthService defines the authentication and authorization service interface.
var _ = dsl.Service("auth", func() {
	dsl.Description("The auth service handles user registration, authentication, and token issuance.")

	// Common errors across the auth endpoints.
	dsl.Error("conflict", dsl.ErrorResult, "The email is already registered")
	dsl.Error("invalid_credentials", dsl.ErrorResult, "Incorrect email or password")
	dsl.Error("unauthorized", dsl.ErrorResult, "User is not authenticated")
	dsl.Error("invalid_token", dsl.ErrorResult, "Provided token is invalid or expired")

	// Base path for the auth service.
	dsl.HTTP(func() {
		dsl.Path("/auth")
	})

	// Signup endpoint.
	dsl.Method("signup", func() {
		dsl.Description("Registers a new user with provided credentials and personal info.")

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

	// Signin endpoint.
	dsl.Method("signin", func() {
		dsl.Description("Authenticates a user and returns a JWT access and refresh token.")

		dsl.Payload(SigninRequest)
		dsl.Result(TokenResponse)
		dsl.Error("invalid_credentials")

		dsl.HTTP(func() {
			dsl.POST("/signin")

			dsl.Body(func() {
				dsl.Attribute("email")
				dsl.Attribute("password")
			})

			dsl.Response("invalid_credentials", dsl.StatusUnauthorized)
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(TokenResponse)
			})
		})
	})

	// Signout endpoint.
	dsl.Method("signout", func() {
		dsl.Description("Logs out an authenticated user by invalidating the access or refresh token.")
		dsl.Security(JWTAuth)

		dsl.Payload(SignoutRequest)
		dsl.Result(dsl.Empty)

		dsl.Error("unauthorized")
		dsl.Error("invalid_token")

		dsl.HTTP(func() {
			dsl.POST("/signout")
			dsl.Response(dsl.StatusNoContent)
			dsl.Response("unauthorized", dsl.StatusUnauthorized)
			dsl.Response("invalid_token", dsl.StatusUnauthorized)
		})
	})
})
