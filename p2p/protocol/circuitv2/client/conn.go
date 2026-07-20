package client

import (
	"net"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var HopTagWeight = 5

type statLimitDuration struct{}
type statLimitData struct{}

var (
	StatLimitDuration = statLimitDuration{}
	StatLimitData     = statLimitData{}
)

type Conn struct {
	stream network.Stream
	remote peer.AddrInfo
	stat   network.ConnStats

	client *Client
}

type NetAddr struct {
	Relay  string
	Remote string
}

var _ net.Addr = (*NetAddr)(nil)

func (n *NetAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (n *NetAddr) String() string { _ = "STUB: not implemented"; return "" }

var _ manet.Conn = (*Conn)(nil)

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) RemoteMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *Conn) LocalMultiaddr() ma.Multiaddr { _ = "STUB: not implemented"; return *new(ma.Multiaddr) }

func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

var _ network.ConnStat = (*Conn)(nil)

func (c *Conn) Stat() network.ConnStats { _ = "STUB: not implemented"; return *new(network.ConnStats) }

func (c *Conn) tagHop() { _ = "STUB: not implemented"; return }

func (c *Conn) untagHop() { _ = "STUB: not implemented"; return }

type capableConnWithStat interface {
	tpt.CapableConn
	network.ConnStat
}

type capableConn struct {
	capableConnWithStat
}

var transportName = ma.ProtocolWithCode(ma.P_CIRCUIT).Name

func (c capableConn) ConnState() network.ConnectionState {
	_ = "STUB: not implemented"
	return *new(network.ConnectionState)
}
