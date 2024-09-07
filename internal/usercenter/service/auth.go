package service

import (
	"context"
	"github.com/costa92/krm/pkg/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
)

// Login authenticates the user credentials and returns a token on success.
func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	log.C(ctx).Infow("Login", "username", rq.Username)
	return nil, nil
}

// Logout invalidates the user token.
func (s *UserCenterService) Logout(ctx context.Context, rq *v1.LogoutRequest) (*emptypb.Empty, error) {
	log.Infof("Logout")
	return nil, nil
}
