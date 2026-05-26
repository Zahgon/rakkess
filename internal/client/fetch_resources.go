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
	"github.com/corneliusweig/rakkess/internal/options"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
)

var (
	// for testing
	getDiscoveryClient = getDiscoveryClientImpl
)

// GroupResource contains the APIGroup and APIResource
type GroupResource struct {
	APIGroup    string
	APIResource metav1.APIResource
}

// Extracts the full name including APIGroup, e.g. 'deployment.apps'
func (g GroupResource) fullName() string { _ = "STUB: not implemented"; return "" }

// FetchAvailableGroupResources fetches a list of known APIResources on the server.
func FetchAvailableGroupResources(opts *options.RakkessOptions) ([]GroupResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDiscoveryClientImpl(opts *options.RakkessOptions) (discovery.CachedDiscoveryInterface, error) {
	_ = "STUB: not implemented"
	return *new(discovery.CachedDiscoveryInterface), nil
}
