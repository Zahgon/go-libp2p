package rcmgr

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/network"
)

type resources struct {
	limit Limit

	nconnsIn, nconnsOut     int
	nstreamsIn, nstreamsOut int
	nfd                     int

	memory int64
}

type resourceScope struct {
	sync.Mutex
	done   bool
	refCnt int

	spanID int

	rc    resources
	owner *resourceScope
	edges []*resourceScope

	name    string
	trace   *trace
	metrics *metrics
}

var _ network.ResourceScope = (*resourceScope)(nil)
var _ network.ResourceScopeSpan = (*resourceScope)(nil)

func newResourceScope(limit Limit, edges []*resourceScope, name string, trace *trace, metrics *metrics) *resourceScope {
	_ = "STUB: not implemented"
	return nil
}

func newResourceScopeSpan(owner *resourceScope, id int) *resourceScope {
	_ = "STUB: not implemented"
	return nil
}

func IsSpan(name string) bool { _ = "STUB: not implemented"; return false }

func addInt64WithOverflow(a int64, b int64) (c int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func mulInt64WithOverflow(a, b int64) (c int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (rc *resources) checkMemory(rsvp int64, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *resources) reserveMemory(size int64, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *resources) releaseMemory(size int64) { _ = "STUB: not implemented"; return }

func (rc *resources) addStream(dir network.Direction) error { _ = "STUB: not implemented"; return nil }

func (rc *resources) addStreams(incount, outcount int) error { _ = "STUB: not implemented"; return nil }

func (rc *resources) removeStream(dir network.Direction) { _ = "STUB: not implemented"; return }

func (rc *resources) removeStreams(incount, outcount int) { _ = "STUB: not implemented"; return }

func (rc *resources) addConn(dir network.Direction, usefd bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *resources) addConns(incount, outcount, fdcount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *resources) removeConn(dir network.Direction, usefd bool) {
	_ = "STUB: not implemented"
	return
}

func (rc *resources) removeConns(incount, outcount, fdcount int) { _ = "STUB: not implemented"; return }

func (rc *resources) stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (s *resourceScope) wrapError(err error) error { _ = "STUB: not implemented"; return nil }

func (s *resourceScope) ReserveMemory(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) reserveMemoryForEdges(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) releaseMemoryForEdges(size int) { _ = "STUB: not implemented"; return }

func (s *resourceScope) ReserveMemoryForChild(size int64, prio uint8) (network.ScopeStat, error) {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat), nil
}

func (s *resourceScope) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (s *resourceScope) ReleaseMemoryForChild(size int64) { _ = "STUB: not implemented"; return }

func (s *resourceScope) AddStream(dir network.Direction) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) addStreamForEdges(dir network.Direction) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) AddStreamForChild(dir network.Direction) (network.ScopeStat, error) {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat), nil
}

func (s *resourceScope) RemoveStream(dir network.Direction) { _ = "STUB: not implemented"; return }

func (s *resourceScope) removeStreamForEdges(dir network.Direction) {
	_ = "STUB: not implemented"
	return
}

func (s *resourceScope) RemoveStreamForChild(dir network.Direction) {
	_ = "STUB: not implemented"
	return
}

func (s *resourceScope) AddConn(dir network.Direction, usefd bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) addConnForEdges(dir network.Direction, usefd bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) AddConnForChild(dir network.Direction, usefd bool) (network.ScopeStat, error) {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat), nil
}

func (s *resourceScope) RemoveConn(dir network.Direction, usefd bool) {
	_ = "STUB: not implemented"
	return
}

func (s *resourceScope) removeConnForEdges(dir network.Direction, usefd bool) {
	_ = "STUB: not implemented"
	return
}

func (s *resourceScope) RemoveConnForChild(dir network.Direction, usefd bool) {
	_ = "STUB: not implemented"
	return
}

func (s *resourceScope) ReserveForChild(st network.ScopeStat) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *resourceScope) ReleaseForChild(st network.ScopeStat) { _ = "STUB: not implemented"; return }

func (s *resourceScope) ReleaseResources(st network.ScopeStat) { _ = "STUB: not implemented"; return }

func (s *resourceScope) nextSpanID() int { _ = "STUB: not implemented"; return 0 }

func (s *resourceScope) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}

func (s *resourceScope) Done() { _ = "STUB: not implemented"; return }

func (s *resourceScope) doneUnlocked() { _ = "STUB: not implemented"; return }

func (s *resourceScope) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (s *resourceScope) IncRef() { _ = "STUB: not implemented"; return }

func (s *resourceScope) DecRef() { _ = "STUB: not implemented"; return }

func (s *resourceScope) IsUnused() bool { _ = "STUB: not implemented"; return false }
