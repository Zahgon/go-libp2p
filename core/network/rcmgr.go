package network

import (
	"context"
	"net"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type ResourceManager interface {
	ResourceScopeViewer

	OpenConnection(dir Direction, usefd bool, endpoint multiaddr.Multiaddr) (ConnManagementScope, error)

	VerifySourceAddress(addr net.Addr) bool

	OpenStream(p peer.ID, dir Direction) (StreamManagementScope, error)

	Close() error
}

type ResourceScopeViewer interface {
	ViewSystem(func(ResourceScope) error) error

	ViewTransient(func(ResourceScope) error) error

	ViewService(string, func(ServiceScope) error) error

	ViewProtocol(protocol.ID, func(ProtocolScope) error) error

	ViewPeer(peer.ID, func(PeerScope) error) error
}

const (
	ReservationPriorityLow uint8 = 101

	ReservationPriorityMedium uint8 = 152

	ReservationPriorityHigh uint8 = 203

	ReservationPriorityAlways uint8 = 255
)

type ResourceScope interface {
	ReserveMemory(size int, prio uint8) error

	ReleaseMemory(size int)

	Stat() ScopeStat

	BeginSpan() (ResourceScopeSpan, error)
}

type ResourceScopeSpan interface {
	ResourceScope

	Done()
}

type ServiceScope interface {
	ResourceScope

	Name() string
}

type ProtocolScope interface {
	ResourceScope

	Protocol() protocol.ID
}

type PeerScope interface {
	ResourceScope

	Peer() peer.ID
}

type ConnManagementScope interface {
	ResourceScopeSpan

	PeerScope() PeerScope

	SetPeer(peer.ID) error
}

type ConnScope interface {
	ResourceScope
}

type StreamManagementScope interface {
	ResourceScopeSpan

	ProtocolScope() ProtocolScope

	SetProtocol(proto protocol.ID) error

	ServiceScope() ServiceScope

	SetService(srv string) error

	PeerScope() PeerScope
}

type StreamScope interface {
	ResourceScope

	SetService(srv string) error
}

type ScopeStat struct {
	NumStreamsInbound  int
	NumStreamsOutbound int
	NumConnsInbound    int
	NumConnsOutbound   int
	NumFD              int

	Memory int64
}

type connManagementScopeKey struct{}

func WithConnManagementScope(ctx context.Context, scope ConnManagementScope) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func UnwrapConnManagementScope(ctx context.Context) (ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(ConnManagementScope), nil
}

type NullResourceManager struct{}

var _ ResourceManager = (*NullResourceManager)(nil)

var _ ResourceScope = (*NullScope)(nil)
var _ ResourceScopeSpan = (*NullScope)(nil)
var _ ServiceScope = (*NullScope)(nil)
var _ ProtocolScope = (*NullScope)(nil)
var _ PeerScope = (*NullScope)(nil)
var _ ConnManagementScope = (*NullScope)(nil)
var _ ConnScope = (*NullScope)(nil)
var _ StreamManagementScope = (*NullScope)(nil)
var _ StreamScope = (*NullScope)(nil)

type NullScope struct{}

func (n *NullResourceManager) ViewSystem(f func(ResourceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NullResourceManager) ViewTransient(f func(ResourceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NullResourceManager) ViewService(_ string, f func(ServiceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NullResourceManager) ViewProtocol(_ protocol.ID, f func(ProtocolScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NullResourceManager) ViewPeer(_ peer.ID, f func(PeerScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NullResourceManager) OpenConnection(_ Direction, _ bool, _ multiaddr.Multiaddr) (ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(ConnManagementScope), nil
}

func (n *NullResourceManager) OpenStream(_ peer.ID, _ Direction) (StreamManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(StreamManagementScope), nil
}

func (*NullResourceManager) VerifySourceAddress(_ net.Addr) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *NullResourceManager) Close() error { _ = "STUB: not implemented"; return nil }

func (n *NullScope) ReserveMemory(_ int, _ uint8) error { _ = "STUB: not implemented"; return nil }
func (n *NullScope) ReleaseMemory(_ int)                { _ = "STUB: not implemented"; return }
func (n *NullScope) Stat() ScopeStat                    { _ = "STUB: not implemented"; return *new(ScopeStat) }
func (n *NullScope) BeginSpan() (ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(ResourceScopeSpan), nil
}
func (n *NullScope) Done()                 { _ = "STUB: not implemented"; return }
func (n *NullScope) Name() string          { _ = "STUB: not implemented"; return "" }
func (n *NullScope) Protocol() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }
func (n *NullScope) Peer() peer.ID         { _ = "STUB: not implemented"; return *new(peer.ID) }
func (n *NullScope) PeerScope() PeerScope  { _ = "STUB: not implemented"; return *new(PeerScope) }
func (n *NullScope) SetPeer(peer.ID) error { _ = "STUB: not implemented"; return nil }
func (n *NullScope) ProtocolScope() ProtocolScope {
	_ = "STUB: not implemented"
	return *new(ProtocolScope)
}
func (n *NullScope) SetProtocol(_ protocol.ID) error { _ = "STUB: not implemented"; return nil }
func (n *NullScope) ServiceScope() ServiceScope {
	_ = "STUB: not implemented"
	return *new(ServiceScope)
}
func (n *NullScope) SetService(_ string) error           { _ = "STUB: not implemented"; return nil }
func (n *NullScope) VerifySourceAddress(_ net.Addr) bool { _ = "STUB: not implemented"; return false }
