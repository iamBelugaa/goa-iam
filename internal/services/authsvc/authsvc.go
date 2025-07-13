package authsvc

import (
	"context"
	"fmt"

	"github.com/iamBelugaa/goa-iam/gen/auth/gen/auth"
	userstore "github.com/iamBelugaa/goa-iam/internal/services/usersvc/store"
	"github.com/iamBelugaa/goa-iam/pkg/logger"
)

type authService struct {
	log       *logger.Logger
	userStore userstore.UserStorer
}

func NewService(log *logger.Logger, userStore userstore.UserStorer) *authService {
	return &authService{log: log, userStore: userStore}
}

func (s *authService) Signup(ctx context.Context, req *auth.SignupRequest) (*auth.SignupResponse, error) {
	if user, err := s.userStore.QueryUserByEmail(ctx, req.Email); err != nil {
		return nil, auth.MakeEmailExists(err)
	} else if user != nil {
		return nil, auth.MakeEmailExists(fmt.Errorf("user with email %s already exists", req.Email))
	}

	return &auth.SignupResponse{Success: true, Message: "", Data: ""}, nil
}

func (s *authService) Signin(ctx context.Context, req *auth.SigninRequest) (*auth.TokenResponse, error) {
	if user, err := s.userStore.QueryUserByEmail(ctx, req.Email); err != nil {
		return nil, auth.MakeNotFound(err)
	} else if user == nil {
		return nil, auth.MakeNotFound(fmt.Errorf("user with email %s doesn't exists", req.Email))
	}

	return &auth.TokenResponse{}, nil
}

func (s *authService) Signout(context.Context, *auth.SignoutRequest) (res *auth.SignoutResponse, err error) {
	return &auth.SignoutResponse{}, nil
}
