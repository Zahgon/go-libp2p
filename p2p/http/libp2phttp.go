package libp2phttp

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	host "github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	logging "github.com/libp2p/go-libp2p/gologshim"
	httpauth "github.com/libp2p/go-libp2p/p2p/http/auth"
	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("libp2phttp")

var WellKnownRequestTimeout = 30 * time.Second

const ProtocolIDForMultistreamSelect = "/http/1.1"
const WellKnownProtocols = "/.well-known/libp2p/protocols"

const LegacyWellKnownProtocols = "/.well-known/libp2p"

const peerMetadataLimit = 8 << 10
const peerMetadataLRUSize = 256

var DefaultNewStreamTimeout = 10 * time.Second

type clientPeerIDContextKey struct{}
type serverPeerIDContextKey struct{}

func ClientPeerID(r *http.Request) peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func ServerPeerID(r *http.Response) peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

type ProtocolMeta struct {
	Path string `json:"path"`
}

type PeerMeta map[protocol.ID]ProtocolMeta

type WellKnownHandler struct {
	wellknownMapMu   sync.Mutex
	wellKnownMapping PeerMeta
	wellKnownCache   []byte
}

func streamHostListen(streamHost host.Host) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (h *WellKnownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *WellKnownHandler) AddProtocolMeta(p protocol.ID, protocolMeta ProtocolMeta) {
	_ = "STUB: not implemented"
	return
}

func (h *WellKnownHandler) RemoveProtocolMeta(p protocol.ID) { _ = "STUB: not implemented"; return }

type Host struct {
	StreamHost host.Host

	ListenAddrs []ma.Multiaddr

	TLSConfig *tls.Config

	InsecureAllowHTTP bool

	ServerPeerIDAuth *httpauth.ServerPeerIDAuth

	ClientPeerIDAuth *httpauth.ClientPeerIDAuth

	ServeMux           *http.ServeMux
	initializeServeMux sync.Once

	DefaultClientRoundTripper *http.Transport

	WellKnownHandler WellKnownHandler

	EnableCompatibilityWithLegacyWellKnownEndpoint bool

	peerMetadata *lru.Cache[peer.ID, PeerMeta]

	createHTTPTransport sync.Once

	createDefaultClientRoundTripper sync.Once
	httpTransport                   *httpTransport
}

type httpTransport struct {
	listenAddrs         []ma.Multiaddr
	listeners           []net.Listener
	closeListeners      chan struct{}
	waitingForListeners chan struct{}
}

func newPeerMetadataCache() *lru.Cache[peer.ID, PeerMeta] { _ = "STUB: not implemented"; return nil }

func (h *Host) httpTransportInit() { _ = "STUB: not implemented"; return }

func (h *Host) serveMuxInit() { _ = "STUB: not implemented"; return }

func (h *Host) Addrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (h *Host) PeerID() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

var ErrNoListeners = errors.New("nothing to listen on")

func (h *Host) setupListeners(listenerErrCh chan error) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Host) Serve() error { _ = "STUB: not implemented"; return nil }

func (h *Host) Close() error { _ = "STUB: not implemented"; return nil }

func (h *Host) SetHTTPHandler(p protocol.ID, handler http.Handler) {
	_ = "STUB: not implemented"
	return
}

func (h *Host) SetHTTPHandlerAtPath(p protocol.ID, path string, handler http.Handler) {
	_ = "STUB: not implemented"
	return
}

type PeerMetadataGetter interface {
	GetPeerMetadata() (PeerMeta, error)
}

type streamRoundTripper struct {
	server peer.ID

	skipAddAddrs bool
	addrsAdded   sync.Once
	serverAddrs  []ma.Multiaddr
	h            host.Host
	httpHost     *Host
}

type streamReadCloser struct {
	io.ReadCloser
	s network.Stream
}

func (s *streamReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (rt *streamRoundTripper) GetPeerMetadata() (PeerMeta, error) {
	_ = "STUB: not implemented"
	return *new(PeerMeta), nil
}

func (rt *streamRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func locationHeaderToMultiaddrURI(original *url.URL, locationHeader string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type roundTripperForSpecificServer struct {
	http.RoundTripper
	ownRoundtripper  bool
	httpHost         *Host
	server           peer.ID
	targetServerAddr string
	sni              string
	scheme           string
	cachedProtos     PeerMeta
}

func (rt *roundTripperForSpecificServer) GetPeerMetadata() (PeerMeta, error) {
	_ = "STUB: not implemented"
	return *new(PeerMeta), nil
}

func (rt *roundTripperForSpecificServer) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rt *roundTripperForSpecificServer) CloseIdleConnections() { _ = "STUB: not implemented"; return }

type namespacedRoundTripper struct {
	http.RoundTripper
	protocolPrefix    string
	protocolPrefixRaw string
}

func (rt *namespacedRoundTripper) GetPeerMetadata() (PeerMeta, error) {
	_ = "STUB: not implemented"
	return *new(PeerMeta), nil
}

func (rt *namespacedRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Host) NamespaceRoundTripper(roundtripper http.RoundTripper, p protocol.ID, server peer.ID) (*namespacedRoundTripper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Host) NamespacedClient(p protocol.ID, server peer.AddrInfo, opts ...RoundTripperOption) (http.Client, error) {
	_ = "STUB: not implemented"
	return *new(http.Client), nil
}

func (h *Host) initDefaultRT() { _ = "STUB: not implemented"; return }

func (h *Host) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Host) NewConstrainedRoundTripper(server peer.AddrInfo, opts ...RoundTripperOption) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

type explodedMultiaddr struct {
	useHTTPS bool
	host     string
	port     string
	sni      string
	httpPath string
	peer     peer.ID
}

func parseMultiaddr(addr ma.Multiaddr) (explodedMultiaddr, error) {
	_ = "STUB: not implemented"
	return *new(explodedMultiaddr), nil
}

var httpComponent, _ = ma.NewComponent("http", "")
var tlsComponent, _ = ma.NewComponent("tls", "")

func normalizeHTTPMultiaddr(addr ma.Multiaddr) (ma.Multiaddr, bool) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), false
}

func (h *Host) getAndStorePeerMetadata(ctx context.Context, roundtripper http.RoundTripper, server peer.ID) (PeerMeta, error) {
	_ = "STUB: not implemented"
	return *new(PeerMeta), nil
}

func requestPeerMeta(ctx context.Context, roundtripper http.RoundTripper, wellKnownResource string) (PeerMeta, error) {
	_ = "STUB: not implemented"
	return *new(PeerMeta), nil
}

func (h *Host) SetPeerMetadata(server peer.ID, meta PeerMeta) { _ = "STUB: not implemented"; return }

func (h *Host) AddPeerMetadata(server peer.ID, meta PeerMeta) { _ = "STUB: not implemented"; return }

func (h *Host) GetPeerMetadata(server peer.ID) (PeerMeta, bool) {
	_ = "STUB: not implemented"
	return *new(PeerMeta), false
}

func (h *Host) RemovePeerMetadata(server peer.ID) { _ = "STUB: not implemented"; return }

func connectionCloseHeaderMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func maybeDecorateContextWithAuthMiddleware(serverAuth *httpauth.ServerPeerIDAuth, next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
