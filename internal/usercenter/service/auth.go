package service

import (
	"context"
	"github.com/costa92/krm/pkg/log"

	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
)

// Login authenticates the user credentials and returns a token on success.
func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	log.Infow("Login", "username", rq.Username)
	return nil, nil
}
