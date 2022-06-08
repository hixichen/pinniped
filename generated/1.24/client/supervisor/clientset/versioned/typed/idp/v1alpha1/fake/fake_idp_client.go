// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "go.pinniped.dev/generated/1.24/client/supervisor/clientset/versioned/typed/idp/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIDPV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIDPV1alpha1) ActiveDirectoryIdentityProviders(namespace string) v1alpha1.ActiveDirectoryIdentityProviderInterface {
	return &FakeActiveDirectoryIdentityProviders{c, namespace}
}

func (c *FakeIDPV1alpha1) LDAPIdentityProviders(namespace string) v1alpha1.LDAPIdentityProviderInterface {
	return &FakeLDAPIdentityProviders{c, namespace}
}

func (c *FakeIDPV1alpha1) OIDCIdentityProviders(namespace string) v1alpha1.OIDCIdentityProviderInterface {
	return &FakeOIDCIdentityProviders{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIDPV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
