// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm
//

package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	clientset "github.com/costa92/krm/pkg/generated/clientset/versioned"
	"github.com/costa92/krm/pkg/generated/informers"
)

// WantsInternalInformerFactory defines a function which sets InformerFactory for admission plugins that need it.
type WantsInternalInformerFactory interface {
	admission.InitializationValidator
	SetInternalInformerFactory(informers.SharedInformerFactory)
}

// WantsInternalClientSet defines a function which sets external ClientSet for admission plugins that need it.
type WantsInternalClientSet interface {
	admission.InitializationValidator
	SetInternalClientSet(clientset.Interface)
}
