// Package server provides initialization and lifecycle management
// for the HTTP server that hosts the Goa generated services.
package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	goahttp "goa.design/goa/v3/http"

	genauth "github.com/iamBelugaa/goa-iam/gen/auth/gen/auth"
	genauthserver "github.com/iamBelugaa/goa-iam/gen/auth/gen/http/auth/server"
	genuserserver "github.com/iamBelugaa/goa-iam/gen/user/gen/http/user/server"
	genuser "github.com/iamBelugaa/goa-iam/gen/user/gen/user"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/internal/services/authsvc"
	"github.com/iamBelugaa/goa-iam/internal/services/usersvc"
	usermemorystore "github.com/iamBelugaa/goa-iam/internal/services/usersvc/store/memory"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

// server encapsulates the application configuration,
// logger, HTTP server instance, and error channel.
type server struct {
	cfg         *config.Config // Application configuration
	log         *logger.Logger // Application logger
	httpServer  *http.Server   // Underlying HTTP server
	serverError chan error     // Channel for capturing async server errors
}

// New creates and configures a new instance of the server.
// It sets up user and auth services, mounts their HTTP handlers,
// and initializes the HTTP server.
func New(logger *logger.Logger, cfg *config.Config) *server {
	// Initialize in-memory user store and user service.
	userStore := usermemorystore.NewMemoryStore()
	userSvc := usersvc.NewService(userStore)
	userEndpoints := genuser.NewEndpoints(userSvc)

	// Initialize auth service using user store and configuration.
	authsvc := authsvc.NewService(logger, userStore, cfg.Auth)
	authEndPoints := genauth.NewEndpoints(authsvc)

	// Create Goa HTTP multiplexer.
	mux := goahttp.NewMuxer()

	// Setup and mount user HTTP handlers.
	userHandlers := genuserserver.New(userEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	genuserserver.Mount(mux, userHandlers)

	// Setup and mount auth HTTP handlers.
	authHandlers := genauthserver.New(authEndPoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	genauthserver.Mount(mux, authHandlers)

	// Log mounted user endpoints.
	for _, mount := range userHandlers.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	// Log mounted auth endpoints.
	for _, mount := range authHandlers.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	return &server{
		cfg:         cfg,
		log:         logger,
		serverError: make(chan error, 1),
		httpServer: &http.Server{
			Handler:      mux,
			IdleTimeout:  cfg.Server.IdleTimeout,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		},
	}
}

// ListenAndServe starts the HTTP server.
func (s *server) ListenAndServe() error {
	go func() {
		s.log.Infow("starting http server", "address", fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.serverError <- err
		}
	}()
	return nil
}

// ListenAndServeTLS is for starting a TLS enabled server.
func (s *server) ListenAndServeTLS() error {
	return nil
}

// Shutdown listens for termination signals or server errors
// and performs a graceful shutdown of the HTTP server.
func (s *server) Shutdown() error {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	// Handle server startup error.
	case err := <-s.serverError:
		s.log.Infow("received server error", "error", err)
		return fmt.Errorf("server error: %w", err)

	// Handle OS shutdown signal.
	case sig := <-signalCh:
		s.log.Infow("shutting down server signal received", "signal", sig)
		s.log.Infow("initiating graceful shutdown", "service", s.cfg.Application.Service)

		// Create context with timeout for graceful shutdown.
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.Server.ShutdownTimeout)
		defer cancel()

		// Attempt graceful shutdown.
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}

		s.log.Infow("graceful shutdown completed successfully", "service", s.cfg.Application.Service)
	}

	return nil
}
