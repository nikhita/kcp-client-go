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

package v1alpha1

import (
	"github.com/kcp-dev/logicalcluster/v2"

	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	"k8s.io/client-go/rest"

	kcpinternalv1alpha1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/apiserverinternal/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpinternalv1alpha1.InternalV1alpha1ClusterInterface = (*InternalV1alpha1ClusterClient)(nil)

type InternalV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *InternalV1alpha1ClusterClient) Cluster(cluster logicalcluster.Name) internalv1alpha1.InternalV1alpha1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &InternalV1alpha1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *InternalV1alpha1ClusterClient) StorageVersions() kcpinternalv1alpha1.StorageVersionClusterInterface {
	return &storageVersionsClusterClient{Fake: c.Fake}
}

var _ internalv1alpha1.InternalV1alpha1Interface = (*InternalV1alpha1Client)(nil)

type InternalV1alpha1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *InternalV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *InternalV1alpha1Client) StorageVersions() internalv1alpha1.StorageVersionInterface {
	return &storageVersionsClient{Fake: c.Fake, Cluster: c.Cluster}
}
