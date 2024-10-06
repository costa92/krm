// Copyright 2024 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/krm.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1beta1 "github.com/costa92/krm/pkg/apis/apps/v1beta1"
	appsv1beta1 "github.com/costa92/krm/pkg/generated/applyconfigurations/apps/v1beta1"
	scheme "github.com/costa92/krm/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChainsGetter has a method to return a ChainInterface.
// A group's client should implement this interface.
type ChainsGetter interface {
	Chains(namespace string) ChainInterface
}

// ChainInterface has methods to work with Chain resources.
type ChainInterface interface {
	Create(ctx context.Context, chain *v1beta1.Chain, opts v1.CreateOptions) (*v1beta1.Chain, error)
	Update(ctx context.Context, chain *v1beta1.Chain, opts v1.UpdateOptions) (*v1beta1.Chain, error)
	UpdateStatus(ctx context.Context, chain *v1beta1.Chain, opts v1.UpdateOptions) (*v1beta1.Chain, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Chain, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ChainList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Chain, err error)
	Apply(ctx context.Context, chain *appsv1beta1.ChainApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Chain, err error)
	ApplyStatus(ctx context.Context, chain *appsv1beta1.ChainApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Chain, err error)
	ChainExpansion
}

// chains implements ChainInterface
type chains struct {
	client rest.Interface
	ns     string
}

// newChains returns a Chains
func newChains(c *AppsV1beta1Client, namespace string) *chains {
	return &chains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chain, and returns the corresponding chain object, and an error if there is any.
func (c *chains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Chain, err error) {
	result = &v1beta1.Chain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Chains that match those selectors.
func (c *chains) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ChainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ChainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chains.
func (c *chains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a chain and creates it.  Returns the server's representation of the chain, and an error, if there is any.
func (c *chains) Create(ctx context.Context, chain *v1beta1.Chain, opts v1.CreateOptions) (result *v1beta1.Chain, err error) {
	result = &v1beta1.Chain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(chain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a chain and updates it. Returns the server's representation of the chain, and an error, if there is any.
func (c *chains) Update(ctx context.Context, chain *v1beta1.Chain, opts v1.UpdateOptions) (result *v1beta1.Chain, err error) {
	result = &v1beta1.Chain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chains").
		Name(chain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(chain).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *chains) UpdateStatus(ctx context.Context, chain *v1beta1.Chain, opts v1.UpdateOptions) (result *v1beta1.Chain, err error) {
	result = &v1beta1.Chain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chains").
		Name(chain.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(chain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the chain and deletes it. Returns an error if one occurs.
func (c *chains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched chain.
func (c *chains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Chain, err error) {
	result = &v1beta1.Chain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied chain.
func (c *chains) Apply(ctx context.Context, chain *appsv1beta1.ChainApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Chain, err error) {
	if chain == nil {
		return nil, fmt.Errorf("chain provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(chain)
	if err != nil {
		return nil, err
	}
	name := chain.Name
	if name == nil {
		return nil, fmt.Errorf("chain.Name must be provided to Apply")
	}
	result = &v1beta1.Chain{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("chains").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *chains) ApplyStatus(ctx context.Context, chain *appsv1beta1.ChainApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Chain, err error) {
	if chain == nil {
		return nil, fmt.Errorf("chain provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(chain)
	if err != nil {
		return nil, err
	}

	name := chain.Name
	if name == nil {
		return nil, fmt.Errorf("chain.Name must be provided to Apply")
	}

	result = &v1beta1.Chain{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("chains").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
