package autonat

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/host/autonat/pb"

	ma "github.com/multiformats/go-multiaddr"
)

const AutoNATProto = "/libp2p/autonat/1.0.0"

func newDialMessage(pi peer.AddrInfo) *pb.Message { _ = "STUB: not implemented"; return nil }

func newDialResponseOK(addr ma.Multiaddr) *pb.Message_DialResponse {
	_ = "STUB: not implemented"
	return nil
}

func newDialResponseError(status pb.Message_ResponseStatus, text string) *pb.Message_DialResponse {
	_ = "STUB: not implemented"
	return nil
}
