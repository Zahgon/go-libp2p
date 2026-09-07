package quicreuse

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"sync"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
)

type Listener interface {
	Accept(context.Context) (*quic.Conn, error)
	Addr() net.Addr
	Multiaddrs() []ma.Multiaddr
	io.Closer
}

type protoConf struct {
	ln                  *listener
	tlsConf             *tls.Config
	allowWindowIncrease func(conn *quic.Conn, delta uint64) bool
}

type quicListener struct {
	l         QUICListener
	transport RefCountedQUICTransport
	running   chan struct{}
	addrs     []ma.Multiaddr

	protocolsMu sync.Mutex
	protocols   map[string]protoConf
}

func newQuicListener(tr RefCountedQUICTransport, quicConfig *quic.Config) (*quicListener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *quicListener) allowWindowIncrease(conn *quic.Conn, delta uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *quicListener) Add(association any, tlsConf *tls.Config, allowWindowIncrease func(conn *quic.Conn, delta uint64) bool, onRemove func()) (*listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *quicListener) Run() error { _ = "STUB: not implemented"; return nil }

func (l *quicListener) Close() error { _ = "STUB: not implemented"; return nil }

const queueLen = 16

type listener struct {
	queue             chan *quic.Conn
	acceptLoopRunning chan struct{}
	addr              net.Addr
	addrs             []ma.Multiaddr
	remove            func()
	closeOnce         sync.Once
}

var _ Listener = &listener{}

func (l *listener) add(c *quic.Conn) { _ = "STUB: not implemented"; return }

func (l *listener) Accept(ctx context.Context) (*quic.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (l *listener) Multiaddrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }
