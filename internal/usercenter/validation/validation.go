// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

package validation

import (
	"github.com/google/wire"
)

// ProviderSet is validator providers.
var ProviderSet = wire.NewSet(New, wire.Bind(new(any), new(*validator)))

// validator struct implements the custom validator interface.
type validator struct{}

// New creates and initializes a custom validator.
// It receives an instance of store.IStore interface as parameter ds
// and returns a new *validator and an error.
func New() (*validator, error) {
	vd := &validator{}

	return vd, nil
}
