package mocks

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

type clock interface {
	Now() time.Time
}

type MockDiscoveryServer struct {
	mx    sync.Mutex
	db    map[string]map[peer.ID]*discoveryRegistration
	clock clock
}

type discoveryRegistration struct {
	info       peer.AddrInfo
	expiration time.Time
}

func NewDiscoveryServer(clock clock) *MockDiscoveryServer { _ = "STUB: not implemented"; return nil }

func (s *MockDiscoveryServer) Advertise(ns string, info peer.AddrInfo, ttl time.Duration) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (s *MockDiscoveryServer) FindPeers(ns string, limit int) (<-chan peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MockDiscoveryClient struct {
	host   host.Host
	server *MockDiscoveryServer
}

func NewDiscoveryClient(h host.Host, server *MockDiscoveryServer) *MockDiscoveryClient {
	_ = "STUB: not implemented"
	return nil
}

func (d *MockDiscoveryClient) Advertise(_ context.Context, ns string, opts ...discovery.Option) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (d *MockDiscoveryClient) FindPeers(_ context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
