package main

import (
	"log"
	"net/http"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

func main() {
	logger, err := logger.NewWithConfig(
		"iam", "1.0.0",
		config.EnvironmentDevelopment,
		&config.Logging{
			RequestLogging:  true,
			RedactSensitive: true,
			Level:           "info",
		},
	)
	if err != nil {
		log.Fatalln("failed to construct logger", err)
	}

	logger.LogRequest("Hello World", http.MethodGet, "/api", "", "", "", http.StatusOK)
	logger.LogRequest("Hello World", http.MethodGet, "/api", "", "", "", http.StatusOK)
	logger.LogRequest("Hello World", http.MethodGet, "/api", "", "", "", http.StatusOK)
	logger.LogRequest("Hello World", http.MethodGet, "/api", "", "", "", http.StatusOK)
	logger.LogRequest("Hello World", http.MethodGet, "/api", "", "", "", http.StatusOK)

	if err := logger.Close(); err != nil {
		log.Fatalln("failed to flush buffered log entries", err)
	}
}
