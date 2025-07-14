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
