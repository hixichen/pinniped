// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.pinniped.dev/generated/1.19/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOIDCProviderConfigs implements OIDCProviderConfigInterface
type FakeOIDCProviderConfigs struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var oidcproviderconfigsResource = schema.GroupVersionResource{Group: "config.pinniped.dev", Version: "v1alpha1", Resource: "oidcproviderconfigs"}

var oidcproviderconfigsKind = schema.GroupVersionKind{Group: "config.pinniped.dev", Version: "v1alpha1", Kind: "OIDCProviderConfig"}

// Get takes name of the oIDCProviderConfig, and returns the corresponding oIDCProviderConfig object, and an error if there is any.
func (c *FakeOIDCProviderConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OIDCProviderConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(oidcproviderconfigsResource, c.ns, name), &v1alpha1.OIDCProviderConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCProviderConfig), err
}

// List takes label and field selectors, and returns the list of OIDCProviderConfigs that match those selectors.
func (c *FakeOIDCProviderConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OIDCProviderConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(oidcproviderconfigsResource, oidcproviderconfigsKind, c.ns, opts), &v1alpha1.OIDCProviderConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OIDCProviderConfigList{ListMeta: obj.(*v1alpha1.OIDCProviderConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.OIDCProviderConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oIDCProviderConfigs.
func (c *FakeOIDCProviderConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(oidcproviderconfigsResource, c.ns, opts))

}

// Create takes the representation of a oIDCProviderConfig and creates it.  Returns the server's representation of the oIDCProviderConfig, and an error, if there is any.
func (c *FakeOIDCProviderConfigs) Create(ctx context.Context, oIDCProviderConfig *v1alpha1.OIDCProviderConfig, opts v1.CreateOptions) (result *v1alpha1.OIDCProviderConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(oidcproviderconfigsResource, c.ns, oIDCProviderConfig), &v1alpha1.OIDCProviderConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCProviderConfig), err
}

// Update takes the representation of a oIDCProviderConfig and updates it. Returns the server's representation of the oIDCProviderConfig, and an error, if there is any.
func (c *FakeOIDCProviderConfigs) Update(ctx context.Context, oIDCProviderConfig *v1alpha1.OIDCProviderConfig, opts v1.UpdateOptions) (result *v1alpha1.OIDCProviderConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(oidcproviderconfigsResource, c.ns, oIDCProviderConfig), &v1alpha1.OIDCProviderConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCProviderConfig), err
}

// Delete takes name of the oIDCProviderConfig and deletes it. Returns an error if one occurs.
func (c *FakeOIDCProviderConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(oidcproviderconfigsResource, c.ns, name), &v1alpha1.OIDCProviderConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOIDCProviderConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(oidcproviderconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OIDCProviderConfigList{})
	return err
}

// Patch applies the patch and returns the patched oIDCProviderConfig.
func (c *FakeOIDCProviderConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OIDCProviderConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(oidcproviderconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OIDCProviderConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OIDCProviderConfig), err
}
