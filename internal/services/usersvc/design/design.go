package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// UserService defines the user management endpoints.
// This service handles CRUD operations for users, roles, and permissions.
var _ = dsl.Service("user", func() {
	dsl.Description("User management service for CRUD operations on users, roles, and permissions")

	// Applies common error types.
	rootdesign.CommonErrors()
	dsl.Error("email_exists", rootdesign.ConflictError, "Email address is already registered")
	dsl.Error("user_not_found", rootdesign.NotFoundError, "User account not found")

	// Base path for the user service.
	dsl.HTTP(func() {
		dsl.Path("/users")
	})

	// List all users.
	dsl.Method("list", func() {
		dsl.Description("List all users in the system.")

		dsl.Payload(dsl.Empty)
		dsl.Result(ListUsersResponse)

		dsl.HTTP(func() {
			dsl.GET("/")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(ListUsersResponse)
			})
		})
	})

	// Get user by ID.
	dsl.Method("getById", func() {
		dsl.Description("Retrieve a user by their ID.")

		dsl.Payload(GetUserByIDRequest)
		dsl.Result(GetUserByIDResponse)

		dsl.Error("not_found")

		dsl.HTTP(func() {
			dsl.GET("/{id}")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(GetUserByIDResponse)
			})
		})
	})

	dsl.Method("create", func() {
		dsl.Description("Create a new user account")

		dsl.Payload(CreateUserRequest)
		dsl.Result(CreateUserResponse)

		dsl.Error("email_exists")

		dsl.HTTP(func() {
			dsl.POST("/")
			dsl.Body(func() {
				dsl.Attribute("firstName")
				dsl.Attribute("lastName")
				dsl.Attribute("email")
				dsl.Attribute("password")
			})

			dsl.Response(dsl.StatusCreated, func() {
				dsl.Body(CreateUserResponse)
			})
		})
	})
})
