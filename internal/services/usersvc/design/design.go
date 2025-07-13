package design

import (
	rootdesign "github.com/iamBelugaa/goa-iam/internal/design"
	"goa.design/goa/v3/dsl"
)

var _ = dsl.Service("user", func() {
	dsl.Description("User service handles user related queries")

	// Common errors across the user endpoints.
	rootdesign.CommonErrors()

	// Base path for the user service.
	dsl.HTTP(func() {
		dsl.Path("/users")
	})

	dsl.Method("list", func() {
		dsl.Description("Get all users")

		dsl.Payload(dsl.Empty)
		dsl.Result(func() {
			dsl.Attribute("users", dsl.ArrayOf(rootdesign.User), "List of users")
		})

		dsl.HTTP(func() {
			dsl.GET("/")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(func() {
					dsl.Attribute("users", dsl.ArrayOf(rootdesign.User), "List of users")
				})
			})
		})
	})
})
