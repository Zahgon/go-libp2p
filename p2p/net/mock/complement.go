package mocknet

import (
	"github.com/libp2p/go-libp2p/core/network"
)

func StreamComplement(s network.Stream) network.Stream {
	_ = "STUB: not implemented"
	return *new(network.Stream)
}

func ConnComplement(c network.Conn) network.Conn {
	_ = "STUB: not implemented"
	return *new(network.Conn)
}
