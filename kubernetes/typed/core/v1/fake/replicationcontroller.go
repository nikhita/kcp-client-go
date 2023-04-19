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

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"

	kcpcorev1 "github.com/kcp-dev/client-go/kubernetes/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var replicationControllersResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "replicationcontrollers"}
var replicationControllersKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ReplicationController"}

type replicationControllersClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *replicationControllersClusterClient) Cluster(clusterPath logicalcluster.Path) kcpcorev1.ReplicationControllersNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &replicationControllersNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ReplicationControllers that match those selectors across all clusters.
func (c *replicationControllersClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ReplicationControllerList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(replicationControllersResource, replicationControllersKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.ReplicationControllerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ReplicationControllerList{ListMeta: obj.(*corev1.ReplicationControllerList).ListMeta}
	for _, item := range obj.(*corev1.ReplicationControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ReplicationControllers across all clusters.
func (c *replicationControllersClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(replicationControllersResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type replicationControllersNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *replicationControllersNamespacer) Namespace(namespace string) corev1client.ReplicationControllerInterface {
	return &replicationControllersClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type replicationControllersClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *replicationControllersClient) Create(ctx context.Context, replicationController *corev1.ReplicationController, opts metav1.CreateOptions) (*corev1.ReplicationController, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(replicationControllersResource, c.ClusterPath, c.Namespace, replicationController), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

func (c *replicationControllersClient) Update(ctx context.Context, replicationController *corev1.ReplicationController, opts metav1.UpdateOptions) (*corev1.ReplicationController, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(replicationControllersResource, c.ClusterPath, c.Namespace, replicationController), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

func (c *replicationControllersClient) UpdateStatus(ctx context.Context, replicationController *corev1.ReplicationController, opts metav1.UpdateOptions) (*corev1.ReplicationController, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(replicationControllersResource, c.ClusterPath, "status", c.Namespace, replicationController), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

func (c *replicationControllersClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(replicationControllersResource, c.ClusterPath, c.Namespace, name, opts), &corev1.ReplicationController{})
	return err
}

func (c *replicationControllersClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(replicationControllersResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ReplicationControllerList{})
	return err
}

func (c *replicationControllersClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.ReplicationController, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(replicationControllersResource, c.ClusterPath, c.Namespace, name), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

// List takes label and field selectors, and returns the list of ReplicationControllers that match those selectors.
func (c *replicationControllersClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ReplicationControllerList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(replicationControllersResource, replicationControllersKind, c.ClusterPath, c.Namespace, opts), &corev1.ReplicationControllerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ReplicationControllerList{ListMeta: obj.(*corev1.ReplicationControllerList).ListMeta}
	for _, item := range obj.(*corev1.ReplicationControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *replicationControllersClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(replicationControllersResource, c.ClusterPath, c.Namespace, opts))
}

func (c *replicationControllersClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.ReplicationController, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicationControllersResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

func (c *replicationControllersClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ReplicationControllerApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ReplicationController, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicationControllersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

func (c *replicationControllersClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ReplicationControllerApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ReplicationController, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicationControllersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.ReplicationController{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ReplicationController), err
}

func (c *replicationControllersClient) GetScale(ctx context.Context, replicationControllerName string, options metav1.GetOptions) (*autoscalingv1.Scale, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetSubresourceAction(replicationControllersResource, c.ClusterPath, "scale", c.Namespace, replicationControllerName), &autoscalingv1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}

func (c *replicationControllersClient) UpdateScale(ctx context.Context, replicationControllerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(replicationControllersResource, c.ClusterPath, "scale", c.Namespace, scale), &autoscalingv1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
