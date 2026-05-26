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

package validation

import (
	"github.com/corneliusweig/rakkess/internal/options"
)

// Options validates RakkessOptions. Fields validated:
// - OutputFormat
// - Verbs
func Options(opts *options.RakkessOptions) error { _ = "STUB: not implemented"; return nil }

func OutputFormat(format string) error { _ = "STUB: not implemented"; return nil }

func verbs(verbs []string) error { _ = "STUB: not implemented"; return nil }
