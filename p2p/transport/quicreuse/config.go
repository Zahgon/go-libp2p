package quicreuse

import (
	"time"

	"github.com/quic-go/quic-go"
)

var quicConfig = &quic.Config{
	MaxIncomingStreams:         256,
	MaxIncomingUniStreams:      5,
	MaxStreamReceiveWindow:     10 * (1 << 20),
	MaxConnectionReceiveWindow: 15 * (1 << 20),
	KeepAlivePeriod:            15 * time.Second,
	Versions:                   []quic.Version{quic.Version1},

	EnableDatagrams: true,

	EnableStreamResetPartialDelivery: true,
}
