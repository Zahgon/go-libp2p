package rcmgr

import (
	"context"
	"net"
	"net/netip"
	"sync"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/x/rate"

	logging "github.com/libp2p/go-libp2p/gologshim"
	"github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("rcmgr")

type resourceManager struct {
	limits Limiter

	connLimiter                    *connLimiter
	connRateLimiter                *rate.Limiter
	verifySourceAddressRateLimiter *rate.Limiter

	trace          *trace
	metrics        *metrics
	disableMetrics bool

	allowlist *Allowlist

	system    *systemScope
	transient *transientScope

	allowlistedSystem    *systemScope
	allowlistedTransient *transientScope

	cancelCtx context.Context
	cancel    func()
	wg        sync.WaitGroup

	mx    sync.Mutex
	svc   map[string]*serviceScope
	proto map[protocol.ID]*protocolScope
	peer  map[peer.ID]*peerScope

	stickyProto map[protocol.ID]struct{}
	stickyPeer  map[peer.ID]struct{}

	connId, streamId int64
}

var _ network.ResourceManager = (*resourceManager)(nil)

type systemScope struct {
	*resourceScope
}

var _ network.ResourceScope = (*systemScope)(nil)

type transientScope struct {
	*resourceScope

	system *systemScope
}

var _ network.ResourceScope = (*transientScope)(nil)

type serviceScope struct {
	*resourceScope

	service string
	rcmgr   *resourceManager

	peers map[peer.ID]*resourceScope
}

var _ network.ServiceScope = (*serviceScope)(nil)

type protocolScope struct {
	*resourceScope

	proto protocol.ID
	rcmgr *resourceManager

	peers map[peer.ID]*resourceScope
}

var _ network.ProtocolScope = (*protocolScope)(nil)

type peerScope struct {
	*resourceScope

	peer  peer.ID
	rcmgr *resourceManager
}

var _ network.PeerScope = (*peerScope)(nil)

type connectionScope struct {
	*resourceScope

	dir           network.Direction
	usefd         bool
	isAllowlisted bool
	rcmgr         *resourceManager
	peer          *peerScope
	endpoint      multiaddr.Multiaddr
	ip            netip.Addr
}

var _ network.ConnScope = (*connectionScope)(nil)
var _ network.ConnManagementScope = (*connectionScope)(nil)

type streamScope struct {
	*resourceScope

	dir   network.Direction
	rcmgr *resourceManager
	peer  *peerScope
	svc   *serviceScope
	proto *protocolScope

	peerProtoScope *resourceScope
	peerSvcScope   *resourceScope
}

var _ network.StreamScope = (*streamScope)(nil)
var _ network.StreamManagementScope = (*streamScope)(nil)

type Option func(*resourceManager) error

func NewResourceManager(limits Limiter, opts ...Option) (network.ResourceManager, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceManager), nil
}

func (r *resourceManager) GetAllowlist() *Allowlist { _ = "STUB: not implemented"; return nil }

func GetAllowlist(rcmgr network.ResourceManager) *Allowlist { _ = "STUB: not implemented"; return nil }

