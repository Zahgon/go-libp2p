package main

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
)

func main() {
	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(websocket.New),
	)

	host, err := libp2p.New(transports)
	if err != nil {
		panic(err)
	}

	host.Close()
}
