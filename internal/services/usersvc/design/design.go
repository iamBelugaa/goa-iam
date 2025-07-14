// Package design contains the service design definitions for Goa.
// This file defines the "user" service, including its methods, errors, and HTTP interface.
package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// UserService defines the user management endpoints.
var _ = dsl.Service("user", func() {
	dsl.Description("User management service for CRUD operations on users, roles, and permissions.")

	// Common domain level error types.
	rootdesign.CommonErrors()

	// Specific service level errors.
	dsl.Error("email_exists", rootdesign.ConflictError, "Email address is already registered.")
	dsl.Error("user_not_found", rootdesign.NotFoundError, "User account not found.")

	// Base URL path for all HTTP endpoints in the user service.
	dsl.HTTP(func() {
		dsl.Path("/users")
	})

	// --- Method: list ---
	dsl.Method("list", func() {
		dsl.Description("List all users in the system.")

		dsl.Payload(dsl.Empty)
		dsl.Result(ListUsersResponse)
		dsl.Error("internal_server_error")

		dsl.HTTP(func() {
			dsl.GET("/")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(ListUsersResponse)
			})
		})
	})

	// --- Method: getById ---
	dsl.Method("getById", func() {
		dsl.Description("Retrieve a user by their unique ID.")

		dsl.Error("user_not_found")
		dsl.Payload(GetUserByIDRequest)
		dsl.Result(GetUserByIDResponse)

		dsl.HTTP(func() {
			dsl.GET("/{id}")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(GetUserByIDResponse)
			})
		})
	})

	// --- Method: create ---
	dsl.Method("create", func() {
		dsl.Description("Create a new user account in the system.")

		dsl.Error("email_exists")
		dsl.Payload(CreateUserRequest)
		dsl.Result(CreateUserResponse)

		dsl.HTTP(func() {
			dsl.POST("/")

			dsl.Body(func() {
				dsl.Attribute("firstName", dsl.String, "User's first name")
				dsl.Attribute("lastName", dsl.String, "User's last name")
				dsl.Attribute("email", dsl.String, "Email used for login")
				dsl.Attribute("password", dsl.String, "Password (plain text, will be hashed)")
			})

			dsl.Response(dsl.StatusCreated, func() {
				dsl.Body(CreateUserResponse)
			})
		})
	})
})
