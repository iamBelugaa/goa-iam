package main

import (
	"log"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

func main() {
	logger, err := logger.NewWithConfig("", "", config.EnvironmentDevelopment, &config.Logging{})
	if err != nil {
		log.Fatalln("failed to construct logger", err)
	}

	if err := logger.Close(); err != nil {
		log.Fatalln("failed to flush buffered log entries", err)
	}
}
