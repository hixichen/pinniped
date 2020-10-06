// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	configv1alpha1 "go.pinniped.dev/generated/1.18/apis/config/v1alpha1"
	versioned "go.pinniped.dev/generated/1.18/client/clientset/versioned"
	internalinterfaces "go.pinniped.dev/generated/1.18/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.pinniped.dev/generated/1.18/client/listers/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OIDCProviderConfigInformer provides access to a shared informer and lister for
// OIDCProviderConfigs.
type OIDCProviderConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OIDCProviderConfigLister
}

type oIDCProviderConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOIDCProviderConfigInformer constructs a new informer for OIDCProviderConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOIDCProviderConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOIDCProviderConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOIDCProviderConfigInformer constructs a new informer for OIDCProviderConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOIDCProviderConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().OIDCProviderConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().OIDCProviderConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&configv1alpha1.OIDCProviderConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *oIDCProviderConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOIDCProviderConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *oIDCProviderConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha1.OIDCProviderConfig{}, f.defaultInformer)
}

func (f *oIDCProviderConfigInformer) Lister() v1alpha1.OIDCProviderConfigLister {
	return v1alpha1.NewOIDCProviderConfigLister(f.Informer().GetIndexer())
}
