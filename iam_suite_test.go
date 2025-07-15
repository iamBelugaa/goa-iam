package goa_iam_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoaIam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IAM Suite")
}
