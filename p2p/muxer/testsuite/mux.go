package mux

import (
	crand "crypto/rand"
	"net"
	"sync"
	"testing"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

var randomness []byte
var Subtests map[string]TransportTest

func init() {

	randomness = make([]byte, 1<<20)
	if _, err := crand.Read(randomness); err != nil {
		panic(err)
	}

	Subtests = make(map[string]TransportTest)
	for _, f := range subtests {
		Subtests[getFunctionName(f)] = f
	}
}

func getFunctionName(i any) string { _ = "STUB: not implemented"; return "" }

type peerScope struct {
	mx     sync.Mutex
	memory int
}

func (p *peerScope) ReserveMemory(size int, _ uint8) error { _ = "STUB: not implemented"; return nil }

func (p *peerScope) ReleaseMemory(size int) { _ = "STUB: not implemented"; return }

func (p *peerScope) Check(t *testing.T) { _ = "STUB: not implemented"; return }

type peerScopeSpan struct {
	peerScope
}

func (p *peerScopeSpan) Done() { _ = "STUB: not implemented"; return }

func (p *peerScope) Stat() network.ScopeStat {
	_ = "STUB: not implemented"
	return *new(network.ScopeStat)
}
func (p *peerScope) BeginSpan() (network.ResourceScopeSpan, error) {
	_ = "STUB: not implemented"
	return *new(network.ResourceScopeSpan), nil
}
func (p *peerScope) Peer() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

var _ network.PeerScope = &peerScope{}

type Options struct {
	tr        network.Multiplexer
	connNum   int
	streamNum int
	msgNum    int
	msgMin    int
	msgMax    int
}

func randBuf(size int) []byte { _ = "STUB: not implemented"; return nil }

func checkErr(t *testing.T, err error) { _ = "STUB: not implemented"; return }

func echoStream(s network.MuxedStream) { _ = "STUB: not implemented"; return }

func GoServe(t *testing.T, tr network.Multiplexer, l net.Listener) (done func()) {
	_ = "STUB: not implemented"
	return nil
}

func SubtestSimpleWrite(t *testing.T, tr network.Multiplexer) { _ = "STUB: not implemented"; return }

func SubtestStress(t *testing.T, opt Options) { _ = "STUB: not implemented"; return }

func tcpPipe(t *testing.T) (net.Conn, net.Conn) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(net.Conn)
}

func SubtestStreamOpenStress(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStreamReset(t *testing.T, tr network.Multiplexer) { _ = "STUB: not implemented"; return }

func SubtestWriteAfterClose(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStreamLeftOpen(t *testing.T, tr network.Multiplexer) { _ = "STUB: not implemented"; return }

func SubtestStress1Conn1Stream1Msg(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn1Stream100Msg(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn100Stream100Msg(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress10Conn10Stream50Msg(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn1000Stream10Msg(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, tr network.Multiplexer) {
	_ = "STUB: not implemented"
	return
}

var subtests = []TransportTest{
	SubtestSimpleWrite,
	SubtestWriteAfterClose,
	SubtestStress1Conn1Stream1Msg,
	SubtestStress1Conn1Stream100Msg,
	SubtestStress1Conn100Stream100Msg,
	SubtestStress10Conn10Stream50Msg,
	SubtestStress1Conn1000Stream10Msg,
	SubtestStress1Conn100Stream100Msg10MB,
	SubtestStreamOpenStress,
	SubtestStreamReset,
	SubtestStreamLeftOpen,
}

func SubtestAll(t *testing.T, tr network.Multiplexer) { _ = "STUB: not implemented"; return }

type TransportTest func(t *testing.T, tr network.Multiplexer)
