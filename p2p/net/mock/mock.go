package mocknet

import (
	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("mocknet")

func WithNPeers(n int) (Mocknet, error) { _ = "STUB: not implemented"; return *new(Mocknet), nil }

func FullMeshLinked(n int) (Mocknet, error) { _ = "STUB: not implemented"; return *new(Mocknet), nil }

func FullMeshConnected(n int) (Mocknet, error) {
	_ = "STUB: not implemented"
	return *new(Mocknet), nil
}
