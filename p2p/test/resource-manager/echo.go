package itest

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

const (
	EchoService = "test.echo"
	EchoProtoID = "/test/echo"
)

var (
	echoLog = logging.Logger("echo")
)

type Echo struct {
	Host host.Host

	mx     sync.Mutex
	status EchoStatus

	beforeReserve, beforeRead, beforeWrite, beforeDone func() error
	done                                               func()
}

type EchoStatus struct {
	StreamsIn                 int
	EchosIn, EchosOut         int
	IOErrors                  int
	ResourceServiceErrors     int
	ResourceReservationErrors int
}

func NewEcho(h host.Host) *Echo { _ = "STUB: not implemented"; return nil }

func (e *Echo) Status() EchoStatus { _ = "STUB: not implemented"; return *new(EchoStatus) }

func (e *Echo) BeforeReserve(f func() error) { _ = "STUB: not implemented"; return }

func (e *Echo) BeforeRead(f func() error) { _ = "STUB: not implemented"; return }

func (e *Echo) BeforeWrite(f func() error) { _ = "STUB: not implemented"; return }

func (e *Echo) BeforeDone(f func() error) { _ = "STUB: not implemented"; return }

func (e *Echo) Done(f func()) { _ = "STUB: not implemented"; return }

func (e *Echo) getBeforeReserve() func() error { _ = "STUB: not implemented"; return nil }

func (e *Echo) getBeforeRead() func() error { _ = "STUB: not implemented"; return nil }

func (e *Echo) getBeforeWrite() func() error { _ = "STUB: not implemented"; return nil }

func (e *Echo) getBeforeDone() func() error { _ = "STUB: not implemented"; return nil }

func (e *Echo) getDone() func() { _ = "STUB: not implemented"; return nil }

func (e *Echo) handleStream(s network.Stream) { _ = "STUB: not implemented"; return }

func (e *Echo) Echo(p peer.ID, what string) error { _ = "STUB: not implemented"; return nil }
