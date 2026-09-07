package core

import (
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	"github.com/multiformats/go-multiaddr"
)

type Multiaddr = multiaddr.Multiaddr

type PeerID = peer.ID

type ProtocolID = protocol.ID

type PeerAddrInfo = peer.AddrInfo

type Host = host.Host

type Network = network.Network

type Conn = network.Conn

type Stream = network.Stream
