package peerstore

import (
	"context"
	"errors"
	"io"
	"math"
	"time"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/record"

	ma "github.com/multiformats/go-multiaddr"
)

var ErrNotFound = errors.New("item not found")

var (
	AddressTTL = time.Hour

	TempAddrTTL = time.Minute * 2

	RecentlyConnectedAddrTTL = time.Minute * 15

	OwnObservedAddrTTL = time.Minute * 30
)

const (
	PermanentAddrTTL = math.MaxInt64 - iota

	ConnectedAddrTTL
)

type Peerstore interface {
	io.Closer

	AddrBook
	KeyBook
	PeerMetadata
	Metrics
	ProtoBook

	PeerInfo(peer.ID) peer.AddrInfo

	Peers() peer.IDSlice

	RemovePeer(peer.ID)
}

type PeerMetadata interface {
	Get(p peer.ID, key string) (any, error)
	Put(p peer.ID, key string, val any) error

	RemovePeer(peer.ID)
}

type AddrBook interface {
	AddAddr(p peer.ID, addr ma.Multiaddr, ttl time.Duration)

	AddAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration)

	SetAddr(p peer.ID, addr ma.Multiaddr, ttl time.Duration)

	SetAddrs(p peer.ID, addrs []ma.Multiaddr, ttl time.Duration)

	UpdateAddrs(p peer.ID, oldTTL time.Duration, newTTL time.Duration)

	Addrs(p peer.ID) []ma.Multiaddr

	AddrStream(context.Context, peer.ID) <-chan ma.Multiaddr

	ClearAddrs(p peer.ID)

	PeersWithAddrs() peer.IDSlice
}

type CertifiedAddrBook interface {
	ConsumePeerRecord(s *record.Envelope, ttl time.Duration) (accepted bool, err error)

	GetPeerRecord(p peer.ID) *record.Envelope
}

func GetCertifiedAddrBook(ab AddrBook) (cab CertifiedAddrBook, ok bool) {
	_ = "STUB: not implemented"
	return *new(CertifiedAddrBook), false
}

type KeyBook interface {
	PubKey(peer.ID) ic.PubKey

	AddPubKey(peer.ID, ic.PubKey) error

	PrivKey(peer.ID) ic.PrivKey

	AddPrivKey(peer.ID, ic.PrivKey) error

	PeersWithKeys() peer.IDSlice

	RemovePeer(peer.ID)
}

type Metrics interface {
	RecordLatency(peer.ID, time.Duration)

	LatencyEWMA(peer.ID) time.Duration

	RemovePeer(peer.ID)
}

type ProtoBook interface {
	GetProtocols(peer.ID) ([]protocol.ID, error)
	AddProtocols(peer.ID, ...protocol.ID) error
	SetProtocols(peer.ID, ...protocol.ID) error
	RemoveProtocols(peer.ID, ...protocol.ID) error

	SupportsProtocols(peer.ID, ...protocol.ID) ([]protocol.ID, error)

	FirstSupportedProtocol(peer.ID, ...protocol.ID) (protocol.ID, error)

	RemovePeer(peer.ID)
}
