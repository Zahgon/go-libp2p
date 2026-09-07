package util

import (
	"context"

	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/peer"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("discovery-util")

func FindPeers(ctx context.Context, d discovery.Discoverer, ns string, opts ...discovery.Option) ([]peer.AddrInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Advertise(ctx context.Context, a discovery.Advertiser, ns string, opts ...discovery.Option) {
	_ = "STUB: not implemented"
	return
}
