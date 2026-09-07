package rcmgr

import (
	"context"
	"encoding/json"
	"io"
	"sync"

	"github.com/libp2p/go-libp2p/core/network"
)

type trace struct {
	path string

	ctx    context.Context
	cancel func()
	wg     sync.WaitGroup

	mx            sync.Mutex
	done          bool
	pendingWrites []any
	reporters     []TraceReporter
}

type TraceReporter interface {
	ConsumeEvent(TraceEvt)
}

func WithTrace(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTraceReporter(reporter TraceReporter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type TraceEvtTyp string

const (
	TraceStartEvt              TraceEvtTyp = "start"
	TraceCreateScopeEvt        TraceEvtTyp = "create_scope"
	TraceDestroyScopeEvt       TraceEvtTyp = "destroy_scope"
	TraceReserveMemoryEvt      TraceEvtTyp = "reserve_memory"
	TraceBlockReserveMemoryEvt TraceEvtTyp = "block_reserve_memory"
	TraceReleaseMemoryEvt      TraceEvtTyp = "release_memory"
	TraceAddStreamEvt          TraceEvtTyp = "add_stream"
	TraceBlockAddStreamEvt     TraceEvtTyp = "block_add_stream"
	TraceRemoveStreamEvt       TraceEvtTyp = "remove_stream"
	TraceAddConnEvt            TraceEvtTyp = "add_conn"
	TraceBlockAddConnEvt       TraceEvtTyp = "block_add_conn"
	TraceRemoveConnEvt         TraceEvtTyp = "remove_conn"
)

type scopeClass struct {
	name string
}

func (s scopeClass) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type TraceEvt struct {
	Time string
	Type TraceEvtTyp

	Scope *scopeClass `json:",omitempty"`
	Name  string      `json:",omitempty"`

	Limit any `json:",omitempty"`

	Priority uint8 `json:",omitempty"`

	Delta    int64 `json:",omitempty"`
	DeltaIn  int   `json:",omitempty"`
	DeltaOut int   `json:",omitempty"`

	Memory int64 `json:",omitempty"`

	StreamsIn  int `json:",omitempty"`
	StreamsOut int `json:",omitempty"`

	ConnsIn  int `json:",omitempty"`
	ConnsOut int `json:",omitempty"`

	FD int `json:",omitempty"`
}

func (t *trace) push(evt TraceEvt) { _ = "STUB: not implemented"; return }

func (t *trace) backgroundWriter(out io.WriteCloser) { _ = "STUB: not implemented"; return }

func (t *trace) writeEvents(pend []any, jout *json.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *trace) Start(limits Limiter) error { _ = "STUB: not implemented"; return nil }

func (t *trace) Close() error { _ = "STUB: not implemented"; return nil }

func (t *trace) CreateScope(scope string, limit Limit) { _ = "STUB: not implemented"; return }

func (t *trace) DestroyScope(scope string) { _ = "STUB: not implemented"; return }

func (t *trace) ReserveMemory(scope string, prio uint8, size, mem int64) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) BlockReserveMemory(scope string, prio uint8, size, mem int64) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) ReleaseMemory(scope string, size, mem int64) { _ = "STUB: not implemented"; return }

func (t *trace) AddStream(scope string, dir network.Direction, nstreamsIn, nstreamsOut int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) BlockAddStream(scope string, dir network.Direction, nstreamsIn, nstreamsOut int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) RemoveStream(scope string, dir network.Direction, nstreamsIn, nstreamsOut int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) AddStreams(scope string, deltaIn, deltaOut, nstreamsIn, nstreamsOut int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) BlockAddStreams(scope string, deltaIn, deltaOut, nstreamsIn, nstreamsOut int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) RemoveStreams(scope string, deltaIn, deltaOut, nstreamsIn, nstreamsOut int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) AddConn(scope string, dir network.Direction, usefd bool, nconnsIn, nconnsOut, nfd int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) BlockAddConn(scope string, dir network.Direction, usefd bool, nconnsIn, nconnsOut, nfd int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) RemoveConn(scope string, dir network.Direction, usefd bool, nconnsIn, nconnsOut, nfd int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) AddConns(scope string, deltaIn, deltaOut, deltafd, nconnsIn, nconnsOut, nfd int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) BlockAddConns(scope string, deltaIn, deltaOut, deltafd, nconnsIn, nconnsOut, nfd int) {
	_ = "STUB: not implemented"
	return
}

func (t *trace) RemoveConns(scope string, deltaIn, deltaOut, deltafd, nconnsIn, nconnsOut, nfd int) {
	_ = "STUB: not implemented"
	return
}
