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

package result

import (
	"github.com/corneliusweig/rakkess/internal/printer"
)

// ResourceAccess holds the access result for all resources.
type ResourceAccess map[string]map[string]Access

// Print implements MatrixPrinter.Print. It prints a tab-separated table with a header.
func (ra ResourceAccess) Table(verbs []string) *printer.Table {
	_ = "STUB: not implemented"
	return nil
}

// table header

// table body
