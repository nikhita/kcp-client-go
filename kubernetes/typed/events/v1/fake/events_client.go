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

	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	"k8s.io/client-go/rest"

	kcpeventsv1 "github.com/kcp-dev/client-go/kubernetes/typed/events/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpeventsv1.EventsV1ClusterInterface = (*EventsV1ClusterClient)(nil)

type EventsV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *EventsV1ClusterClient) Cluster(clusterPath logicalcluster.Path) eventsv1.EventsV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &EventsV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *EventsV1ClusterClient) Events() kcpeventsv1.EventClusterInterface {
	return &eventsClusterClient{Fake: c.Fake}
}

var _ eventsv1.EventsV1Interface = (*EventsV1Client)(nil)

type EventsV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *EventsV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *EventsV1Client) Events(namespace string) eventsv1.EventInterface {
	return &eventsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
