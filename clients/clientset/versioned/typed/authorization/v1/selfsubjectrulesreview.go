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
	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	authorizationv1client "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

// SelfSubjectRulesReviewsClusterGetter has a method to return a SelfSubjectRulesReviewsClusterInterface.
// A group's cluster client should implement this interface.
type SelfSubjectRulesReviewsClusterGetter interface {
	SelfSubjectRulesReviews() SelfSubjectRulesReviewsClusterInterface
}

// SelfSubjectRulesReviewsClusterInterface can scope down to one cluster and return a authorizationv1client.SelfSubjectRulesReviewInterface.
type SelfSubjectRulesReviewsClusterInterface interface {
	Cluster(logicalcluster.Name) authorizationv1client.SelfSubjectRulesReviewInterface
}

type selfSubjectRulesReviewsClusterInterface struct {
	clientCache kcpclient.Cache[*authorizationv1client.AuthorizationV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *selfSubjectRulesReviewsClusterInterface) Cluster(name logicalcluster.Name) authorizationv1client.SelfSubjectRulesReviewInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).SelfSubjectRulesReviews()
}
