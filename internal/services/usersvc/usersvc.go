// Package usersvc provides user-related business logic for the IAM system.
package usersvc

import (
	"context"
	"fmt"

	genuser "github.com/iamBelugaa/goa-iam/gen/user"

	userstore "github.com/iamBelugaa/goa-iam/internal/services/usersvc/store"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
	"github.com/iamBelugaa/goa-iam/pkg/redact"
)

// service implements user-related operations backed by a user store.
type service struct {
	log   *logger.Logger       // Logger for structured logging
	store userstore.UserStorer // Interface to the underlying user storage
}

// NewService creates a new user service instance with the provided store.
func NewService(log *logger.Logger, userStore userstore.UserStorer) *service {
	return &service{store: userStore, log: log}
}

// List returns all users in the system.
func (s *service) List(ctx context.Context) (*genuser.ListUsersResponse, error) {
	s.log.Infow("list users request received")

	users, err := s.store.List(ctx)
	if err != nil {
		s.log.Infow("list users error", "error", err)
		return nil, genuser.MakeInternalServerError(err)
	}

	s.log.Infow("list users request successful", "totalUsers", len(users))
	return &genuser.ListUsersResponse{
		Success: true,
		Data:    users,
		Message: "User's list fetched successfully",
	}, nil
}

// GetByID retrieves a user by their unique ID.
func (s *service) GetByID(ctx context.Context, req *genuser.GetUserByIDPayload) (*genuser.GetUserByIDResponse, error) {
	s.log.Infow("getUserById request received", "userId", req.ID)

	user, err := s.store.QueryById(ctx, req.ID)
	if err != nil {
		s.log.Infow("getUserById error", "userId", req.ID, "error", err)
		return nil, genuser.MakeUserNotFound(err)
	}
	if user == nil {
		s.log.Infow("getUserById error", "userId", req.ID, "error", err)
		return nil, genuser.MakeUserNotFound(fmt.Errorf("user with id %s doesn't exist", req.ID))
	}

	s.log.Infow("getUserById request successful", "user", *user)
	return &genuser.GetUserByIDResponse{
		Success: true,
		Data:    user,
		Message: "User fetched successfully",
	}, nil
}

// Create registers a new user in the system.
func (s *service) Create(ctx context.Context, req *genuser.CreateUserRequest) (*genuser.CreateUserResponse, error) {
	s.log.Infow(
		"create user request received",
		"email", redact.RedactEmail(req.Email),
		"firstName", req.FirstName, "lastName", req.LastName,
		"password", redact.RedactSensitiveData(req.Password),
	)

	user, err := s.store.Create(ctx, req)
	if err != nil {
		s.log.Infow("create user error", "email", redact.RedactEmail(req.Email), "error", err)
		return nil, genuser.MakeEmailExists(err)
	}

	s.log.Infow("create user request successful", "user", *user)
	return &genuser.CreateUserResponse{
		Success: true,
		Data:    user,
		Message: "User created successfully",
	}, nil
}
