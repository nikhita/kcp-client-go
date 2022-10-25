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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	storagev1beta1 "k8s.io/api/storage/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	storagev1beta1listers "k8s.io/client-go/listers/storage/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// StorageClassClusterLister can list StorageClasses across all workspaces, or scope down to a StorageClassLister for one workspace.
type StorageClassClusterLister interface {
	List(selector labels.Selector) (ret []*storagev1beta1.StorageClass, err error)
	Cluster(cluster logicalcluster.Name) storagev1beta1listers.StorageClassLister
}

type storageClassClusterLister struct {
	indexer cache.Indexer
}

// NewStorageClassClusterLister returns a new StorageClassClusterLister.
func NewStorageClassClusterLister(indexer cache.Indexer) *storageClassClusterLister {
	return &storageClassClusterLister{indexer: indexer}
}

// List lists all StorageClasses in the indexer across all workspaces.
func (s *storageClassClusterLister) List(selector labels.Selector) (ret []*storagev1beta1.StorageClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storagev1beta1.StorageClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get StorageClasses.
func (s *storageClassClusterLister) Cluster(cluster logicalcluster.Name) storagev1beta1listers.StorageClassLister {
	return &storageClassLister{indexer: s.indexer, cluster: cluster}
}

// storageClassLister implements the storagev1beta1listers.StorageClassLister interface.
type storageClassLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all StorageClasses in the indexer for a workspace.
func (s *storageClassLister) List(selector labels.Selector) (ret []*storagev1beta1.StorageClass, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*storagev1beta1.StorageClass)
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

// Get retrieves the StorageClass from the indexer for a given workspace and name.
func (s *storageClassLister) Get(name string) (*storagev1beta1.StorageClass, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storagev1beta1.Resource("StorageClass"), name)
	}
	return obj.(*storagev1beta1.StorageClass), nil
}