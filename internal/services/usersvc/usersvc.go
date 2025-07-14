// Package usersvc provides user-related business logic for the IAM system.
package usersvc

import (
	"context"
	"fmt"

	genuser "github.com/iamBelugaa/goa-iam/gen/user/gen/user"
	userstore "github.com/iamBelugaa/goa-iam/internal/services/usersvc/store"
)

// service implements user-related operations backed by a user store.
type service struct {
	store userstore.UserStorer // Interface to the underlying user storage
}

// NewService creates a new user service instance with the provided store.
func NewService(userStore userstore.UserStorer) *service {
	return &service{store: userStore}
}

// List returns all users in the system.
func (s *service) List(ctx context.Context) (*genuser.ListUsersResponse, error) {
	users, err := s.store.List(ctx)
	if err != nil {
		return nil, genuser.MakeInternalServerError(err)
	}

	return &genuser.ListUsersResponse{
		Success: true,
		Data:    users,
		Message: "User's list fetched successfully",
	}, nil
}

// GetByID retrieves a user by their unique ID.
func (s *service) GetByID(ctx context.Context, req *genuser.GetUserByIDPayload) (*genuser.GetUserByIDResponse, error) {
	user, err := s.store.QueryById(ctx, req.ID)
	if err != nil {
		return nil, genuser.MakeUserNotFound(err)
	}
	if user == nil {
		return nil, genuser.MakeUserNotFound(fmt.Errorf("user with id %s doesn't exist", req.ID))
	}

	return &genuser.GetUserByIDResponse{
		Success: true,
		Data:    user,
		Message: "User fetched successfully",
	}, nil
}

// Create registers a new user in the system.
func (s *service) Create(ctx context.Context, req *genuser.CreateUserRequest) (*genuser.CreateUserResponse, error) {
	user, err := s.store.Create(ctx, req)
	if err != nil {
		return nil, genuser.MakeEmailExists(err)
	}

	return &genuser.CreateUserResponse{
		Success: true,
		Data:    user,
		Message: "User created successfully",
	}, nil
}
