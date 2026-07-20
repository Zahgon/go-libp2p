package main

import (
	"github.com/libp2p/go-libp2p"
)

func main() {

	transports := libp2p.ChainOptions()

	host, err := libp2p.New(transports)
	if err != nil {
		panic(err)
	}

	host.Close()
}
