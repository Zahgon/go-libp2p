package mocknetwork

import (
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	gomock "go.uber.org/mock/gomock"
)

type MockConnManagementScope struct {
	ctrl     *gomock.Controller
	recorder *MockConnManagementScopeMockRecorder
	isgomock struct{}
}

type MockConnManagementScopeMockRecorder struct {
	mock *MockConnManagementScope
}

func NewMockConnManagementScope(ctrl *gomock.Controller) *MockConnManagementScope {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) EXPECT() *MockConnManagementScopeMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}

func (mr *MockConnManagementScopeMockRecorder) BeginSpan() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) Done() { _ = "STUB: not implemented"; return }

func (mr *MockConnManagementScopeMockRecorder) Done() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) PeerScope() network.PeerScope {
	_ = "STUB: not implemented"
	return *new(network.PeerScope)
}

func (mr *MockConnManagementScopeMockRecorder) PeerScope() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (mr *MockConnManagementScopeMockRecorder) ReleaseMemory(size any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) ReserveMemory(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockConnManagementScopeMockRecorder) ReserveMemory(size, prio any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) SetPeer(arg0 peer.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockConnManagementScopeMockRecorder) SetPeer(arg0 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockConnManagementScope) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (mr *MockConnManagementScopeMockRecorder) Stat() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
