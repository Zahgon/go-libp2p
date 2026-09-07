package network

import (
	"github.com/libp2p/go-libp2p/core/protocol"
)

type Stream interface {
	MuxedStream

	ID() string

	Protocol() protocol.ID
	SetProtocol(id protocol.ID) error

	Stat() Stats

	Conn() Conn

	Scope() StreamScope

	ResetWithError(errCode StreamErrorCode) error
}
