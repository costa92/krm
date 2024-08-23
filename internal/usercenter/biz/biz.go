package biz

import (
	"github.com/costa92/krm/internal/usercenter/auth"
	"github.com/costa92/krm/internal/usercenter/store"
	"github.com/costa92/krm/pkg/authn"
	"github.com/google/wire"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface{}

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
