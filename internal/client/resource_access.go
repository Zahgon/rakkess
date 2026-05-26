/*
Copyright 2020 Cornelius Weig

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

package client

import (
	"context"

	"github.com/corneliusweig/rakkess/internal/client/result"
	authv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

// CheckResourceAccess determines the access rights for the given GroupResources and verbs.
// Since it needs to do a lot of requests, the SelfSubjectAccessReviewInterface needs to
// be configured for high queries per second.
func CheckResourceAccess(ctx context.Context, sar authv1.SelfSubjectAccessReviewInterface, grs []GroupResource, verbs []string, namespace *string) result.ResourceAccess {
	_ = "STUB: not implemented"
	// guards res
	return *new(result.ResourceAccess)
}

// copy captured variables

// This seems to be a bug in kubernetes. If namespace is set for non-namespaced
// resources, the access is reported as "allowed", but in fact it is forbidden.
