package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func getEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return intVal
}

func getEnvBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return boolVal
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	durationVal, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return durationVal
}

func getEnvSlice(key string, fallback []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	parts := strings.Split(val, ",")
	if len(parts) < 2 {
		return fallback
	}

	return parts
}

func ToEnvironment(str string) Environment {
	switch strings.ToLower(str) {
	case "prod", "production":
		return EnvironmentProduction
	case "dev", "develop", "development", "local":
		return EnvironmentDevelopment
	default:
		return EnvironmentDevelopment
	}
}
