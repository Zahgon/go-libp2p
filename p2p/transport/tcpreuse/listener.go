package tcpreuse

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/transport"
	logging "github.com/libp2p/go-libp2p/gologshim"
	"github.com/libp2p/go-libp2p/p2p/net/reuseport"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

const acceptQueueSize = 64

const acceptTimeout = 30 * time.Second

var log = logging.Logger("tcp-demultiplex")

type ConnMgr struct {
	enableReuseport bool
	reuse           reuseport.Transport
	upgrader        transport.Upgrader

	mx        sync.Mutex
	listeners map[string]*multiplexedListener
}

func NewConnMgr(enableReuseport bool, upgrader transport.Upgrader) *ConnMgr {
	_ = "STUB: not implemented"
	return nil
}

func (t *ConnMgr) gatedMaListen(listenAddr ma.Multiaddr) (transport.GatedMaListener, error) {
	_ = "STUB: not implemented"
	return *new(transport.GatedMaListener), nil
}

func (t *ConnMgr) useReuseport() bool { _ = "STUB: not implemented"; return false }

func getTCPAddr(listenAddr ma.Multiaddr) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func (t *ConnMgr) DemultiplexedListen(laddr ma.Multiaddr, connType DemultiplexedConnType) (transport.GatedMaListener, error) {
	_ = "STUB: not implemented"
	return *new(transport.GatedMaListener), nil
}

var _ transport.GatedMaListener = &demultiplexedListener{}

type multiplexedListener struct {
	transport.GatedMaListener
	listeners map[DemultiplexedConnType]*demultiplexedListener
	mx        sync.RWMutex

	ctx     context.Context
	closeFn func() error
	wg      sync.WaitGroup
}

var ErrListenerExists = errors.New("listener already exists for this conn type on this address")

func (m *multiplexedListener) DemultiplexedListen(connType DemultiplexedConnType) (transport.GatedMaListener, error) {
	_ = "STUB: not implemented"
	return *new(transport.GatedMaListener), nil
}

func (m *multiplexedListener) run() error { _ = "STUB: not implemented"; return nil }

func (m *multiplexedListener) Close() error { _ = "STUB: not implemented"; return nil }

func (m *multiplexedListener) closeListener() error { _ = "STUB: not implemented"; return nil }

func (m *multiplexedListener) removeDemultiplexedListener(c DemultiplexedConnType) {
	_ = "STUB: not implemented"
	return
}

type demultiplexedListener struct {
	buffer     chan *connWithScope
	inner      transport.GatedMaListener
	ctx        context.Context
	cancelFunc context.CancelFunc
	closeFn    func() error
}

func (m *demultiplexedListener) Accept() (manet.Conn, network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), *new(network.ConnManagementScope), nil
}

func (m *demultiplexedListener) Close() error { _ = "STUB: not implemented"; return nil }

func (m *demultiplexedListener) Multiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

func (m *demultiplexedListener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
