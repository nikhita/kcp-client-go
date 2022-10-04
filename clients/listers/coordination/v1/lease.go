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

	coordinationv1 "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	coordinationv1listers "k8s.io/client-go/listers/coordination/v1"
	"k8s.io/client-go/tools/cache"
)

// LeaseClusterLister can list Leases across all workspaces, or scope down to a LeaseLister for one workspace.
type LeaseClusterLister interface {
	List(selector labels.Selector) (ret []*coordinationv1.Lease, err error)
	Cluster(cluster logicalcluster.Name) coordinationv1listers.LeaseLister
}

type leaseClusterLister struct {
	indexer cache.Indexer
}

// NewLeaseClusterLister returns a new LeaseClusterLister.
func NewLeaseClusterLister(indexer cache.Indexer) *leaseClusterLister {
	return &leaseClusterLister{indexer: indexer}
}

// List lists all Leases in the indexer across all workspaces.
func (s *leaseClusterLister) List(selector labels.Selector) (ret []*coordinationv1.Lease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*coordinationv1.Lease))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Leases.
func (s *leaseClusterLister) Cluster(cluster logicalcluster.Name) coordinationv1listers.LeaseLister {
	return &leaseLister{indexer: s.indexer, cluster: cluster}
}

// leaseLister implements the coordinationv1listers.LeaseLister interface.
type leaseLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all Leases in the indexer for a workspace.
func (s *leaseLister) List(selector labels.Selector) (ret []*coordinationv1.Lease, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*coordinationv1.Lease)
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

// Leases returns an object that can list and get Leases in one namespace.
func (s *leaseLister) Leases(namespace string) coordinationv1listers.LeaseNamespaceLister {
	return &leaseNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// leaseNamespaceLister implements the coordinationv1listers.LeaseNamespaceLister interface.
type leaseNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all Leases in the indexer for a given workspace and namespace.
func (s *leaseNamespaceLister) List(selector labels.Selector) (ret []*coordinationv1.Lease, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*coordinationv1.Lease)
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

// Get retrieves the Lease from the indexer for a given workspace, namespace and name.
func (s *leaseNamespaceLister) Get(name string) (*coordinationv1.Lease, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(coordinationv1.Resource("Lease"), name)
	}
	return obj.(*coordinationv1.Lease), nil
}
