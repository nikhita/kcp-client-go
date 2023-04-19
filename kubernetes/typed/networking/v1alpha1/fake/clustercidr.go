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
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnetworkingv1alpha1 "k8s.io/client-go/applyconfigurations/networking/v1alpha1"
	networkingv1alpha1client "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var clusterCIDRsResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1alpha1", Resource: "clustercidrs"}
var clusterCIDRsKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1alpha1", Kind: "ClusterCIDR"}

type clusterCIDRsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterCIDRsClusterClient) Cluster(clusterPath logicalcluster.Path) networkingv1alpha1client.ClusterCIDRInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterCIDRsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ClusterCIDRs that match those selectors across all clusters.
func (c *clusterCIDRsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1alpha1.ClusterCIDRList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterCIDRsResource, clusterCIDRsKind, logicalcluster.Wildcard, opts), &networkingv1alpha1.ClusterCIDRList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1alpha1.ClusterCIDRList{ListMeta: obj.(*networkingv1alpha1.ClusterCIDRList).ListMeta}
	for _, item := range obj.(*networkingv1alpha1.ClusterCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterCIDRs across all clusters.
func (c *clusterCIDRsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterCIDRsResource, logicalcluster.Wildcard, opts))
}

type clusterCIDRsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *clusterCIDRsClient) Create(ctx context.Context, clusterCIDR *networkingv1alpha1.ClusterCIDR, opts metav1.CreateOptions) (*networkingv1alpha1.ClusterCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterCIDRsResource, c.ClusterPath, clusterCIDR), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}

func (c *clusterCIDRsClient) Update(ctx context.Context, clusterCIDR *networkingv1alpha1.ClusterCIDR, opts metav1.UpdateOptions) (*networkingv1alpha1.ClusterCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterCIDRsResource, c.ClusterPath, clusterCIDR), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}

func (c *clusterCIDRsClient) UpdateStatus(ctx context.Context, clusterCIDR *networkingv1alpha1.ClusterCIDR, opts metav1.UpdateOptions) (*networkingv1alpha1.ClusterCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterCIDRsResource, c.ClusterPath, "status", clusterCIDR), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}

func (c *clusterCIDRsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterCIDRsResource, c.ClusterPath, name, opts), &networkingv1alpha1.ClusterCIDR{})
	return err
}

func (c *clusterCIDRsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterCIDRsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1alpha1.ClusterCIDRList{})
	return err
}

func (c *clusterCIDRsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*networkingv1alpha1.ClusterCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterCIDRsResource, c.ClusterPath, name), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}

// List takes label and field selectors, and returns the list of ClusterCIDRs that match those selectors.
func (c *clusterCIDRsClient) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1alpha1.ClusterCIDRList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterCIDRsResource, clusterCIDRsKind, c.ClusterPath, opts), &networkingv1alpha1.ClusterCIDRList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1alpha1.ClusterCIDRList{ListMeta: obj.(*networkingv1alpha1.ClusterCIDRList).ListMeta}
	for _, item := range obj.(*networkingv1alpha1.ClusterCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterCIDRsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterCIDRsResource, c.ClusterPath, opts))
}

func (c *clusterCIDRsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*networkingv1alpha1.ClusterCIDR, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterCIDRsResource, c.ClusterPath, name, pt, data, subresources...), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}

func (c *clusterCIDRsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1alpha1.ClusterCIDRApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1alpha1.ClusterCIDR, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterCIDRsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}

func (c *clusterCIDRsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsnetworkingv1alpha1.ClusterCIDRApplyConfiguration, opts metav1.ApplyOptions) (*networkingv1alpha1.ClusterCIDR, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterCIDRsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &networkingv1alpha1.ClusterCIDR{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1alpha1.ClusterCIDR), err
}
