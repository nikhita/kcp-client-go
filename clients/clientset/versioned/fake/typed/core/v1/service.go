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
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"

	kcpcorev1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var servicesResource = schema.GroupVersionResource{Group: "core", Version: "V1", Resource: "services"}
var servicesKind = schema.GroupVersionKind{Group: "core", Version: "V1", Kind: "Service"}

type servicesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *servicesClusterClient) Cluster(cluster logicalcluster.Name) kcpcorev1.ServicesNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &servicesNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of Services that match those selectors across all clusters.
func (c *servicesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(servicesResource, servicesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.ServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceList{ListMeta: obj.(*corev1.ServiceList).ListMeta}
	for _, item := range obj.(*corev1.ServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Services across all clusters.
func (c *servicesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(servicesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type servicesNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *servicesNamespacer) Namespace(namespace string) corev1client.ServiceInterface {
	return &servicesClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type servicesClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *servicesClient) Create(ctx context.Context, service *corev1.Service, opts metav1.CreateOptions) (*corev1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(servicesResource, c.Cluster, c.Namespace, service), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}

func (c *servicesClient) Update(ctx context.Context, service *corev1.Service, opts metav1.UpdateOptions) (*corev1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(servicesResource, c.Cluster, c.Namespace, service), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}

func (c *servicesClient) UpdateStatus(ctx context.Context, service *corev1.Service, opts metav1.UpdateOptions) (*corev1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(servicesResource, c.Cluster, "status", c.Namespace, service), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}

func (c *servicesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(servicesResource, c.Cluster, c.Namespace, name, opts), &corev1.Service{})
	return err
}

func (c *servicesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(servicesResource, c.Cluster, c.Namespace, name), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}

// List takes label and field selectors, and returns the list of Services that match those selectors.
func (c *servicesClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(servicesResource, servicesKind, c.Cluster, c.Namespace, opts), &corev1.ServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ServiceList{ListMeta: obj.(*corev1.ServiceList).ListMeta}
	for _, item := range obj.(*corev1.ServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *servicesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(servicesResource, c.Cluster, c.Namespace, opts))
}

func (c *servicesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(servicesResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}

func (c *servicesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Service, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(servicesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}

func (c *servicesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Service, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(servicesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.Service{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Service), err
}