func (r *resourceManager) ViewSystem(f func(network.ResourceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) ViewTransient(f func(network.ResourceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) ViewService(srv string, f func(network.ServiceScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) ViewProtocol(proto protocol.ID, f func(network.ProtocolScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) ViewPeer(p peer.ID, f func(network.PeerScope) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) getServiceScope(svc string) *serviceScope {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) getProtocolScope(proto protocol.ID) *protocolScope {
	_ = "STUB: not implemented"
	return nil
}

func (r *resourceManager) setStickyProtocol(proto protocol.ID) { _ = "STUB: not implemented"; return }

func (r *resourceManager) getPeerScope(p peer.ID) *peerScope { _ = "STUB: not implemented"; return nil }

func (r *resourceManager) setStickyPeer(p peer.ID) { _ = "STUB: not implemented"; return }

func (r *resourceManager) nextConnId() int64 { _ = "STUB: not implemented"; return 0 }

func (r *resourceManager) nextStreamId() int64 { _ = "STUB: not implemented"; return 0 }

func (r *resourceManager) VerifySourceAddress(addr net.Addr) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *resourceManager) OpenConnectionNoIP(dir network.Direction, usefd bool, endpoint multiaddr.Multiaddr) (network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(network.ConnManagementScope), nil
}

func (r *resourceManager) OpenConnection(dir network.Direction, usefd bool, endpoint multiaddr.Multiaddr) (network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(network.ConnManagementScope), nil
}

func (r *resourceManager) openConnection(dir network.Direction, usefd bool, endpoint multiaddr.Multiaddr, ip netip.Addr) (network.ConnManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(network.ConnManagementScope), nil
}

func (r *resourceManager) OpenStream(p peer.ID, dir network.Direction) (network.StreamManagementScope, error) {
	_ = "STUB: not implemented"
	return *new(network.StreamManagementScope), nil
}

func (r *resourceManager) Close() error { _ = "STUB: not implemented"; return nil }

func (r *resourceManager) background() { _ = "STUB: not implemented"; return }

func (r *resourceManager) gc() { _ = "STUB: not implemented"; return }

func newSystemScope(limit Limit, rcmgr *resourceManager, name string) *systemScope {
	_ = "STUB: not implemented"
	return nil
}

func newTransientScope(limit Limit, rcmgr *resourceManager, name string, systemScope *resourceScope) *transientScope {
	_ = "STUB: not implemented"
	return nil
}

func newServiceScope(service string, limit Limit, rcmgr *resourceManager) *serviceScope {
	_ = "STUB: not implemented"
	return nil
}

func newProtocolScope(proto protocol.ID, limit Limit, rcmgr *resourceManager) *protocolScope {
	_ = "STUB: not implemented"
	return nil
}

func newPeerScope(p peer.ID, limit Limit, rcmgr *resourceManager) *peerScope {
	_ = "STUB: not implemented"
	return nil
}

func newConnectionScope(dir network.Direction, usefd bool, limit Limit, rcmgr *resourceManager, endpoint multiaddr.Multiaddr, ip netip.Addr) *connectionScope {
	_ = "STUB: not implemented"
	return nil
}

func newAllowListedConnectionScope(dir network.Direction, usefd bool, limit Limit, rcmgr *resourceManager, endpoint multiaddr.Multiaddr) *connectionScope {
	_ = "STUB: not implemented"
	return nil
}

func newStreamScope(dir network.Direction, limit Limit, peer *peerScope, rcmgr *resourceManager) *streamScope {
	_ = "STUB: not implemented"
	return nil
}

func IsSystemScope(name string) bool { _ = "STUB: not implemented"; return false }

func IsTransientScope(name string) bool { _ = "STUB: not implemented"; return false }

func streamScopeName(streamId int64) string { _ = "STUB: not implemented"; return "" }

func IsStreamScope(name string) bool { _ = "STUB: not implemented"; return false }

func connScopeName(streamId int64) string { _ = "STUB: not implemented"; return "" }

func IsConnScope(name string) bool { _ = "STUB: not implemented"; return false }

func peerScopeName(p peer.ID) string { _ = "STUB: not implemented"; return "" }

func PeerStrInScopeName(name string) string { _ = "STUB: not implemented"; return "" }

func ParseProtocolScopeName(name string) string { _ = "STUB: not implemented"; return "" }

func (s *serviceScope) Name() string { _ = "STUB: not implemented"; return "" }

func (s *serviceScope) getPeerScope(p peer.ID) *resourceScope {
	_ = "STUB: not implemented"
	return nil
}

func (s *protocolScope) Protocol() protocol.ID { _ = "STUB: not implemented"; return *new(protocol.ID) }

func (s *protocolScope) getPeerScope(p peer.ID) *resourceScope {
	_ = "STUB: not implemented"
	return nil
}

func (s *peerScope) Peer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (s *connectionScope) PeerScope() network.PeerScope {
	_ = "STUB: not implemented"
	return *new(network.PeerScope)
}

func (s *connectionScope) Done() { _ = "STUB: not implemented"; return }

func (s *connectionScope) transferAllowedToStandard() (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *connectionScope) SetPeer(p peer.ID) error { _ = "STUB: not implemented"; return nil }

func (s *streamScope) ProtocolScope() network.ProtocolScope {
	_ = "STUB: not implemented"
	return *new(network.ProtocolScope)
}

func (s *streamScope) SetProtocol(proto protocol.ID) error { _ = "STUB: not implemented"; return nil }

func (s *streamScope) ServiceScope() network.ServiceScope {
	_ = "STUB: not implemented"
	return *new(network.ServiceScope)
}

func (s *streamScope) SetService(svc string) error { _ = "STUB: not implemented"; return nil }

func (s *streamScope) PeerScope() network.PeerScope {
	_ = "STUB: not implemented"
	return *new(network.PeerScope)
}
