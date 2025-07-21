package design

import (
	"goa.design/goa/v3/dsl"
)

// User represents a registered user in the system.
var User = dsl.Type("User", func() {
	dsl.Description("User defines the information of a system user, including identifying information, contact details, timestamps, and status.")

	dsl.Attribute("id", dsl.String, "User's unique identifier", func() {
		dsl.Format(dsl.FormatUUID)
		dsl.Example("4d2efde6-448a-4c26-a69a-26c2f9a6de4a")
	})

	dsl.Attribute("firstName", dsl.String, "User's first name", func() {
		dsl.Example("John")
	})

	dsl.Attribute("lastName", dsl.String, "User's last name", func() {
		dsl.Example("Doe")
	})

	dsl.Attribute("email", dsl.String, "User's email address", func() {
		dsl.Description("Primary email address used for communication and login.")
		dsl.Format(dsl.FormatEmail)
		dsl.Example("work", "john@work.com")
		dsl.Example("personal", "john@gmail.com")
	})

	dsl.Attribute("status", dsl.String, "Indicates current status of the user", func() {
		dsl.Example("active")
	})

	dsl.Attribute("createdAt", dsl.String, "Timestamp when the user was created", func() {
		dsl.Description("Timestamp representing when the user account was created.")
		dsl.Format(dsl.FormatDateTime)
		dsl.Example("2025-01-01T00:00:00Z")
	})

	dsl.Attribute("updatedAt", dsl.String, "Timestamp when the user was last updated", func() {
		dsl.Description("Timestamp representing the last time the user's data was modified.")
		dsl.Format(dsl.FormatDateTime)
		dsl.Example("2025-06-15T13:45:30Z")
	})

	dsl.Required("id", "firstName", "lastName", "email", "status", "createdAt", "updatedAt")
})

// ListUsersResponse represents the structure of a list of all users.
var ListUsersResponse = dsl.Type("ListUsersResponse", func() {
	dsl.Description("Response returned when listing all users.")

	dsl.Reference(SuccessResponse)

	dsl.Attribute("success", dsl.Boolean, "Whether the request was successful.")
	dsl.Attribute("message", dsl.String, "Human-readable message explaining the result.")
	dsl.Attribute("data", dsl.ArrayOf(User), "List of users returned in the response", func() {
		dsl.Example([]any{
			map[string]any{
				"id":        "4d2efde6-448a-4c26-a69a-26c2f9a6de4a",
				"firstName": "John",
				"lastName":  "Doe",
				"email":     "john@gmail.com",
				"status":    "active",
				"createdAt": "2025-01-01T00:00:00Z",
				"updatedAt": "2025-01-01T00:00:00Z",
			},
		})
	})

	dsl.Required("success", "message", "data")
})

// GetUserByIDRequest defines the payload to retrieve a single user by ID.
var GetUserByIDRequest = dsl.Type("GetUserByIDPayload", func() {
	dsl.Description("Payload for retrieving a user by unique identifier.")

	dsl.Attribute("id", dsl.String, "User's unique identifier", func() {
		dsl.Format(dsl.FormatUUID)
		dsl.Example("4d2efde6-448a-4c26-a69a-26c2f9a6de4a")
	})

	dsl.Required("id")
})

// GetUserByIDResponse defines the response when a user is retrieved by ID.
var GetUserByIDResponse = dsl.Type("GetUserByIDResponse", func() {
	dsl.Description("Response returned when retrieving a user by ID.")

	dsl.Reference(SuccessResponse)

	dsl.Attribute("success", dsl.Boolean, "Whether the request was successful.")
	dsl.Attribute("message", dsl.String, "Human-readable message explaining the result.")
	dsl.Attribute("data", User, "User returned in the response.")
})

// CreateUserRequest defines the payload for creating a new user.
var CreateUserRequest = dsl.Type("CreateUserRequest", func() {
	dsl.Description("Payload for user creation. Includes user identity and authentication credentials.")

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
		dsl.Example("john.doe@example.com")
	})

	dsl.Attribute("password", dsl.String, "User's password (8-128 characters)", func() {
		dsl.MinLength(8)
		dsl.MaxLength(128)
		dsl.Example("secure-password")
	})

	dsl.Required("firstName", "lastName", "email", "password")
})

// CreateUserResponse defines the structure of the response when a new user is created.
var CreateUserResponse = dsl.Type("CreateUserResponse", func() {
	dsl.Description("Response returned after successfully creating a user.")

	dsl.Reference(SuccessResponse)

	dsl.Attribute("success", dsl.Boolean, "Whether the request was successful.")
	dsl.Attribute("message", dsl.String, "Human-readable message explaining the result.")
	dsl.Attribute("data", User, "Details of the created user.", func() {
		dsl.Example(map[string]any{
			"id":        "4d2efde6-448a-4c26-a69a-26c2f9a6de4a",
			"firstName": "John",
			"lastName":  "Doe",
			"email":     "john@gmail.com",
			"status":    "active",
			"createdAt": "2025-01-01T00:00:00Z",
			"updatedAt": "2025-01-01T00:00:00Z",
		})
	})
})

// UserService defines the user management endpoints.
var _ = dsl.Service("user", func() {
	dsl.Description("User management service for CRUD operations on users, roles, and permissions.")

	// Common domain level error types.
	commonErrors()

	// Specific service level errors.
	dsl.Error("email_exists", ConflictError, "Email address is already registered.")
	dsl.Error("user_not_found", NotFoundError, "User account not found.")

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
