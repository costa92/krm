package biz

import "github.com/google/wire"

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface{}

type biz struct{}

// NewBiz returns a pointer to a new instance of the biz struct.
func NewBiz() *biz {
	return &biz{}
}
