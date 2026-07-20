package metrics

import (
	"time"

	"github.com/libp2p/go-flow-metrics"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type BandwidthCounter struct {
	totalIn  flow.Meter
	totalOut flow.Meter

	protocolIn  flow.MeterRegistry
	protocolOut flow.MeterRegistry

	peerIn  flow.MeterRegistry
	peerOut flow.MeterRegistry
}

func NewBandwidthCounter() *BandwidthCounter { _ = "STUB: not implemented"; return nil }

func (bwc *BandwidthCounter) LogSentMessage(size int64) { _ = "STUB: not implemented"; return }

func (bwc *BandwidthCounter) LogRecvMessage(size int64) { _ = "STUB: not implemented"; return }

func (bwc *BandwidthCounter) LogSentMessageStream(size int64, proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

func (bwc *BandwidthCounter) LogRecvMessageStream(size int64, proto protocol.ID, p peer.ID) {
	_ = "STUB: not implemented"
	return
}

func (bwc *BandwidthCounter) GetBandwidthForPeer(p peer.ID) (out Stats) {
	_ = "STUB: not implemented"
	return *new(Stats)
}

func (bwc *BandwidthCounter) GetBandwidthForProtocol(proto protocol.ID) (out Stats) {
	_ = "STUB: not implemented"
	return *new(Stats)
}

func (bwc *BandwidthCounter) GetBandwidthTotals() (out Stats) {
	_ = "STUB: not implemented"
	return *new(Stats)
}

func (bwc *BandwidthCounter) GetBandwidthByPeer() map[peer.ID]Stats {
	_ = "STUB: not implemented"
	return nil
}

func (bwc *BandwidthCounter) GetBandwidthByProtocol() map[protocol.ID]Stats {
	_ = "STUB: not implemented"
	return nil
}

func (bwc *BandwidthCounter) Reset() { _ = "STUB: not implemented"; return }

func (bwc *BandwidthCounter) TrimIdle(since time.Time) { _ = "STUB: not implemented"; return }
