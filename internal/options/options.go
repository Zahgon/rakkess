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

package options

import (
	"bytes"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/discovery"
	v1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

// RakkessOptions holds all user configuration options.
type RakkessOptions struct {
	ConfigFlags      *genericclioptions.ConfigFlags
	Verbs            []string
	AsServiceAccount string
	OutputFormat     string
	Streams          *genericclioptions.IOStreams
}

// NewRakkessOptions creates RakkessOptions with defaults.
func NewRakkessOptions() *RakkessOptions { _ = "STUB: not implemented"; return nil }

// Sets up options with in-memory buffers as in- and output-streams
func NewTestRakkessOptions() (*RakkessOptions, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// GetAuthClient creates a client for SelfSubjectAccessReviews with high queries per second.
func (o *RakkessOptions) GetAuthClient() (v1.SelfSubjectAccessReviewInterface, error) {
	_ = "STUB: not implemented"
	return *new(v1.SelfSubjectAccessReviewInterface), nil
}

// DiscoveryClient creates a kubernetes discovery client.
func (o *RakkessOptions) DiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	_ = "STUB: not implemented"
	return *new(discovery.CachedDiscoveryInterface), nil
}

func (o *RakkessOptions) ExpandServiceAccount() error { _ = "STUB: not implemented"; return nil }

func (o *RakkessOptions) namespacedServiceAccount() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExpandVerbs expands wildcard verbs `*` and `all`.
func (o *RakkessOptions) ExpandVerbs() { _ = "STUB: not implemented"; return }
