package holepunch

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	logging "github.com/libp2p/go-libp2p/gologshim"
	"github.com/libp2p/go-libp2p/p2p/protocol/identify"

	ma "github.com/multiformats/go-multiaddr"
)

const defaultDirectDialTimeout = 10 * time.Second

const Protocol protocol.ID = "/libp2p/dcutr"

var log = logging.Logger("p2p-holepunch")

var StreamTimeout = 1 * time.Minute

const (
	ServiceName = "libp2p.holepunch"

	maxMsgSize = 4 * 1024
)

var ErrClosed = errors.New("hole punching service closing")

type Option func(*Service) error

func DirectDialTimeout(timeout time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type Service struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	host host.Host

	ids identify.IDService

	listenAddrs func() []ma.Multiaddr

	directDialTimeout time.Duration
	holePuncherMx     sync.Mutex
	holePuncher       *holePuncher

	hasPublicAddrsChan chan struct{}

	tracer *tracer
	filter AddrFilter

	refCount sync.WaitGroup
}

func NewService(h host.Host, ids identify.IDService, listenAddrs func() []ma.Multiaddr, opts ...Option) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) waitForPublicAddr() { _ = "STUB: not implemented"; return }

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Service) incomingHolePunch(str network.Stream) (rtt time.Duration, remoteAddrs []ma.Multiaddr, ownAddrs []ma.Multiaddr, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil, nil, nil
}

func (s *Service) handleNewStream(str network.Stream) { _ = "STUB: not implemented"; return }

func (s *Service) DirectConnect(p peer.ID) error { _ = "STUB: not implemented"; return nil }
