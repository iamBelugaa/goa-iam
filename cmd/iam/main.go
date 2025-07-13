package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error : %v", err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to parse env : %v", err)
	}

	logger, err := logger.NewWithConfig(
		cfg.Application.Service,
		cfg.Application.Version,
		cfg.Application.Environment, cfg.Logging,
	)
	if err != nil {
		return fmt.Errorf("failed to construct logger : %v", err)
	}

	logger.LogRequest("Request 1", http.MethodGet, "/api/users", "", "", "", http.StatusOK)
	logger.LogRequest("Request 2", http.MethodPost, "/api/users", "", "", "", http.StatusOK)

	if err := logger.Close(); err != nil {
		return fmt.Errorf("failed to flush buffered log entries : %v", err)
	}

	return nil
}
