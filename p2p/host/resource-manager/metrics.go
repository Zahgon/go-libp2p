package rcmgr

import (
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type MetricsReporter interface {
	AllowConn(dir network.Direction, usefd bool)

	BlockConn(dir network.Direction, usefd bool)

	AllowStream(p peer.ID, dir network.Direction)

	BlockStream(p peer.ID, dir network.Direction)

	AllowPeer(p peer.ID)

	BlockPeer(p peer.ID)

	AllowProtocol(proto protocol.ID)

	BlockProtocol(proto protocol.ID)

	BlockProtocolPeer(proto protocol.ID, p peer.ID)

	AllowService(svc string)

	BlockService(svc string)

	BlockServicePeer(svc string, p peer.ID)

	AllowMemory(size int)

	BlockMemory(size int)
}

type metrics struct {
	reporter MetricsReporter
}

func WithMetrics(reporter MetricsReporter) Option { _ = "STUB: not implemented"; return *new(Option) }

func (m *metrics) AllowConn(dir network.Direction, usefd bool) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockConn(dir network.Direction, usefd bool) { _ = "STUB: not implemented"; return }

func (m *metrics) AllowStream(p peer.ID, dir network.Direction) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockStream(p peer.ID, dir network.Direction) { _ = "STUB: not implemented"; return }

func (m *metrics) AllowPeer(p peer.ID) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockPeer(p peer.ID) { _ = "STUB: not implemented"; return }

func (m *metrics) AllowProtocol(proto protocol.ID) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockProtocol(proto protocol.ID) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockProtocolPeer(proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

func (m *metrics) AllowService(svc string) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockService(svc string) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockServicePeer(svc string, p peer.ID) { _ = "STUB: not implemented"; return }

func (m *metrics) AllowMemory(size int) { _ = "STUB: not implemented"; return }

func (m *metrics) BlockMemory(size int) { _ = "STUB: not implemented"; return }
