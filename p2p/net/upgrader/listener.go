package upgrader

import (
	"context"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/transport"

	logging "github.com/libp2p/go-libp2p/gologshim"
	manet "github.com/multiformats/go-multiaddr/net"
)

var log = logging.Logger("upgrader")

type listener struct {
	transport.GatedMaListener

	transport transport.Transport
	upgrader  *upgrader
	rcmgr     network.ResourceManager

	incoming chan transport.CapableConn
	err      error

	threshold *threshold

	ctx    context.Context
	cancel func()
}

var _ transport.Listener = (*listener)(nil)

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *listener) handleIncoming() { _ = "STUB: not implemented"; return }

func (l *listener) Accept() (transport.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.CapableConn), nil
}

func (l *listener) String() string { _ = "STUB: not implemented"; return "" }

type gatedMaListener struct {
	manet.Listener
	rcmgr     network.ResourceManager
	connGater connmgr.ConnectionGater
}

var _ transport.GatedMaListener = &gatedMaListener{}

func (l *gatedMaListener) Accept() (manet.Conn, network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(manet.Conn), *new(network.ConnManagementScope), nil
}
