package logger_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"

	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("Helpers", func() {
	Context("UserID", func() {
		It("should return a zap.Field with correct key and value", func() {
			userID := "user:123"
			field := logger.UserID(userID)

			Expect(field.Key).To(Equal("user_id"))
			Expect(field.Type).To(Equal(zapcore.StringType))
			Expect(field.String).To(Equal(userID))
		})
	})

	Context("RequestID", func() {
		It("should return a zap.Field with correct key and value", func() {
			requestID := "request:123"
			field := logger.RequestID(requestID)

			Expect(field.Key).To(Equal("request_id"))
			Expect(field.Type).To(Equal(zapcore.StringType))
			Expect(field.String).To(Equal(requestID))
		})
	})

	Context("RequestID", func() {
		It("should return a zap.Field with correct key and value", func() {
			operation := "read:users"
			field := logger.Operation(operation)

			Expect(field.Key).To(Equal("operation"))
			Expect(field.Type).To(Equal(zapcore.StringType))
			Expect(field.String).To(Equal(operation))
		})
	})
})
