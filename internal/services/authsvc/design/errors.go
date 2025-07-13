package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// AuthErrors defines errors specific to authentication operations.
func AuthErrors() {
	// Invalid credentials provided during signin.
	dsl.Error("invalid_credentials", rootdesign.UnauthorizedError, "Invalid email or password")

	// Invalid confirm password during signup.
	dsl.Error("password_mismatch", rootdesign.ValidationError, "Password and confirmation password do not match")

	// Email already exists during registration.
	dsl.Error("email_exists", rootdesign.ConflictError, "Email address is already registered")

	// User account not found
	dsl.Error("user_not_found", rootdesign.NotFoundError, "User account not found")

	// JWT token is invalid or malformed.
	dsl.Error("invalid_token", rootdesign.UnauthorizedError, "Invalid or expired token")

	// Session has expired.
	dsl.Error("session_expired", rootdesign.UnauthorizedError, "Session has expired")
}
