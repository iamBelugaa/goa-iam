package design

import (
	"goa.design/goa/v3/dsl"
)

var User = dsl.Type("User", func() {
	dsl.Attribute("id", dsl.String, "User Id", func() {
		dsl.Format(dsl.FormatUUID)
	})

	dsl.Attribute("firstName", dsl.String, "First Name", func() {
		dsl.Example("John")
	})

	dsl.Attribute("lastName", dsl.String, "Last Name", func() {
		dsl.Example("Doe")
	})
})
