package biz

//go:generate mockgen -self_package github.com/costa92/krm/internal/usercenter/biz -destination mock_biz.go -package biz github.com/costa92/krm/internal/usercenter/biz IBiz

import (
	"github.com/costa92/krm/internal/usercenter/biz/user"
	"github.com/google/wire"

	"github.com/costa92/krm/internal/usercenter/auth"
	"github.com/costa92/krm/internal/usercenter/store"
	"github.com/costa92/krm/pkg/authn"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds    store.IStore
	authn authn.Authenticator
	auth  auth.AuthProvider
}

// NewBiz returns a pointer to a new instance of the biz struct.
func NewBiz(ds store.IStore, authn authn.Authenticator, auth auth.AuthProvider) *biz {
	return &biz{
		ds:    ds,
		authn: authn,
		auth:  auth,
	}
}

// Users returns a new instance of the UserBiz interface.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
