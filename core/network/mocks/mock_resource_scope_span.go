package mocknetwork

import (
	network "github.com/libp2p/go-libp2p/core/network"
	gomock "go.uber.org/mock/gomock"
)

type MockResourceScopeSpan struct {
	ctrl     *gomock.Controller
	recorder *MockResourceScopeSpanMockRecorder
	isgomock struct{}
}

type MockResourceScopeSpanMockRecorder struct {
	mock *MockResourceScopeSpan
}

func NewMockResourceScopeSpan(ctrl *gomock.Controller) *MockResourceScopeSpan {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceScopeSpan) EXPECT() *MockResourceScopeSpanMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceScopeSpan) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}

func (mr *MockResourceScopeSpanMockRecorder) BeginSpan() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceScopeSpan) Done() { _ = "STUB: not implemented"; return }

func (mr *MockResourceScopeSpanMockRecorder) Done() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceScopeSpan) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (mr *MockResourceScopeSpanMockRecorder) ReleaseMemory(size any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceScopeSpan) ReserveMemory(size int, prio uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockResourceScopeSpanMockRecorder) ReserveMemory(size, prio any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceScopeSpan) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}

func (mr *MockResourceScopeSpanMockRecorder) Stat() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
