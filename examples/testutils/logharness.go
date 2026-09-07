package testutils

import (
	"bufio"
	"bytes"
	"testing"
)

type LogHarness struct {
	buf       bytes.Buffer
	sequences []*Sequence
}

type Expectation interface {
	IsMatch(line string) bool
	String() string
}

func (h *LogHarness) Run(t *testing.T, f func()) { _ = "STUB: not implemented"; return }

func (h *LogHarness) Expect(s string) { _ = "STUB: not implemented"; return }

func (h *LogHarness) ExpectPrefix(s string) { _ = "STUB: not implemented"; return }

func (h *LogHarness) NewSequence(name string) *Sequence { _ = "STUB: not implemented"; return nil }

type prefix string

func (p prefix) IsMatch(line string) bool { _ = "STUB: not implemented"; return false }

func (p prefix) String() string { _ = "STUB: not implemented"; return "" }

type text string

func (t text) IsMatch(line string) bool { _ = "STUB: not implemented"; return false }

func (t text) String() string { _ = "STUB: not implemented"; return "" }

type Sequence struct {
	name string
	exp  []Expectation
}

func (seq *Sequence) Assert(t *testing.T, s *bufio.Scanner) { _ = "STUB: not implemented"; return }

func (seq *Sequence) Expect(s string) { _ = "STUB: not implemented"; return }

func (seq *Sequence) ExpectPrefix(s string) { _ = "STUB: not implemented"; return }
