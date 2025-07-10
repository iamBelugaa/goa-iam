// Package config defines configuration types used across the IAM service.
package config

import "time"

// Environment defines the type for representing different runtime environments.
type Environment string

// Supported environment constants.
var (
	EnvironmentProduction  Environment = "PRODUCTION"
	EnvironmentDevelopment Environment = "DEVELOPMENT"
)

// Logging holds settings for how logging should behave in different environments.
type Logging struct {
	Level           string `json:"level"`
	RequestLogging  bool   `json:"requestLogging"`
	RedactSensitive bool   `json:"redactSensitive"`
}

// Server holds HTTP server configuration.
type Server struct {
	TLS             *TLS          `json:"tls"`
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	ReadTimeout     time.Duration `json:"readTimeout"`
	IdleTimeout     time.Duration `json:"idleTimeout"`
	WriteTimeout    time.Duration `json:"writeTimeout"`
	ShutdownTimeout time.Duration `json:"shutdownTimeout"`
}

// TLS holds TLS configuration.
type TLS struct {
	Enable   bool   `json:"enable"`
	KeyFile  string `json:"keyFile"`
	CertFile string `json:"certFile"`
}

// Cors holds CORS configuration.
type Cors struct {
	Enable         bool     `json:"enable"`
	AllowedOrigins []string `json:"allowedOrigins"`
	AllowedHeaders []string `json:"allowedHeaders"`
	AllowedMethods []string `json:"allowedMethods"`
	ExposedHeaders []string `json:"exposedHeaders"`
}

// App holds application specific configuration.
type Application struct {
	Cors        *Cors       `json:"cors"`
	Service     string      `json:"service"`
	Version     string      `json:"version"`
	Environment Environment `json:"environment"`
}

type Config struct {
	Server      *Server      `json:"server"`
	Logging     *Logging     `json:"logging"`
	Application *Application `json:"application"`
}

func Load() (*Config, error) {
	return &Config{
		Server: &Server{
			TLS: &TLS{
				Enable:   getEnvBool("TLS_ENABLED", false),
				KeyFile:  getEnv("TLS_KEY_FILE", ""),
				CertFile: getEnv("TLS_CERT_FILE", ""),
			},
			Host:            getEnv("SERVER_HOST", "0.0.0.0"),
			Port:            getEnvInt("SERVER_PORT", 8080),
			ReadTimeout:     getEnvDuration("SERVER_READ_TIMEOUT", time.Second*10),
			IdleTimeout:     getEnvDuration("SERVER_IDLE_TIMEOUT", time.Second*25),
			WriteTimeout:    getEnvDuration("SERVER_WRITE_TIMEOUT", time.Second*10),
			ShutdownTimeout: getEnvDuration("SERVER_SHUTDOWN_TIMEOUT", time.Second*30),
		},
		Logging: &Logging{
			Level:           getEnv("LOG_LEVEL", "INFO"),
			RequestLogging:  getEnvBool("LOG_REQUEST_LOGGING", true),
			RedactSensitive: getEnvBool("LOG_REDACT_SENSITIVE", true),
		},
		Application: &Application{
			Cors: &Cors{
				Enable:         getEnvBool("CORS_ENABLED", true),
				AllowedOrigins: getEnvSlice("ALLOWED_ORIGINS", []string{}),
				AllowedHeaders: getEnvSlice("ALLOWED_HEADERS", []string{}),
				AllowedMethods: getEnvSlice("ALLOWED_METHODS", []string{}),
				ExposedHeaders: getEnvSlice("EXPOSED_HEADERS", []string{}),
			},
			Version:     getEnv("APP_VERSION", "0.1.0"),
			Service:     getEnv("SERVICE_NAME", "IAM-PLATFORM"),
			Environment: ToEnvironment(getEnv("APP_ENVIRONMENT", "development")),
		},
	}, nil
}
