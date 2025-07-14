// Package userstore defines the interface for interacting with the user data storage layer.
package userstore

import (
	"context"

	"github.com/iamBelugaa/goa-iam/gen/user/gen/user"
)

// UserStorer defines the contract for managing user data in a storage backend.
type UserStorer interface {
	// QueryById retrieves a user by their unique user ID.
	QueryById(ctx context.Context, userID string) (*user.User, error)

	// QueryByEmail retrieves a user by their email address.
	QueryByEmail(ctx context.Context, email string) (*user.User, error)

	// Create stores a new user in the storage backend using the provided request payload.
	Create(ctx context.Context, cmd *user.CreateUserRequest) (*user.User, error)
}
