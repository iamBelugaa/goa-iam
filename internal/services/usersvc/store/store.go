package userstore

import (
	"context"

	"github.com/iamBelugaa/goa-iam/gen/user/gen/user"
)

type UserStorer interface {
	QueryUserById(context context.Context, userID string) (*user.User, error)
	QueryUserByEmail(context context.Context, email string) (*user.User, error)
}
