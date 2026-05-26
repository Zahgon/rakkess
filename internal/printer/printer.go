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

package printer

import (
	"io"
	"sync"
)

type color int

const (
	red    = color(31)
	green  = color(32)
	purple = color(35)
	none   = color(0)
)

var (
	isTerminal = isTerminalImpl
	once       sync.Once
)

type Outcome uint8

const (
	None Outcome = iota
	Up
	Down
	Err
)

type Row struct {
	Intro   []string
	Entries []Outcome
}
type Table struct {
	Headers []string
	Rows    []Row
}

func TableWithHeaders(headers []string) *Table { _ = "STUB: not implemented"; return nil }

func (p *Table) AddRow(intro []string, outcomes ...Outcome) { _ = "STUB: not implemented"; return }

func (p *Table) Render(out io.Writer, outputFormat string) { _ = "STUB: not implemented"; return }

// table header

// table body

// FIXME

func humanreadableAccessCode(o Outcome) string { _ = "STUB: not implemented"; return "" }

// ✓

// ✕

func colored(wrap func(Outcome) string) func(Outcome) string { _ = "STUB: not implemented"; return nil }

func asciiAccessCode(o Outcome) string { _ = "STUB: not implemented"; return "" }
