package mocknetwork

import (
	network "github.com/libp2p/go-libp2p/core/network"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
	gomock "go.uber.org/mock/gomock"
)

type MockProtocolScope struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolScopeMockRecorder
	isgomock struct{}
}

type MockProtocolScopeMockRecorder struct {
	mock *MockProtocolScope
}

func NewMockProtocolScope(ctrl *gomock.Controller) *MockProtocolScope {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockProtocolScope) EXPECT() *MockProtocolScopeMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockProtocolScope) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}

func (mr *MockProtocolScopeMockRecorder) BeginSpan() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockProtocolScope) Protocol() protocol.ID {
	_ = "STUB: not implemented"
	return *new(protocol.ID)
}

func (mr *MockProtocolScopeMockRecorder) Protocol() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockProtocolScope) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (mr *MockProtocolScopeMockRecorder) ReleaseMemory(size any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockProtocolScope) ReserveMemory(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockProtocolScopeMockRecorder) ReserveMemory(size, prio any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockProtocolScope) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (mr *MockProtocolScopeMockRecorder) Stat() *gomock.Call { _ = "STUB: not implemented"; return nil }
