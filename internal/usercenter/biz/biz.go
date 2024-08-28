package biz

import (
	"github.com/google/wire"

	"github.com/costa92/krm/internal/usercenter/auth"
	"github.com/costa92/krm/internal/usercenter/store"
	"github.com/costa92/krm/pkg/authn"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz any

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
