package blankhost

import (
	"context"

	"github.com/libp2p/go-libp2p/core/connmgr"
	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"

	logging "github.com/libp2p/go-libp2p/gologshim"

	ma "github.com/multiformats/go-multiaddr"
	mstream "github.com/multiformats/go-multistream"
)

var log = logging.Logger("blankhost")

type BlankHost struct {
	n        network.Network
	mux      *mstream.MultistreamMuxer[protocol.ID]
	cmgr     connmgr.ConnManager
	eventbus event.Bus
	emitters struct {
		evtLocalProtocolsUpdated event.Emitter
	}
}

type config struct {
	cmgr     connmgr.ConnManager
	eventBus event.Bus
}

type Option = func(cfg *config)

func WithConnectionManager(cmgr connmgr.ConnManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEventBus(eventBus event.Bus) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewBlankHost(n network.Network, options ...Option) *BlankHost {
	_ = "STUB: not implemented"
	return nil
}

func (bh *BlankHost) initSignedRecord() error { _ = "STUB: not implemented"; return nil }

var _ host.Host = (*BlankHost)(nil)

func (bh *BlankHost) Addrs() []ma.Multiaddr { _ = "STUB: not implemented"; return nil }

func (bh *BlankHost) Close() error { _ = "STUB: not implemented"; return nil }

func (bh *BlankHost) Connect(ctx context.Context, ai peer.AddrInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (bh *BlankHost) Peerstore() peerstore.Peerstore {
	_ = "STUB: not implemented"
	return *new(peerstore.Peerstore)
}

func (bh *BlankHost) ID() peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }

func (bh *BlankHost) NewStream(ctx context.Context, p peer.ID, protos ...protocol.ID) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (bh *BlankHost) RemoveStreamHandler(pid protocol.ID) { _ = "STUB: not implemented"; return }

func (bh *BlankHost) SetStreamHandler(pid protocol.ID, handler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func (bh *BlankHost) SetStreamHandlerMatch(pid protocol.ID, m func(protocol.ID) bool, handler network.StreamHandler) {
	_ = "STUB: not implemented"
	return
}

func (bh *BlankHost) newStreamHandler(s network.Stream) { _ = "STUB: not implemented"; return }

func (bh *BlankHost) Mux() protocol.Switch { _ = "STUB: not implemented"; return *new(protocol.Switch) }

func (bh *BlankHost) Network() network.Network {
	_ = "STUB: not implemented"
	return *new(network.Network)
}

func (bh *BlankHost) ConnManager() connmgr.ConnManager {
	_ = "STUB: not implemented"
	return *new(connmgr.ConnManager)
}

func (bh *BlankHost) EventBus() event.Bus { _ = "STUB: not implemented"; return *new(event.Bus) }
