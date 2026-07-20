package ttransport

import (
	"os"
	"strconv"
	"testing"
	"time"

	crand "crypto/rand"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

var randomness []byte

var StressTestTimeout = 1 * time.Minute

func init() {

	randomness = make([]byte, 1<<20)
	if _, err := crand.Read(randomness); err != nil {
		panic(err)
	}

	if timeout := os.Getenv("TEST_STRESS_TIMEOUT_MS"); timeout != "" {
		if v, err := strconv.ParseInt(timeout, 10, 32); err == nil {
			StressTestTimeout = time.Duration(v) * time.Millisecond
		}
	}
}

type Options struct {
	ConnNum   int
	StreamNum int
	MsgNum    int
	MsgMin    int
	MsgMax    int
}

func fullClose(t *testing.T, s network.MuxedStream) { _ = "STUB: not implemented"; return }

func randBuf(size int) []byte { _ = "STUB: not implemented"; return nil }

func echoStream(t *testing.T, s network.MuxedStream) { _ = "STUB: not implemented"; return }

func echo(t *testing.T, c transport.CapableConn) { _ = "STUB: not implemented"; return }

func serve(t *testing.T, l transport.Listener) { _ = "STUB: not implemented"; return }

func SubtestStress(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID, opt Options) {
	_ = "STUB: not implemented"
	return
}

func SubtestStreamOpenStress(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStreamReset(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn1Stream1Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn1Stream100Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn100Stream100Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStressManyConn10Stream50Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn1000Stream10Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}

func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	_ = "STUB: not implemented"
	return
}
