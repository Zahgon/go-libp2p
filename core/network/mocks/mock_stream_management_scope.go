package mocknetwork

import (
	network "github.com/libp2p/go-libp2p/core/network"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
	gomock "go.uber.org/mock/gomock"
)

type MockStreamManagementScope struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagementScopeMockRecorder
	isgomock struct{}
}

type MockStreamManagementScopeMockRecorder struct {
	mock *MockStreamManagementScope
}

func NewMockStreamManagementScope(ctrl *gomock.Controller) *MockStreamManagementScope {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) EXPECT() *MockStreamManagementScopeMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}

func (mr *MockStreamManagementScopeMockRecorder) BeginSpan() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) Done() { _ = "STUB: not implemented"; return }

func (mr *MockStreamManagementScopeMockRecorder) Done() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) PeerScope() network.PeerScope {
	_ = "STUB: not implemented"
	return *new(network.PeerScope)
}

func (mr *MockStreamManagementScopeMockRecorder) PeerScope() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) ProtocolScope() network.ProtocolScope {
	_ = "STUB: not implemented"
	return *new(network.ProtocolScope)
}

func (mr *MockStreamManagementScopeMockRecorder) ProtocolScope() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (mr *MockStreamManagementScopeMockRecorder) ReleaseMemory(size any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) ReserveMemory(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockStreamManagementScopeMockRecorder) ReserveMemory(size, prio any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) ServiceScope() network.ServiceScope {
	_ = "STUB: not implemented"
	return *new(network.ServiceScope)
}

func (mr *MockStreamManagementScopeMockRecorder) ServiceScope() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) SetProtocol(proto protocol.ID) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockStreamManagementScopeMockRecorder) SetProtocol(proto any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) SetService(srv string) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockStreamManagementScopeMockRecorder) SetService(srv any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockStreamManagementScope) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (mr *MockStreamManagementScopeMockRecorder) Stat() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
