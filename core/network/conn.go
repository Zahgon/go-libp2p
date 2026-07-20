package network

import (
	"context"
	"io"

	ic "github.com/libp2p/go-libp2p/core/crypto"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	ma "github.com/multiformats/go-multiaddr"
)

type ConnErrorCode uint32

type ConnError struct {
	Remote         bool
	ErrorCode      ConnErrorCode
	TransportError error
}

func (c *ConnError) Error() string { _ = "STUB: not implemented"; return "" }

func (c *ConnError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (c *ConnError) Unwrap() []error { _ = "STUB: not implemented"; return nil }

const (
	ConnNoError                   ConnErrorCode = 0
	ConnProtocolNegotiationFailed ConnErrorCode = 0x1000
	ConnResourceLimitExceeded     ConnErrorCode = 0x1001
	ConnRateLimited               ConnErrorCode = 0x1002
	ConnProtocolViolation         ConnErrorCode = 0x1003
	ConnSupplanted                ConnErrorCode = 0x1004
	ConnGarbageCollected          ConnErrorCode = 0x1005
	ConnShutdown                  ConnErrorCode = 0x1006
	ConnGated                     ConnErrorCode = 0x1007
	ConnCodeOutOfRange            ConnErrorCode = 0x1008
)

type Conn interface {
	io.Closer

	ConnSecurity
	ConnMultiaddrs
	ConnStat
	ConnScoper

	CloseWithError(errCode ConnErrorCode) error

	ID() string

	NewStream(context.Context) (Stream, error)

	GetStreams() []Stream

	IsClosed() bool

	As(target any) bool
}

type ConnectionState struct {
	StreamMultiplexer protocol.ID

	Security protocol.ID

	Transport string

	UsedEarlyMuxerNegotiation bool
}

type ConnSecurity interface {
	LocalPeer() peer.ID

	RemotePeer() peer.ID

	RemotePublicKey() ic.PubKey

	ConnState() ConnectionState
}

type ConnMultiaddrs interface {
	LocalMultiaddr() ma.Multiaddr

	RemoteMultiaddr() ma.Multiaddr
}

type ConnStat interface {
	Stat() ConnStats
}

type ConnScoper interface {
	Scope() ConnScope
}
