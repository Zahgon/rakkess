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
	"github.com/corneliusweig/rakkess/internal/options"
	clientv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

var (
	// for testing
	getRbacClient = getRbacClientImpl
)

const (
	clusterRoleName = "ClusterRole"
	roleName        = "Role"
)

// GetSubjectAccess determines subjects with access to the given resource.
func GetSubjectAccess(ctx context.Context, opts *options.RakkessOptions, resource, resourceName string) (*result.SubjectAccess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveRoleBindings(ctx context.Context, cli clientv1.RoleBindingsGetter, sa *result.SubjectAccess, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func resolveClusterRoleBindings(ctx context.Context, cli clientv1.ClusterRoleBindingsGetter, sa *result.SubjectAccess) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchMatchingClusterRoles(ctx context.Context, rbacClient clientv1.ClusterRolesGetter, sa *result.SubjectAccess) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchMatchingRoles(ctx context.Context, rbacClient clientv1.RolesGetter, sa *result.SubjectAccess, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func getRbacClientImpl(o *options.RakkessOptions) (clientv1.RbacV1Interface, error) {
	_ = "STUB: not implemented"
	return *new(clientv1.RbacV1Interface), nil
}
