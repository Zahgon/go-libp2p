package libp2pwebtransport

import (
	"context"
	"crypto/tls"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p/core/connmgr"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/pnet"
	tpt "github.com/libp2p/go-libp2p/core/transport"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/libp2p/go-libp2p/p2p/transport/quicreuse"

	"github.com/benbjohnson/clock"
	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/webtransport-go"
)

var log = logging.Logger("webtransport")

const webtransportHTTPEndpoint = "/.well-known/libp2p-webtransport"

const errorCodeConnectionGating = 0x47415445

const certValidity = 14 * 24 * time.Hour

type Option func(*transport) error

func WithClock(cl clock.Clock) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLSClientConfig(c *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHandshakeTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type transport struct {
	privKey ic.PrivKey
	pid     peer.ID
	clock   clock.Clock

	connManager *quicreuse.ConnManager
	rcmgr       network.ResourceManager
	gater       connmgr.ConnectionGater

	listenOnce     sync.Once
	listenOnceErr  error
	certManager    *certManager
	hasCertManager atomic.Bool
	staticTLSConf  *tls.Config
	tlsClientConf  *tls.Config

	noise *noise.Transport

	connMx           sync.Mutex
	conns            map[*quic.Conn]*conn
	handshakeTimeout time.Duration
}

var _ tpt.Transport = &transport{}
var _ tpt.Resolver = &transport{}
var _ io.Closer = &transport{}

func New(key ic.PrivKey, psk pnet.PSK, connManager *quicreuse.ConnManager, gater connmgr.ConnectionGater, rcmgr network.ResourceManager, opts ...Option) (tpt.Transport, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Transport), nil
}

func (t *transport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (t *transport) dialWithScope(ctx context.Context, raddr ma.Multiaddr, p peer.ID, scope network.ConnManagementScope) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (t *transport) dial(ctx context.Context, addr ma.Multiaddr, url, sni string, certHashes []multihash.DecodedMultihash) (*webtransport.Session, *quic.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (t *transport) upgrade(ctx context.Context, sess *webtransport.Session, p peer.ID, certHashes []multihash.DecodedMultihash) (*connSecurityMultiaddrs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeCertHashesFromProtobuf(b [][]byte) ([]multihash.DecodedMultihash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *transport) CanDial(addr ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (t *transport) Listen(laddr ma.Multiaddr) (tpt.Listener, error) {
	_ = "STUB: not implemented"
	return *new(tpt.Listener), nil
}

func (t *transport) Protocols() []int { _ = "STUB: not implemented"; return nil }

func (t *transport) Proxy() bool { _ = "STUB: not implemented"; return false }

func (t *transport) Close() error { _ = "STUB: not implemented"; return nil }

func (t *transport) allowWindowIncrease(conn *quic.Conn, size uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *transport) addConn(conn *quic.Conn, c *conn) { _ = "STUB: not implemented"; return }

func (t *transport) removeConn(conn *quic.Conn) { _ = "STUB: not implemented"; return }

func extractSNI(maddr ma.Multiaddr) (sni string, foundSniComponent bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (t *transport) Resolve(_ context.Context, maddr ma.Multiaddr) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *transport) AddCertHashes(m ma.Multiaddr) (ma.Multiaddr, bool) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), false
}
