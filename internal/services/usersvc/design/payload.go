package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

// ListUsersResponse represents the structure of a list of all users.
var ListUsersResponse = dsl.Type("ListUsersResponse", func() {
	dsl.Description("Response returned when listing all users.")

	dsl.Reference(rootdesign.SuccessResponse)

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

	dsl.Reference(rootdesign.SuccessResponse)

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

	dsl.Reference(rootdesign.SuccessResponse)

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
