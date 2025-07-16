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
	Level string `json:"level"`
}

// Server holds HTTP server configuration.
type Server struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	ReadTimeout     time.Duration `json:"readTimeout"`
	IdleTimeout     time.Duration `json:"idleTimeout"`
	WriteTimeout    time.Duration `json:"writeTimeout"`
	ShutdownTimeout time.Duration `json:"shutdownTimeout"`
}

// App holds application specific configuration.
type Application struct {
	Service     string      `json:"service"`
	Version     string      `json:"version"`
	Environment Environment `json:"environment"`
}

type Auth struct {
	Issuer              string        `json:"issuer"`
	Secret              string        `json:"secret"`
	Audience            string        `json:"audience"`
	AccessTokenExpTime  time.Duration `json:"accessTokenExpTime"`
	RefreshTokenExpTime time.Duration `json:"refreshTokenExpTime"`
}

type Config struct {
	Server      *Server      `json:"server"`
	Auth        *Auth        `json:"auth"`
	Logging     *Logging     `json:"logging"`
	Application *Application `json:"application"`
}

func Load() (*Config, error) {
	return &Config{
		Server: &Server{
			Host:            getEnv("SERVER_HOST", "0.0.0.0"),
			Port:            getEnvInt("SERVER_PORT", 8080),
			ReadTimeout:     getEnvDuration("SERVER_READ_TIMEOUT", time.Second*10),
			IdleTimeout:     getEnvDuration("SERVER_IDLE_TIMEOUT", time.Second*25),
			WriteTimeout:    getEnvDuration("SERVER_WRITE_TIMEOUT", time.Second*10),
			ShutdownTimeout: getEnvDuration("SERVER_SHUTDOWN_TIMEOUT", time.Second*30),
		},
		Auth: &Auth{
			Audience:            getEnv("AUTH_AUDIENCE", "http://localhost:8080"),
			Issuer:              getEnv("AUTH_ISSUER", "https://issuer.iam.support"),
			AccessTokenExpTime:  getEnvDuration("AUTH_ACCESS_TOKEN_EXP_TIME", time.Hour),
			RefreshTokenExpTime: getEnvDuration("AUTH_REFRESH_TOKEN_EXP_TIME", time.Hour*24*60),
			Secret:              getEnv("AUTH_SECRET", "9916ce66f41d25276ab5923ce5e62ef7fbb6e046bb3072a507bf0362bae0d63d"),
		},
		Logging: &Logging{
			Level: getEnv("LOG_LEVEL", "INFO"),
		},
		Application: &Application{
			Version:     getEnv("APP_VERSION", "0.1.0"),
			Service:     getEnv("SERVICE_NAME", "IAM-PLATFORM"),
			Environment: ToEnvironment(getEnv("APP_ENVIRONMENT", "development")),
		},
	}, nil
}
