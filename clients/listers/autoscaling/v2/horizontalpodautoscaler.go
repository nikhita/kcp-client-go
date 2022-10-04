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

package v2

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	autoscalingv2 "k8s.io/api/autoscaling/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	autoscalingv2listers "k8s.io/client-go/listers/autoscaling/v2"
	"k8s.io/client-go/tools/cache"
)

// HorizontalPodAutoscalerClusterLister can list HorizontalPodAutoscalers across all workspaces, or scope down to a HorizontalPodAutoscalerLister for one workspace.
type HorizontalPodAutoscalerClusterLister interface {
	List(selector labels.Selector) (ret []*autoscalingv2.HorizontalPodAutoscaler, err error)
	Cluster(cluster logicalcluster.Name) autoscalingv2listers.HorizontalPodAutoscalerLister
}

type horizontalPodAutoscalerClusterLister struct {
	indexer cache.Indexer
}

// NewHorizontalPodAutoscalerClusterLister returns a new HorizontalPodAutoscalerClusterLister.
func NewHorizontalPodAutoscalerClusterLister(indexer cache.Indexer) *horizontalPodAutoscalerClusterLister {
	return &horizontalPodAutoscalerClusterLister{indexer: indexer}
}

// List lists all HorizontalPodAutoscalers in the indexer across all workspaces.
func (s *horizontalPodAutoscalerClusterLister) List(selector labels.Selector) (ret []*autoscalingv2.HorizontalPodAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*autoscalingv2.HorizontalPodAutoscaler))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get HorizontalPodAutoscalers.
func (s *horizontalPodAutoscalerClusterLister) Cluster(cluster logicalcluster.Name) autoscalingv2listers.HorizontalPodAutoscalerLister {
	return &horizontalPodAutoscalerLister{indexer: s.indexer, cluster: cluster}
}

// horizontalPodAutoscalerLister implements the autoscalingv2listers.HorizontalPodAutoscalerLister interface.
type horizontalPodAutoscalerLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all HorizontalPodAutoscalers in the indexer for a workspace.
func (s *horizontalPodAutoscalerLister) List(selector labels.Selector) (ret []*autoscalingv2.HorizontalPodAutoscaler, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*autoscalingv2.HorizontalPodAutoscaler)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// HorizontalPodAutoscalers returns an object that can list and get HorizontalPodAutoscalers in one namespace.
func (s *horizontalPodAutoscalerLister) HorizontalPodAutoscalers(namespace string) autoscalingv2listers.HorizontalPodAutoscalerNamespaceLister {
	return &horizontalPodAutoscalerNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// horizontalPodAutoscalerNamespaceLister implements the autoscalingv2listers.HorizontalPodAutoscalerNamespaceLister interface.
type horizontalPodAutoscalerNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all HorizontalPodAutoscalers in the indexer for a given workspace and namespace.
func (s *horizontalPodAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*autoscalingv2.HorizontalPodAutoscaler, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*autoscalingv2.HorizontalPodAutoscaler)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the HorizontalPodAutoscaler from the indexer for a given workspace, namespace and name.
func (s *horizontalPodAutoscalerNamespaceLister) Get(name string) (*autoscalingv2.HorizontalPodAutoscaler, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(autoscalingv2.Resource("HorizontalPodAutoscaler"), name)
	}
	return obj.(*autoscalingv2.HorizontalPodAutoscaler), nil
}
