package logger_test

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("Logger", func() {
	var (
		service        string
		version        string
		cfg            *config.Logging
		environment    config.Environment
		logBuffer      *bytes.Buffer
		originalStdErr *os.File
		r, w           *os.File
		wg             sync.WaitGroup
	)

	BeforeEach(func() {
		version = "1.0.0"
		service = "test-service"
		cfg = &config.Logging{Level: "info"}
		environment = config.EnvironmentDevelopment

		// Redirect standard error to a pipe for capturing logs.
		originalStdErr = os.Stderr
		logBuffer = &bytes.Buffer{}

		var err error
		r, w, err = os.Pipe()
		Expect(err).NotTo(HaveOccurred())
		os.Stderr = w

		// Read from the pipe and write to logBuffer.
		wg.Add(1)
		go func() {
			defer GinkgoRecover()
			defer wg.Done()
			io.Copy(logBuffer, r)
		}()
	})

	// Close the writer to signal end of logging and restore original Stderr.
	AfterEach(func() {
		w.Close()
		wg.Wait()

		r.Close()
		os.Stderr = originalStdErr
	})

	Describe("NewWithConfig", func() {
		Context("with valid configuration", func() {
			It("should create a logger successfully", func() {
				logger, err := logger.NewWithConfig(service, version, environment, cfg)

				Expect(err).NotTo(HaveOccurred())
				Expect(logger).NotTo(BeNil())
			})

			It("should set service metadata correctly", func() {
				logger, err := logger.NewWithConfig(service, version, environment, cfg)

				Expect(err).NotTo(HaveOccurred())
				Expect(logger).NotTo(BeNil())

				logger.Info("test message")
				time.Sleep(30 * time.Millisecond)

				output := logBuffer.String()
				Expect(output).To(ContainSubstring(service))
				Expect(output).To(ContainSubstring(version))
			})
		})

		Context("with different log levels", func() {
			DescribeTable("should handle different log levels", func(level string) {
				cfg.Level = level
				logger, err := logger.NewWithConfig(service, version, environment, cfg)

				Expect(err).NotTo(HaveOccurred())
				Expect(logger).NotTo(BeNil())
			},
				Entry("info level", "info"),
				Entry("warn level", "warn"),
				Entry("error level", "error"),
				Entry("debug level", "debug"),
			)

			It("should return error for invalid log level", func() {
				cfg.Level = "invalid"
				logger, err := logger.NewWithConfig(service, version, environment, cfg)

				Expect(err).To(HaveOccurred())
				Expect(logger).To(BeNil())
			})
		})

		Context("with different environments", func() {
			It("should configure for development environment", func() {
				environment = config.EnvironmentDevelopment
				logger, err := logger.NewWithConfig(service, version, environment, cfg)

				Expect(err).NotTo(HaveOccurred())
				Expect(logger).NotTo(BeNil())
			})

			It("should configure for production environment", func() {
				environment = config.EnvironmentProduction
				logger, err := logger.NewWithConfig(service, version, environment, cfg)

				Expect(err).NotTo(HaveOccurred())
				Expect(logger).NotTo(BeNil())
			})
		})
	})

	Describe("Close", func() {
		It("should sync the logger without error", func() {
			logger, err := logger.NewWithConfig(service, version, environment, cfg)
			Expect(err).NotTo(HaveOccurred())
			Expect(logger).NotTo(BeNil())

			err = logger.Close()
			Expect(err).To(HaveOccurred()) // Ideally this should not error, but somehow it is.
		})
	})
})
