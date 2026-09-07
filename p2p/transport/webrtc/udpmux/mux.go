package udpmux

import (
	"context"
	"net"
	"sync"

	logging "github.com/libp2p/go-libp2p/gologshim"
	"github.com/pion/ice/v4"
	"github.com/pion/stun/v3"
)

var log = logging.Logger("webrtc-udpmux")

const ReceiveBufSize = 1500

const maxAddrsPerUfrag = 32

type Candidate struct {
	LocalUfrag string

	RemoteUfrag string

	RemotePwd string
	Addr      *net.UDPAddr
}

type UDPMux struct {
	socket net.PacketConn

	queue chan Candidate

	mx sync.Mutex

	ufragMap map[ufragConnKey]*muxedConnection

	addrMap map[string]*muxedConnection

	ufragAddrMap map[ufragConnKey][]net.Addr

	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

var _ ice.UDPMux = &UDPMux{}

func NewUDPMux(socket net.PacketConn) *UDPMux { _ = "STUB: not implemented"; return nil }

func (mux *UDPMux) Start() { _ = "STUB: not implemented"; return }

func (mux *UDPMux) GetListenAddresses() []net.Addr { _ = "STUB: not implemented"; return nil }

func (mux *UDPMux) GetConn(localUfrag string, addr net.Addr) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

func (mux *UDPMux) Close() error { _ = "STUB: not implemented"; return nil }

func (mux *UDPMux) writeTo(buf []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mux *UDPMux) readLoop() { _ = "STUB: not implemented"; return }

func (mux *UDPMux) processPacket(buf []byte, addr net.Addr) (processed bool) {
	_ = "STUB: not implemented"
	return false
}

func (mux *UDPMux) Accept(ctx context.Context) (Candidate, error) {
	_ = "STUB: not implemented"
	return *new(Candidate), nil
}

type ufragConnKey struct {
	localUfrag string
	isIPv6     bool
}

func credentialsFromSTUNMessage(msg *stun.Message) (localUfrag, remoteUfrag, remotePwd string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func (mux *UDPMux) RemoveConnByUfrag(localUfrag string) { _ = "STUB: not implemented"; return }

func (mux *UDPMux) getOrCreateConn(localUfrag string, isIPv6 bool, _ *UDPMux, addr net.Addr) (created bool, _ *muxedConnection) {
	_ = "STUB: not implemented"
	return false, nil
}
