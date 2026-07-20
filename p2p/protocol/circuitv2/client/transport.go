package client

import (
	"context"
	"io"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

var circuitProtocol = ma.ProtocolWithCode(ma.P_CIRCUIT)
var circuitAddr = ma.Cast(circuitProtocol.VCode)

func AddTransport(h host.Host, upgrader transport.Upgrader) error {
	_ = "STUB: not implemented"
	return nil
}

var _ transport.Transport = (*Client)(nil)

var _ transport.SkipResolver = (*Client)(nil)
var _ io.Closer = (*Client)(nil)

func (c *Client) SkipResolve(_ context.Context, _ ma.Multiaddr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Client) Dial(ctx context.Context, a ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (c *Client) dialAndUpgrade(ctx context.Context, a ma.Multiaddr, p peer.ID, connScope network.ConnManagementScope) (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (c *Client) CanDial(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (c *Client) Listen(addr ma.Multiaddr) (transport.Listener, error) {
	_ = "STUB: not implemented"
	return *new(transport.Listener), nil
}

func (c *Client) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (c *Client) Proxy() bool { _ = "STUB: not implemented"; return false }
