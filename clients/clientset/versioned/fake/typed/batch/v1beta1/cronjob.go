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
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsbatchv1beta1 "k8s.io/client-go/applyconfigurations/batch/v1beta1"
	batchv1beta1client "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	"k8s.io/client-go/testing"

	kcpbatchv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/batch/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var cronJobsResource = schema.GroupVersionResource{Group: "batch", Version: "V1beta1", Resource: "cronjobs"}
var cronJobsKind = schema.GroupVersionKind{Group: "batch", Version: "V1beta1", Kind: "CronJob"}

type cronJobsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cronJobsClusterClient) Cluster(cluster logicalcluster.Name) kcpbatchv1beta1.CronJobsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cronJobsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of CronJobs that match those selectors across all clusters.
func (c *cronJobsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*batchv1beta1.CronJobList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cronJobsResource, cronJobsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &batchv1beta1.CronJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batchv1beta1.CronJobList{ListMeta: obj.(*batchv1beta1.CronJobList).ListMeta}
	for _, item := range obj.(*batchv1beta1.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested CronJobs across all clusters.
func (c *cronJobsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cronJobsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type cronJobsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *cronJobsNamespacer) Namespace(namespace string) batchv1beta1client.CronJobInterface {
	return &cronJobsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type cronJobsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *cronJobsClient) Create(ctx context.Context, cronJob *batchv1beta1.CronJob, opts metav1.CreateOptions) (*batchv1beta1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(cronJobsResource, c.Cluster, c.Namespace, cronJob), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}

func (c *cronJobsClient) Update(ctx context.Context, cronJob *batchv1beta1.CronJob, opts metav1.UpdateOptions) (*batchv1beta1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(cronJobsResource, c.Cluster, c.Namespace, cronJob), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}

func (c *cronJobsClient) UpdateStatus(ctx context.Context, cronJob *batchv1beta1.CronJob, opts metav1.UpdateOptions) (*batchv1beta1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(cronJobsResource, c.Cluster, "status", c.Namespace, cronJob), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}

func (c *cronJobsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(cronJobsResource, c.Cluster, c.Namespace, name, opts), &batchv1beta1.CronJob{})
	return err
}

func (c *cronJobsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(cronJobsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &batchv1beta1.CronJobList{})
	return err
}

func (c *cronJobsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*batchv1beta1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(cronJobsResource, c.Cluster, c.Namespace, name), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}

// List takes label and field selectors, and returns the list of CronJobs that match those selectors.
func (c *cronJobsClient) List(ctx context.Context, opts metav1.ListOptions) (*batchv1beta1.CronJobList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cronJobsResource, cronJobsKind, c.Cluster, c.Namespace, opts), &batchv1beta1.CronJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batchv1beta1.CronJobList{ListMeta: obj.(*batchv1beta1.CronJobList).ListMeta}
	for _, item := range obj.(*batchv1beta1.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cronJobsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cronJobsResource, c.Cluster, c.Namespace, opts))
}

func (c *cronJobsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*batchv1beta1.CronJob, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cronJobsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}

func (c *cronJobsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsbatchv1beta1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1beta1.CronJob, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cronJobsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}

func (c *cronJobsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsbatchv1beta1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1beta1.CronJob, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cronJobsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &batchv1beta1.CronJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*batchv1beta1.CronJob), err
}
