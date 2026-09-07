package nat

import (
	"context"
	"errors"
	"net/netip"
	"sync"
	"sync/atomic"
	"time"

	logging "github.com/libp2p/go-libp2p/gologshim"

	"github.com/libp2p/go-libp2p/p2p/net/nat/internal/nat"
)

var ErrNoMapping = errors.New("mapping not established")

var log = logging.Logger("nat")

const MappingDuration = time.Minute

const CacheTime = 15 * time.Second

const DiscoveryTimeout = 10 * time.Second

const rediscoveryThreshold = 3

type entry struct {
	protocol string
	port     int
}

var discoverGateway = nat.DiscoverGateway

func DiscoverNAT(ctx context.Context) (*NAT, error) { _ = "STUB: not implemented"; return nil, nil }

type NAT struct {
	natmu sync.Mutex
	nat   nat.NAT

	consecutiveFailures int
	rediscovering       bool

	extAddr atomic.Pointer[netip.Addr]

	refCount  sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc

	mappingmu sync.RWMutex
	closed    bool
	mappings  map[entry]int
}

func (nat *NAT) Close() error { _ = "STUB: not implemented"; return nil }

func (nat *NAT) GetMapping(protocol string, port int) (addr netip.AddrPort, found bool) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), false
}

func (nat *NAT) AddMapping(ctx context.Context, protocol string, port int) error {
	_ = "STUB: not implemented"
	return nil
}

func (nat *NAT) RemoveMapping(ctx context.Context, protocol string, port int) error {
	_ = "STUB: not implemented"
	return nil
}

func (nat *NAT) background() { _ = "STUB: not implemented"; return }

func (nat *NAT) establishMapping(ctx context.Context, protocol string, internalPort int) (externalPort int) {
	_ = "STUB: not implemented"
	return 0
}

func (nat *NAT) rediscoverNAT() { _ = "STUB: not implemented"; return }

func minTime(a, b time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func getExternalAddress(natInstance nat.NAT) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}
