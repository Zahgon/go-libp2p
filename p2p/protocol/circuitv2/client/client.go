package client

import (
	"context"
	"io"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("p2p-circuit")

type Client struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	host      host.Host
	upgrader  transport.Upgrader

	incoming chan accept

	mx          sync.Mutex
	activeDials map[peer.ID]*completion
	hopCount    map[peer.ID]int
}

var _ io.Closer = &Client{}
var _ transport.Transport = &Client{}

type accept struct {
	conn          *Conn
	writeResponse func() error
}

type completion struct {
	ch    chan struct{}
	relay peer.ID
	err   error
}

func New(h host.Host, upgrader transport.Upgrader) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Start() { _ = "STUB: not implemented"; return }

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }
