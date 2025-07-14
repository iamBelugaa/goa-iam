package usersvc

import (
	"context"
	"fmt"

	genuser "github.com/iamBelugaa/goa-iam/gen/user/gen/user"
	userstore "github.com/iamBelugaa/goa-iam/internal/services/usersvc/store"
)

type service struct {
	store userstore.UserStorer
}

func NewService(userStore userstore.UserStorer) *service {
	return &service{store: userStore}
}

// List all users in the system.
func (s *service) List(ctx context.Context) (*genuser.ListUsersResponse, error) {
	return nil, nil
}

// Retrieve a user by their ID.
func (s *service) GetByID(ctx context.Context, req *genuser.GetUserByIDPayload) (*genuser.GetUserByIDResponse, error) {
	user, err := s.store.QueryById(ctx, req.ID)
	if err != nil {
		return nil, genuser.MakeNotFound(err)
	}
	if user == nil {
		return nil, genuser.MakeNotFound(fmt.Errorf("user with id %s doesn't exists", req.ID))
	}

	return &genuser.GetUserByIDResponse{Success: true, Data: user, Message: "User fetch successfully"}, nil
}

// Create a new user account.
func (s *service) Create(ctx context.Context, req *genuser.CreateUserRequest) (*genuser.CreateUserResponse, error) {
	user, err := s.store.Create(ctx, req)
	if err != nil {
		return nil, genuser.MakeEmailExists(err)
	}
	return &genuser.CreateUserResponse{Success: true, Data: user, Message: "User created successfully"}, nil
}
