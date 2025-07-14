package main

import (
	"fmt"
	"log"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/internal/server"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error : %v", err)
	}
}

// run loads configuration, initializes the logger and HTTP server,
// and starts listening for HTTP requests.
func run() (err error) {
	// Load configuration from environment variables.
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to parse env : %v", err)
	}

	// Initialize structured logger with configuration.
	logger, err := logger.NewWithConfig(
		cfg.Application.Service, cfg.Application.Version, cfg.Application.Environment, cfg.Logging,
	)
	if err != nil {
		return fmt.Errorf("failed to construct logger : %v", err)
	}

	// Ensure logs are flushed on exit.
	defer func() {
		if loggerCloseErr := logger.Close(); loggerCloseErr != nil {
			err = fmt.Errorf("failed to flush buffered log entries : %v", err)
		}
	}()

	// Initialize the HTTP server with configuration and logger.
	server := server.New(logger, cfg)

	// Start serving HTTP requests.
	server.ListenAndServe()

	// Wait for shutdown signal or error and gracefully shut down the server.
	return server.Shutdown()
}
