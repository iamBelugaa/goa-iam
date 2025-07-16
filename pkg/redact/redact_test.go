package redact_test

import (
	"testing"

	"github.com/iamBelugaa/goa-iam/pkg/redact"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redact Suite")
}

var _ = Describe("Redact", func() {
	Describe("RedactEmail", func() {
		It("should redact the user part of the email", func() {
			output := redact.RedactEmail("john@doe.com")
			Expect(output).To(Equal("j***@doe.com"))
		})

		It("should work for empty string", func() {
			output := redact.RedactEmail("")
			Expect(output).To(Equal("[REDACTED]"))
		})

		It("should work for invalid email address", func() {
			output := redact.RedactEmail("JohnDoe.com")
			Expect(output).To(Equal("[REDACTED]"))
		})
	})

	Describe("RedactSensitiveData", func() {
		It("should redact sensitive data correctly", func() {
			output := redact.RedactSensitiveData("1234567890")
			Expect(output).To(Equal("12***90"))
		})

		It("should work for shorter data", func() {
			output := redact.RedactSensitiveData("123")
			Expect(output).To(Equal("[REDACTED]"))
		})
	})
})
