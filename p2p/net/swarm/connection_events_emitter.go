package swarm

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

type peerConnectednessEventType int

const (
	removeConnEvent peerConnectednessEventType = iota
	addConnEvent
)

type peerConnectednessEvent struct {
	PeerID peer.ID
	Type   peerConnectednessEventType
}

type connectionEventsEmitter struct {
	peerConnectednessCh chan peerConnectednessEvent

	lastConnectednessEvent map[peer.ID]network.Connectedness

	connectedness func(peer.ID) network.Connectedness

	onConnected func(*Conn)

	onDisconnected func(*Conn)

	emitter event.Emitter

	closeMu sync.Mutex
	closed  bool
	wg      sync.WaitGroup

	loopWG sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	notifsLk          sync.Mutex
	connected         map[*Conn]struct{}
	pendingDisconnect map[*Conn]struct{}
}

func newConnectionEventsEmitter(
	connectedness func(peer.ID) network.Connectedness,
	emitter event.Emitter,
	onConnected func(*Conn),
	onDisconnected func(*Conn),
) *connectionEventsEmitter {
	_ = "STUB: not implemented"
	return nil
}

func (c *connectionEventsEmitter) AddConn(conn *Conn) { _ = "STUB: not implemented"; return }

func (c *connectionEventsEmitter) RemoveConn(conn *Conn) { _ = "STUB: not implemented"; return }

func (c *connectionEventsEmitter) Close() { _ = "STUB: not implemented"; return }

func (c *connectionEventsEmitter) runEmitter() { _ = "STUB: not implemented"; return }

func (c *connectionEventsEmitter) notifyPeer(pce peerConnectednessEvent) {
	_ = "STUB: not implemented"
	return
}
