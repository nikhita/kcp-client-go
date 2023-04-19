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

	"github.com/kcp-dev/logicalcluster/v3"

	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	authorizationv1client "k8s.io/client-go/kubernetes/typed/authorization/v1"

	kcpauthorizationv1 "github.com/kcp-dev/client-go/kubernetes/typed/authorization/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var localSubjectAccessReviewsResource = schema.GroupVersionResource{Group: "authorization.k8s.io", Version: "v1", Resource: "localsubjectaccessreviews"}
var localSubjectAccessReviewsKind = schema.GroupVersionKind{Group: "authorization.k8s.io", Version: "v1", Kind: "LocalSubjectAccessReview"}

type localSubjectAccessReviewsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *localSubjectAccessReviewsClusterClient) Cluster(clusterPath logicalcluster.Path) kcpauthorizationv1.LocalSubjectAccessReviewsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &localSubjectAccessReviewsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

type localSubjectAccessReviewsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *localSubjectAccessReviewsNamespacer) Namespace(namespace string) authorizationv1client.LocalSubjectAccessReviewInterface {
	return &localSubjectAccessReviewsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type localSubjectAccessReviewsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *localSubjectAccessReviewsClient) Create(ctx context.Context, localSubjectAccessReview *authorizationv1.LocalSubjectAccessReview, opts metav1.CreateOptions) (*authorizationv1.LocalSubjectAccessReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(localSubjectAccessReviewsResource, c.ClusterPath, c.Namespace, localSubjectAccessReview), &authorizationv1.LocalSubjectAccessReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorizationv1.LocalSubjectAccessReview), err
}
