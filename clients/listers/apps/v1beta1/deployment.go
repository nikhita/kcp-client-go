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

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	appsv1beta1listers "k8s.io/client-go/listers/apps/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// DeploymentClusterLister can list Deployments across all workspaces, or scope down to a DeploymentLister for one workspace.
type DeploymentClusterLister interface {
	List(selector labels.Selector) (ret []*appsv1beta1.Deployment, err error)
	Cluster(cluster logicalcluster.Name) appsv1beta1listers.DeploymentLister
}

type deploymentClusterLister struct {
	indexer cache.Indexer
}

// NewDeploymentClusterLister returns a new DeploymentClusterLister.
func NewDeploymentClusterLister(indexer cache.Indexer) *deploymentClusterLister {
	return &deploymentClusterLister{indexer: indexer}
}

// List lists all Deployments in the indexer across all workspaces.
func (s *deploymentClusterLister) List(selector labels.Selector) (ret []*appsv1beta1.Deployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*appsv1beta1.Deployment))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Deployments.
func (s *deploymentClusterLister) Cluster(cluster logicalcluster.Name) appsv1beta1listers.DeploymentLister {
	return &deploymentLister{indexer: s.indexer, cluster: cluster}
}

// deploymentLister implements the appsv1beta1listers.DeploymentLister interface.
type deploymentLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all Deployments in the indexer for a workspace.
func (s *deploymentLister) List(selector labels.Selector) (ret []*appsv1beta1.Deployment, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(s.cluster))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*appsv1beta1.Deployment)
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

// Deployments returns an object that can list and get Deployments in one namespace.
func (s *deploymentLister) Deployments(namespace string) appsv1beta1listers.DeploymentNamespaceLister {
	return &deploymentNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// deploymentNamespaceLister implements the appsv1beta1listers.DeploymentNamespaceLister interface.
type deploymentNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all Deployments in the indexer for a given workspace and namespace.
func (s *deploymentNamespaceLister) List(selector labels.Selector) (ret []*appsv1beta1.Deployment, err error) {
	selectAll := selector == nil || selector.Empty()

	list, err := s.indexer.ByIndex(kcpcache.ClusterAndNamespaceIndexName, kcpcache.ClusterAndNamespaceIndexKey(s.cluster, s.namespace))
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*appsv1beta1.Deployment)
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

// Get retrieves the Deployment from the indexer for a given workspace, namespace and name.
func (s *deploymentNamespaceLister) Get(name string) (*appsv1beta1.Deployment, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1beta1.Resource("Deployment"), name)
	}
	return obj.(*appsv1beta1.Deployment), nil
}
