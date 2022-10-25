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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	appsv1listers "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
)

// ReplicaSetClusterLister can list ReplicaSets across all workspaces, or scope down to a ReplicaSetLister for one workspace.
type ReplicaSetClusterLister interface {
	List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error)
	Cluster(cluster logicalcluster.Name) appsv1listers.ReplicaSetLister
}

type replicaSetClusterLister struct {
	indexer cache.Indexer
}

// NewReplicaSetClusterLister returns a new ReplicaSetClusterLister.
func NewReplicaSetClusterLister(indexer cache.Indexer) *replicaSetClusterLister {
	return &replicaSetClusterLister{indexer: indexer}
}

// List lists all ReplicaSets in the indexer across all workspaces.
func (s *replicaSetClusterLister) List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*appsv1.ReplicaSet))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ReplicaSets.
func (s *replicaSetClusterLister) Cluster(cluster logicalcluster.Name) appsv1listers.ReplicaSetLister {
	return &replicaSetLister{indexer: s.indexer, cluster: cluster}
}

// replicaSetLister implements the appsv1listers.ReplicaSetLister interface.
type replicaSetLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all ReplicaSets in the indexer for a workspace.
func (s *replicaSetLister) List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*appsv1.ReplicaSet)
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

// ReplicaSets returns an object that can list and get ReplicaSets in one namespace.
func (s *replicaSetLister) ReplicaSets(namespace string) appsv1listers.ReplicaSetNamespaceLister {
	return &replicaSetNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// replicaSetNamespaceLister implements the appsv1listers.ReplicaSetNamespaceLister interface.
type replicaSetNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all ReplicaSets in the indexer for a given workspace and namespace.
func (s *replicaSetNamespaceLister) List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error) {
	selectAll := selector == nil || selector.Empty()

	var list []interface{}
	if s.namespace == metav1.NamespaceAll {
		list, err = s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	} else {
		list, err = s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	}
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*appsv1.ReplicaSet)
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

// Get retrieves the ReplicaSet from the indexer for a given workspace, namespace and name.
func (s *replicaSetNamespaceLister) Get(name string) (*appsv1.ReplicaSet, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1.Resource("ReplicaSet"), name)
	}
	return obj.(*appsv1.ReplicaSet), nil
}