package rcmgr

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type ResourceScopeLimiter interface {
	Limit() Limit
	SetLimit(Limit)
}

var _ ResourceScopeLimiter = (*resourceScope)(nil)

type ResourceManagerState interface {
	ListServices() []string
	ListProtocols() []protocol.ID
	ListPeers() []peer.ID

	Stat() ResourceManagerStat
}

type ResourceManagerStat struct {
	System    network.ScopeStat
	Transient network.ScopeStat
	Services  map[string]network.ScopeStat
	Protocols map[protocol.ID]network.ScopeStat
	Peers     map[peer.ID]network.ScopeStat
}

var _ ResourceManagerState = (*resourceManager)(nil)

func (s *resourceScope) Limit() Limit { _ = "STUB: not implemented"; return *new(Limit) }

func (s *resourceScope) SetLimit(limit Limit) { _ = "STUB: not implemented"; return }

func (s *protocolScope) SetLimit(limit Limit) { _ = "STUB: not implemented"; return }

func (s *peerScope) SetLimit(limit Limit) { _ = "STUB: not implemented"; return }

func (r *resourceManager) ListServices() []string { _ = "STUB: not implemented"; return nil }

func (r *resourceManager) ListProtocols() []protocol.ID { _ = "STUB: not implemented"; return nil }

func (r *resourceManager) ListPeers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (r *resourceManager) Stat() (result ResourceManagerStat) {
	_ = "STUB: not implemented"
	return *new(ResourceManagerStat)
}

func (r *resourceManager) GetConnLimit() int { _ = "STUB: not implemented"; return 0 }
