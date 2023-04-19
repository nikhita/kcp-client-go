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

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsextensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	extensionsv1beta1client "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var podSecurityPoliciesResource = schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "podsecuritypolicies"}
var podSecurityPoliciesKind = schema.GroupVersionKind{Group: "extensions", Version: "v1beta1", Kind: "PodSecurityPolicy"}

type podSecurityPoliciesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *podSecurityPoliciesClusterClient) Cluster(clusterPath logicalcluster.Path) extensionsv1beta1client.PodSecurityPolicyInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podSecurityPoliciesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PodSecurityPolicies that match those selectors across all clusters.
func (c *podSecurityPoliciesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.PodSecurityPolicyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(podSecurityPoliciesResource, podSecurityPoliciesKind, logicalcluster.Wildcard, opts), &extensionsv1beta1.PodSecurityPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensionsv1beta1.PodSecurityPolicyList{ListMeta: obj.(*extensionsv1beta1.PodSecurityPolicyList).ListMeta}
	for _, item := range obj.(*extensionsv1beta1.PodSecurityPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PodSecurityPolicies across all clusters.
func (c *podSecurityPoliciesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(podSecurityPoliciesResource, logicalcluster.Wildcard, opts))
}

type podSecurityPoliciesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *podSecurityPoliciesClient) Create(ctx context.Context, podSecurityPolicy *extensionsv1beta1.PodSecurityPolicy, opts metav1.CreateOptions) (*extensionsv1beta1.PodSecurityPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(podSecurityPoliciesResource, c.ClusterPath, podSecurityPolicy), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}

func (c *podSecurityPoliciesClient) Update(ctx context.Context, podSecurityPolicy *extensionsv1beta1.PodSecurityPolicy, opts metav1.UpdateOptions) (*extensionsv1beta1.PodSecurityPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(podSecurityPoliciesResource, c.ClusterPath, podSecurityPolicy), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}

func (c *podSecurityPoliciesClient) UpdateStatus(ctx context.Context, podSecurityPolicy *extensionsv1beta1.PodSecurityPolicy, opts metav1.UpdateOptions) (*extensionsv1beta1.PodSecurityPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(podSecurityPoliciesResource, c.ClusterPath, "status", podSecurityPolicy), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}

func (c *podSecurityPoliciesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(podSecurityPoliciesResource, c.ClusterPath, name, opts), &extensionsv1beta1.PodSecurityPolicy{})
	return err
}

func (c *podSecurityPoliciesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(podSecurityPoliciesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &extensionsv1beta1.PodSecurityPolicyList{})
	return err
}

func (c *podSecurityPoliciesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*extensionsv1beta1.PodSecurityPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(podSecurityPoliciesResource, c.ClusterPath, name), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}

// List takes label and field selectors, and returns the list of PodSecurityPolicies that match those selectors.
func (c *podSecurityPoliciesClient) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.PodSecurityPolicyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(podSecurityPoliciesResource, podSecurityPoliciesKind, c.ClusterPath, opts), &extensionsv1beta1.PodSecurityPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensionsv1beta1.PodSecurityPolicyList{ListMeta: obj.(*extensionsv1beta1.PodSecurityPolicyList).ListMeta}
	for _, item := range obj.(*extensionsv1beta1.PodSecurityPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *podSecurityPoliciesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(podSecurityPoliciesResource, c.ClusterPath, opts))
}

func (c *podSecurityPoliciesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*extensionsv1beta1.PodSecurityPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(podSecurityPoliciesResource, c.ClusterPath, name, pt, data, subresources...), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}

func (c *podSecurityPoliciesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsextensionsv1beta1.PodSecurityPolicyApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.PodSecurityPolicy, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(podSecurityPoliciesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}

func (c *podSecurityPoliciesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsextensionsv1beta1.PodSecurityPolicyApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.PodSecurityPolicy, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(podSecurityPoliciesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &extensionsv1beta1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.PodSecurityPolicy), err
}
