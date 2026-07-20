package libp2pwebrtc

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/p2p/transport/webrtc/pb"
	"github.com/libp2p/go-msgio/pbio"

	"github.com/pion/datachannel"
	"github.com/pion/webrtc/v4"
)

const (
	maxSendMessageSize = 16384

	protoOverhead = 5

	varintOverhead = 2

	maxTotalControlMessagesSize = 50

	maxFINACKWait = 10 * time.Second

	maxReceiveMessageSize = 256<<10 + 1<<10
)

type receiveState uint8

const (
	receiveStateReceiving receiveState = iota
	receiveStateDataRead
	receiveStateReset
)

type sendState uint8

const (
	sendStateSending sendState = iota
	sendStateDataSent
	sendStateDataReceived
	sendStateReset
)

type stream struct {
	mx sync.Mutex

	readerMx  sync.Mutex
	reader    pbio.Reader
	readError error

	nextMessage  *pb.Message
	receiveState receiveState

	writer             pbio.Writer
	writeStateChanged  chan struct{}
	sendState          sendState
	writeDeadline      time.Time
	writeError         error
	maxSendMessageSize int

	controlMessageReaderOnce sync.Once

	controlMessageReaderEndTime time.Time

	onDoneOnce          sync.Once
	onDone              func()
	id                  uint16
	dataChannel         *datachannel.DataChannel
	closeForShutdownErr error
}

var _ network.MuxedStream = &stream{}

func newStream(
	channel *webrtc.DataChannel,
	rwc datachannel.ReadWriteCloser,
	maxSendMessageSize int,
	onDone func(),
) *stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *stream) ResetWithError(errCode network.StreamErrorCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) closeForShutdown(closeErr error) { _ = "STUB: not implemented"; return }

func (s *stream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) processIncomingFlag(msg *pb.Message) { _ = "STUB: not implemented"; return }

func (s *stream) spawnControlMessageReader() { _ = "STUB: not implemented"; return }

func (s *stream) cleanup() { _ = "STUB: not implemented"; return }
