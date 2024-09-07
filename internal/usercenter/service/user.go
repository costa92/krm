package service

import (
	"context"
	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
	"github.com/costa92/krm/pkg/log"
)

// CreateUser

func (s *UserCenterService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.UserReply, error) {
	log.C(ctx).Infow("CreateUser", "req", req)
	return s.biz.Users().Create(ctx, req)
}
