package userstore

import (
	"context"

	"github.com/iamBelugaa/goa-iam/gen/user/gen/user"
)

type UserStorer interface {
	QueryById(context context.Context, userID string) (*user.User, error)
	QueryByEmail(context context.Context, email string) (*user.User, error)
	Create(context context.Context, cmd *user.CreateUserRequest) (*user.User, error)
}
