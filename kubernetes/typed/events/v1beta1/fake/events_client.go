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

	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	"k8s.io/client-go/rest"

	kcpeventsv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/events/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpeventsv1beta1.EventsV1beta1ClusterInterface = (*EventsV1beta1ClusterClient)(nil)

type EventsV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *EventsV1beta1ClusterClient) Cluster(cluster logicalcluster.Name) eventsv1beta1.EventsV1beta1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &EventsV1beta1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *EventsV1beta1ClusterClient) Events() kcpeventsv1beta1.EventClusterInterface {
	return &eventsClusterClient{Fake: c.Fake}
}

var _ eventsv1beta1.EventsV1beta1Interface = (*EventsV1beta1Client)(nil)

type EventsV1beta1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *EventsV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *EventsV1beta1Client) Events(namespace string) eventsv1beta1.EventInterface {
	return &eventsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}