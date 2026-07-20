package network

import (
	ma "github.com/multiformats/go-multiaddr"
)

type Notifiee interface {
	Listen(Network, ma.Multiaddr)
	ListenClose(Network, ma.Multiaddr)
	Connected(Network, Conn)
	Disconnected(Network, Conn)
}

type NotifyBundle struct {
	ListenF      func(Network, ma.Multiaddr)
	ListenCloseF func(Network, ma.Multiaddr)

	ConnectedF    func(Network, Conn)
	DisconnectedF func(Network, Conn)
}

var _ Notifiee = (*NotifyBundle)(nil)

func (nb *NotifyBundle) Listen(n Network, a ma.Multiaddr) { _ = "STUB: not implemented"; return }

func (nb *NotifyBundle) ListenClose(n Network, a ma.Multiaddr) { _ = "STUB: not implemented"; return }

func (nb *NotifyBundle) Connected(n Network, c Conn) { _ = "STUB: not implemented"; return }

func (nb *NotifyBundle) Disconnected(n Network, c Conn) { _ = "STUB: not implemented"; return }

var GlobalNoopNotifiee = &NoopNotifiee{}

type NoopNotifiee struct{}

var _ Notifiee = (*NoopNotifiee)(nil)

func (nn *NoopNotifiee) Connected(_ Network, _ Conn)           { _ = "STUB: not implemented"; return }
func (nn *NoopNotifiee) Disconnected(_ Network, _ Conn)        { _ = "STUB: not implemented"; return }
func (nn *NoopNotifiee) Listen(_ Network, _ ma.Multiaddr)      { _ = "STUB: not implemented"; return }
func (nn *NoopNotifiee) ListenClose(_ Network, _ ma.Multiaddr) { _ = "STUB: not implemented"; return }
