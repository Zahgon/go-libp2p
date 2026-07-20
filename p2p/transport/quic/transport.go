package libp2pquic

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/pnet"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	p2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"github.com/libp2p/go-libp2p/p2p/transport/quicreuse"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
	mafmt "github.com/multiformats/go-multiaddr-fmt"
	"github.com/quic-go/quic-go"
)

const ListenOrder = 1

var log = logging.Logger("quic-transport")

var ErrHolePunching = errors.New("hole punching attempted; no active dial")

var HolePunchTimeout = 5 * time.Second

type transport struct {
	privKey     ic.PrivKey
	localPeer   peer.ID
	identity    *p2ptls.Identity
	connManager *quicreuse.ConnManager
	gater       connmgr.ConnectionGater
	rcmgr       network.ResourceManager

	holePunchingMx sync.Mutex
	holePunching   map[holePunchKey]*activeHolePunch

	rndMx sync.Mutex
	rnd   rand.Rand

	connMx sync.Mutex
	conns  map[*quic.Conn]*conn

	listenersMu sync.Mutex

	listeners map[string][]*virtualListener
}

var _ tpt.Transport = &transport{}

type holePunchKey struct {
	addr string
	peer peer.ID
}

type activeHolePunch struct {
	connCh    chan tpt.CapableConn
	fulfilled bool
}

func NewTransport(key ic.PrivKey, connManager *quicreuse.ConnManager, psk pnet.PSK, gater connmgr.ConnectionGater, rcmgr network.ResourceManager, opts ...Option) (tpt.Transport, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Transport), nil
}

func (t *transport) ListenOrder() int { _ = "STUB: not implemented"; return 0 }

func (t *transport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (_c tpt.CapableConn, _err error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (t *transport) dialWithScope(ctx context.Context, raddr ma.Multiaddr, p peer.ID, scope network.ConnManagementScope) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (t *transport) addConn(conn *quic.Conn, c *conn) { _ = "STUB: not implemented"; return }

func (t *transport) removeConn(conn *quic.Conn) { _ = "STUB: not implemented"; return }

func (t *transport) holePunch(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

var dialMatcher = mafmt.And(mafmt.IP, mafmt.Base(ma.P_UDP), mafmt.Base(ma.P_QUIC_V1))

func (t *transport) CanDial(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (t *transport) Listen(addr ma.Multiaddr) (tpt.Listener, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Listener), nil
}

func (t *transport) allowWindowIncrease(conn *quic.Conn, size uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *transport) Proxy() bool { _ = "STUB: not implemented"; return false }

func (t *transport) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (t *transport) String() string { _ = "STUB: not implemented"; return "" }

func (t *transport) Close() error { _ = "STUB: not implemented"; return nil }

func (t *transport) CloseVirtualListener(l *virtualListener) error {
	_ = "STUB: not implemented"
	return nil
}
