//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ConfigMapsClusterGetter has a method to return a ConfigMapsClusterInterface.
// A group's cluster client should implement this interface.
type ConfigMapsClusterGetter interface {
	ConfigMaps() ConfigMapsClusterInterface
}

// ConfigMapsClusterInterface can operate on ConfigMaps across all clusters,
// or scope down to one cluster and return a ConfigMapsNamespacer.
type ConfigMapsClusterInterface interface {
	Cluster(logicalcluster.Name) ConfigMapsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type configMapsClusterInterface struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *configMapsClusterInterface) Cluster(name logicalcluster.Name) ConfigMapsNamespacer {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &configMapsNamespacer{clientCache: c.clientCache, name: name}
}

// List returns the entire collection of all ConfigMaps across all clusters.
func (c *configMapsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ConfigMaps(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all ConfigMaps across all clusters.
func (c *configMapsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ConfigMaps(metav1.NamespaceAll).Watch(ctx, opts)
}

// ConfigMapsNamespacer can scope to objects within a namespace, returning a corev1client.ConfigMapInterface.
type ConfigMapsNamespacer interface {
	Namespace(string) corev1client.ConfigMapInterface
}

type configMapsNamespacer struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
	name        logicalcluster.Name
}

func (n *configMapsNamespacer) Namespace(namespace string) corev1client.ConfigMapInterface {
	return n.clientCache.ClusterOrDie(n.name).ConfigMaps(namespace)
}
