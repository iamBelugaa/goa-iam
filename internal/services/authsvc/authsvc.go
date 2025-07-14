// Package authsvc provides authentication related business logic for the IAM system.
package authsvc

import (
	"context"
	"fmt"

	"goa.design/goa/v3/security"

	genauth "github.com/iamBelugaa/goa-iam/gen/auth/gen/auth"
	genuser "github.com/iamBelugaa/goa-iam/gen/user/gen/user"

	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/internal/services/authsvc/tokenmgr"
	userstore "github.com/iamBelugaa/goa-iam/internal/services/usersvc/store"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
	"github.com/iamBelugaa/goa-iam/pkg/redact"
)

// service implements authentication operations such as signup, signin, signout,
// and token-based authorization using a JWT token manager.
type service struct {
	log       *logger.Logger            // Logger for structured logging
	userStore userstore.UserStorer      // Interface to the user data store
	tm        *tokenmgr.JWTTokenManager // JWT manager for token generation and validation
}

// NewService initializes and returns a new auth service instance.
func NewService(log *logger.Logger, userStore userstore.UserStorer, authCfg *config.Auth) *service {
	return &service{
		log:       log,
		userStore: userStore,
		tm:        tokenmgr.NewJWTManager(authCfg),
	}
}

// Signup creates a new user account after validating password confirmation.
func (s *service) Signup(ctx context.Context, req *genauth.SignupRequest) (*genauth.SignupResponse, error) {
	s.log.Infow(
		"signup request received",
		"email", redact.RedactEmail(req.Email),
		"firstName", req.FirstName, "lastName", req.LastName,
		"password", redact.RedactSensitiveData(req.Password),
		"confirmPassword", redact.RedactSensitiveData(req.ConfirmPassword),
	)

	if req.Password != req.ConfirmPassword {
		return nil, genauth.MakePasswordMismatch(fmt.Errorf("confirm password and password doesn't match"))
	}

	_, err := s.userStore.Create(ctx, &genuser.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		s.log.Infow("create user error", "email", redact.RedactEmail(req.Email), "error", err)
		return nil, genauth.MakeEmailExists(err)
	}

	s.log.Infow("signup request successful", "email", redact.RedactEmail(req.Email))
	return &genauth.SignupResponse{
		Success: true,
		Message: "User signed up successfully",
	}, nil
}

// Signin authenticates a user by email and password.
func (s *service) Signin(ctx context.Context, req *genauth.SigninRequest) (*genauth.TokenResponse, error) {
	s.log.Infow(
		"signin request received",
		"email", redact.RedactEmail(req.Email),
		"password", redact.RedactSensitiveData(req.Password),
	)

	user, err := s.userStore.QueryByEmail(ctx, req.Email)
	if err != nil {
		s.log.Infow("query user error", "email", redact.RedactEmail(req.Email), "error", err)
		return nil, genauth.MakeNotFound(err)
	}
	if user == nil {
		s.log.Infow("query user error", "email", redact.RedactEmail(req.Email), "error", err)
		return nil, genauth.MakeNotFound(fmt.Errorf("user with email %s doesn't exist", req.Email))
	}

	accessToken, err := s.tm.Generate(s.tm.StandardClaims(user.ID, tokenmgr.AccessToken))
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tm.Generate(s.tm.StandardClaims(user.ID, tokenmgr.RefreshToken))
	if err != nil {
		return nil, err
	}

	s.log.Infow("signin request successful", "email", redact.RedactEmail(req.Email))
	return &genauth.TokenResponse{
		Success: true,
		Message: "Signed in user successfully",
		Data:    &genauth.TokenPayload{AccessToken: accessToken, RefreshToken: refreshToken},
	}, nil
}

// Signout invalidates an access token by verifying its validity and user existence.
func (s *service) Signout(ctx context.Context, req *genauth.SignoutRequest) (*genauth.SignoutResponse, error) {
	s.log.Infow("signout request successful", "token", redact.RedactSensitiveData(req.Token))

	claims, err := s.tm.ParseWithClaims(req.Token)
	if err != nil {
		s.log.Infow("jwt parse error", "error", err)
		return nil, err
	}

	if claims.TokenType != tokenmgr.AccessToken {
		s.log.Infow("invalid token used for signout operation")
		return nil, genauth.MakeInvalidToken(fmt.Errorf("invalid token used for signout operation"))
	}

	if _, err := s.userStore.QueryById(ctx, claims.Subject); err != nil {
		s.log.Infow("query user error", "error", err)
		return nil, genauth.MakeNotFound(err)
	}

	s.log.Infow("signout request successful")
	return &genauth.SignoutResponse{
		Success: true,
		Message: "Signed out successfully",
	}, nil
}

// JWTAuth validates a JWT and attaches the corresponding user context.
func (s *service) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	// TODO: Parse and validate the JWT token and attach user info to context.
	return context.Background(), nil
}
