package service

import (
	"context"

	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
)

// Login authenticates the user credentials and returns a token on success.
func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	return nil, nil
}
