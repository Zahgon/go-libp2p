package mocknetwork

import (
	net "net"

	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
	multiaddr "github.com/multiformats/go-multiaddr"
	gomock "go.uber.org/mock/gomock"
)

type MockResourceManager struct {
	ctrl     *gomock.Controller
	recorder *MockResourceManagerMockRecorder
	isgomock struct{}
}

type MockResourceManagerMockRecorder struct {
	mock *MockResourceManager
}

func NewMockResourceManager(ctrl *gomock.Controller) *MockResourceManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) EXPECT() *MockResourceManagerMockRecorder {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) Close() error { _ = "STUB: not implemented"; return nil }

func (mr *MockResourceManagerMockRecorder) Close() *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) OpenConnection(dir network.Direction, usefd bool, endpoint multiaddr.Multiaddr) (network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(network.ConnManagementScope), nil
}

func (mr *MockResourceManagerMockRecorder) OpenConnection(dir, usefd, endpoint any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) OpenStream(p peer.ID, dir network.Direction) (network.StreamManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(network.StreamManagementScope), nil
}

func (mr *MockResourceManagerMockRecorder) OpenStream(p, dir any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) VerifySourceAddress(addr net.Addr) bool {
	_ = "STUB: not implemented"
	return false
}

func (mr *MockResourceManagerMockRecorder) VerifySourceAddress(addr any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) ViewPeer(arg0 peer.ID, arg1 func(network.PeerScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockResourceManagerMockRecorder) ViewPeer(arg0, arg1 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) ViewProtocol(arg0 protocol.ID, arg1 func(network.ProtocolScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockResourceManagerMockRecorder) ViewProtocol(arg0, arg1 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) ViewService(arg0 string, arg1 func(network.ServiceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockResourceManagerMockRecorder) ViewService(arg0, arg1 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) ViewSystem(arg0 func(network.ResourceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockResourceManagerMockRecorder) ViewSystem(arg0 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockResourceManager) ViewTransient(arg0 func(network.ResourceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockResourceManagerMockRecorder) ViewTransient(arg0 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}
