/*
Copyright 2022 The KCP Authors.

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

package dynamic

import (
	"context"
	"fmt"
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"

	thirdpartydynamic "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/dynamic"
)

var _ ClusterInterface = (*ClusterClientset)(nil)

type ClusterClientset struct {
	clientCache kcpclient.Cache[dynamic.Interface]
}

// Cluster scopes the client down to a particular cluster.
func (c *ClusterClientset) Cluster(name logicalcluster.Name) dynamic.Interface {
	return c.clientCache.ClusterOrDie(name)
}

func (c *ClusterClientset) Resource(resource schema.GroupVersionResource) ResourceClusterInterface {
	return &ClusterResourceClient{clientCache: c.clientCache, resource: resource}
}

// NewForConfig creates a new ClusterClientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ClusterClientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new ClusterClientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*ClusterClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	cache := kcpclient.NewCache(c, httpClient, &kcpclient.Constructor[dynamic.Interface]{
		NewForConfigAndClient: thirdpartydynamic.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.New("root")); err != nil {
		return nil, err
	}

	return &ClusterClientset{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new ClusterClientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterClientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

type ClusterResourceClient struct {
	clientCache kcpclient.Cache[dynamic.Interface]
	resource    schema.GroupVersionResource
}

// Cluster scopes the client down to a particular cluster.
func (c *ClusterResourceClient) Cluster(name logicalcluster.Name) dynamic.NamespaceableResourceInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).Resource(c.resource)
}

// List returns the entire collection of all resources across all clusters.
func (c *ClusterResourceClient) List(ctx context.Context, opts metav1.ListOptions) (*unstructured.UnstructuredList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Resource(c.resource).List(ctx, opts)
}

// Watch begins to watch all resources across all clusters.
func (c *ClusterResourceClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Resource(c.resource).Watch(ctx, opts)
}
