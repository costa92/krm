package user

//go:generate mockgen -self_package github.com/costa92/krm//internal/usercenter/biz/user -destination mock_user.go -package user github.com/costa92/krm/internal/usercenter/biz/user UserBiz

import (
	"context"
	"github.com/costa92/krm/internal/usercenter/store"
	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
)

type UserBiz interface {
	Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.UserReply, error)
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

// New returns a new instance of userBiz.
func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}
