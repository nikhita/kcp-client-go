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

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	"k8s.io/client-go/rest"

	kcpnodev1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/node/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpnodev1alpha1.NodeV1alpha1ClusterInterface = (*NodeV1alpha1ClusterClient)(nil)

type NodeV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *NodeV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) nodev1alpha1.NodeV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &NodeV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *NodeV1alpha1ClusterClient) RuntimeClasses() kcpnodev1alpha1.RuntimeClassClusterInterface {
	return &runtimeClassesClusterClient{Fake: c.Fake}
}

var _ nodev1alpha1.NodeV1alpha1Interface = (*NodeV1alpha1Client)(nil)

type NodeV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *NodeV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *NodeV1alpha1Client) RuntimeClasses() nodev1alpha1.RuntimeClassInterface {
	return &runtimeClassesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
