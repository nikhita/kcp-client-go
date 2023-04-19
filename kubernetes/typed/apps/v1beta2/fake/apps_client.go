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

	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	"k8s.io/client-go/rest"

	kcpappsv1beta2 "github.com/kcp-dev/client-go/kubernetes/typed/apps/v1beta2"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpappsv1beta2.AppsV1beta2ClusterInterface = (*AppsV1beta2ClusterClient)(nil)

type AppsV1beta2ClusterClient struct {
	*kcptesting.Fake
}

func (c *AppsV1beta2ClusterClient) Cluster(clusterPath logicalcluster.Path) appsv1beta2.AppsV1beta2Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AppsV1beta2Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AppsV1beta2ClusterClient) StatefulSets() kcpappsv1beta2.StatefulSetClusterInterface {
	return &statefulSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1beta2ClusterClient) Deployments() kcpappsv1beta2.DeploymentClusterInterface {
	return &deploymentsClusterClient{Fake: c.Fake}
}

func (c *AppsV1beta2ClusterClient) DaemonSets() kcpappsv1beta2.DaemonSetClusterInterface {
	return &daemonSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1beta2ClusterClient) ReplicaSets() kcpappsv1beta2.ReplicaSetClusterInterface {
	return &replicaSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1beta2ClusterClient) ControllerRevisions() kcpappsv1beta2.ControllerRevisionClusterInterface {
	return &controllerRevisionsClusterClient{Fake: c.Fake}
}

var _ appsv1beta2.AppsV1beta2Interface = (*AppsV1beta2Client)(nil)

type AppsV1beta2Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AppsV1beta2Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AppsV1beta2Client) StatefulSets(namespace string) appsv1beta2.StatefulSetInterface {
	return &statefulSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1beta2Client) Deployments(namespace string) appsv1beta2.DeploymentInterface {
	return &deploymentsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1beta2Client) DaemonSets(namespace string) appsv1beta2.DaemonSetInterface {
	return &daemonSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1beta2Client) ReplicaSets(namespace string) appsv1beta2.ReplicaSetInterface {
	return &replicaSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1beta2Client) ControllerRevisions(namespace string) appsv1beta2.ControllerRevisionInterface {
	return &controllerRevisionsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
