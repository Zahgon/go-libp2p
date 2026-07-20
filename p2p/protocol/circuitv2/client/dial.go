package client

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

const maxMessageSize = 4096

var DialTimeout = time.Minute
var DialRelayTimeout = 5 * time.Second

type relayError struct {
	err string
}

func (e relayError) Error() string { _ = "STUB: not implemented"; return "" }

func newRelayError(t string, args ...any) error { _ = "STUB: not implemented"; return nil }

func isRelayError(err error) bool { _ = "STUB: not implemented"; return false }

func (c *Client) dial(ctx context.Context, a ma.Multiaddr, p peer.ID) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) dialPeer(ctx context.Context, relay, dest peer.AddrInfo) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) connect(s network.Stream, dest peer.AddrInfo) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
