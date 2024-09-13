package v1beta1

import (
	v1beta1 "github.com/costa92/krm/pkg/apis/apps/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
)

type ChainLister interface {
	List(selector labels.Selector) (ret []*v1beta1.Chain, err error)
}
