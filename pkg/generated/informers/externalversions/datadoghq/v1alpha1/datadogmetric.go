// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	datadoghqv1alpha1 "github.com/DataDog/datadog-operator/pkg/apis/datadoghq/v1alpha1"
	versioned "github.com/DataDog/datadog-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/DataDog/datadog-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/DataDog/datadog-operator/pkg/generated/listers/datadoghq/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DatadogMetricInformer provides access to a shared informer and lister for
// DatadogMetrics.
type DatadogMetricInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DatadogMetricLister
}

type datadogMetricInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDatadogMetricInformer constructs a new informer for DatadogMetric type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatadogMetricInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatadogMetricInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDatadogMetricInformer constructs a new informer for DatadogMetric type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatadogMetricInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DatadoghqV1alpha1().DatadogMetrics(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DatadoghqV1alpha1().DatadogMetrics(namespace).Watch(options)
			},
		},
		&datadoghqv1alpha1.DatadogMetric{},
		resyncPeriod,
		indexers,
	)
}

func (f *datadogMetricInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatadogMetricInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *datadogMetricInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&datadoghqv1alpha1.DatadogMetric{}, f.defaultInformer)
}

func (f *datadogMetricInformer) Lister() v1alpha1.DatadogMetricLister {
	return v1alpha1.NewDatadogMetricLister(f.Informer().GetIndexer())
}
