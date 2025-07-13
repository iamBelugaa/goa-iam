package userstore

import (
	"context"
	"fmt"

	"github.com/iamBelugaa/goa-iam/gen/user/gen/user"
)

type inMemoryStore struct {
	emailToIdMap map[string]string
	users        map[string]*user.User
}

func NewInMemoryStore() *inMemoryStore {
	return &inMemoryStore{}
}

func (in *inMemoryStore) QueryUserById(context context.Context, userID string) (*user.User, error) {
	user, ok := in.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with id %s doesn't exists", userID)
	}
	return user, nil
}

func (in *inMemoryStore) QueryUserByEmail(context context.Context, email string) (*user.User, error) {
	userID, ok := in.emailToIdMap[email]
	if !ok {
		return nil, fmt.Errorf("user with email %s doesn't exists", email)
	}

	user, ok := in.users[userID]
	if !ok {
		return nil, fmt.Errorf("user with email %s doesn't exists", email)
	}

	return user, nil
}
