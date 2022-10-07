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

package v1beta1

import (
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/apimachinery/pkg/runtime/schema"
	policyv1beta1client "k8s.io/client-go/kubernetes/typed/policy/v1beta1"

	kcppolicyv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/policy/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var evictionsResource = schema.GroupVersionResource{Group: "policy", Version: "V1beta1", Resource: "evictions"}
var evictionsKind = schema.GroupVersionKind{Group: "policy", Version: "V1beta1", Kind: "Eviction"}

type evictionsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *evictionsClusterClient) Cluster(cluster logicalcluster.Name) kcppolicyv1beta1.EvictionsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &evictionsNamespacer{Fake: c.Fake, Cluster: cluster}
}

type evictionsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *evictionsNamespacer) Namespace(namespace string) policyv1beta1client.EvictionInterface {
	return &evictionsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type evictionsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}
