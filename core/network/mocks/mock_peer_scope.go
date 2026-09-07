package mocknetwork

import (
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	gomock "go.uber.org/mock/gomock"
)

type MockPeerScope struct {
	ctrl     *gomock.Controller
	recorder *MockPeerScopeMockRecorder
	isgomock struct{}
}

type MockPeerScopeMockRecorder struct {
	mock *MockPeerScope
}

func NewMockPeerScope(ctrl *gomock.Controller) *MockPeerScope {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockPeerScope) EXPECT() *MockPeerScopeMockRecorder { _ = "STUB: not implemented"; return nil }

func (m *MockPeerScope) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}

func (mr *MockPeerScopeMockRecorder) BeginSpan() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockPeerScope) Peer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (mr *MockPeerScopeMockRecorder) Peer() *gomock.Call { _ = "STUB: not implemented"; return nil }

func (m *MockPeerScope) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (mr *MockPeerScopeMockRecorder) ReleaseMemory(size any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockPeerScope) ReserveMemory(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockPeerScopeMockRecorder) ReserveMemory(size, prio any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockPeerScope) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (mr *MockPeerScopeMockRecorder) Stat() *gomock.Call { _ = "STUB: not implemented"; return nil }
