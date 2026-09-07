package libp2pquic

import (
	"sync"

	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
)

const acceptBufferPerVersion = 4

type virtualListener struct {
	*listener
	udpAddr       string
	version       quic.Version
	t             *transport
	acceptRunnner *acceptLoopRunner
	acceptChan    chan acceptVal
}

var _ tpt.Listener = &virtualListener{}

func (l *virtualListener) Multiaddr() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

func (l *virtualListener) Close() error { _ = "STUB: not implemented"; return nil }

func (l *virtualListener) Accept() (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

type acceptVal struct {
	conn tpt.CapableConn
	err  error
}

type acceptLoopRunner struct {
	acceptSem chan struct{}

	muxerMu     sync.Mutex
	muxer       map[quic.Version]chan acceptVal
	muxerClosed bool
}

func (r *acceptLoopRunner) AcceptForVersion(v quic.Version) chan acceptVal {
	_ = "STUB: not implemented"
	return nil
}

func (r *acceptLoopRunner) RmAcceptForVersion(v quic.Version, err error) {
	_ = "STUB: not implemented"
	return
}

func (r *acceptLoopRunner) sendErrAndClose(err error) { _ = "STUB: not implemented"; return }

func (r *acceptLoopRunner) innerAccept(l *listener, expectedVersion quic.Version, bufferedConnChan chan acceptVal) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}

func (r *acceptLoopRunner) Accept(l *listener, expectedVersion quic.Version, bufferedConnChan chan acceptVal) (tpt.CapableConn, error) {
	_ = "STUB: not implemented"
	return *new(tpt.CapableConn), nil
}
