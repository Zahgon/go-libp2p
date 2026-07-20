package relay

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	pbv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/pb"

	logging "github.com/libp2p/go-libp2p/gologshim"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	ServiceName = "libp2p.relay/v2"

	ReservationTagWeight = 10

	StreamTimeout    = time.Minute
	ConnectTimeout   = 30 * time.Second
	HandshakeTimeout = time.Minute

	relayHopTag      = "relay-v2-hop"
	relayHopTagValue = 2

	maxMessageSize = 4096
)

var log = logging.Logger("relay")

type Relay struct {
	ctx    context.Context
	cancel func()

	reservationAddrFilter ReservationAddressFilterFunc

	host        host.Host
	rc          Resources
	acl         ACLFilter
	constraints *constraints
	scope       network.ResourceScopeSpan
	notifiee    network.Notifiee

	mx     sync.Mutex
	rsvp   map[peer.ID]time.Time
	conns  map[peer.ID]int
	closed bool

	selfAddr ma.Multiaddr

	metricsTracer MetricsTracer
}

func New(h host.Host, opts ...Option) (*Relay, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Relay) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Relay) handleStream(s network.Stream) { _ = "STUB: not implemented"; return }

func (r *Relay) handleReserve(s network.Stream) pbv2.Status {
	_ = "STUB: not implemented"
	return *new(pbv2.Status)
}

func (r *Relay) handleConnect(s network.Stream, msg *pbv2.HopMessage) pbv2.Status {
	_ = "STUB: not implemented"
	return *new(pbv2.Status)
}

func (r *Relay) addConn(p peer.ID) { _ = "STUB: not implemented"; return }

func (r *Relay) rmConn(p peer.ID) { _ = "STUB: not implemented"; return }

func (r *Relay) relayLimited(src, dest network.Stream, srcID, destID peer.ID, limit int64, done func()) {
	_ = "STUB: not implemented"
	return
}

func (r *Relay) relayUnlimited(src, dest network.Stream, srcID, destID peer.ID, done func()) {
	_ = "STUB: not implemented"
	return
}

var errInvalidWrite = errors.New("invalid write result")

func (r *Relay) copyWithBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *Relay) handleError(s network.Stream, status pbv2.Status) {
	_ = "STUB: not implemented"
	return
}

func (r *Relay) writeResponse(s network.Stream, status pbv2.Status, rsvp *pbv2.Reservation, limit *pbv2.Limit) error {
	_ = "STUB: not implemented"
	return nil
}

func makeReservationMsg(
	reservationAddrFilter ReservationAddressFilterFunc,
	signingKey crypto.PrivKey,
	selfID peer.ID,
	selfAddrs []ma.Multiaddr,
	p peer.ID,
	expire time.Time,
) *pbv2.Reservation {
	_ = "STUB: not implemented"
	return nil
}

func (r *Relay) makeLimitMsg(_ peer.ID) *pbv2.Limit { _ = "STUB: not implemented"; return nil }

func (r *Relay) background() { _ = "STUB: not implemented"; return }

func (r *Relay) gc() { _ = "STUB: not implemented"; return }

func (r *Relay) disconnected(n network.Network, c network.Conn) { _ = "STUB: not implemented"; return }

func isRelayAddr(a ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }
