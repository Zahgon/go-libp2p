package autonat

import (
	"context"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

type AutoNAT interface {
	Status() network.Reachability
	io.Closer
}

type Client interface {
	DialBack(ctx context.Context, p peer.ID) error
}

type AddrFunc func() []ma.Multiaddr

type Option func(*config) error
