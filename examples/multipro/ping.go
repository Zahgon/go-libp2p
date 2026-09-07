package main

import (
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"

	p2p "github.com/libp2p/go-libp2p/examples/multipro/pb"
)

const pingRequest = "/ping/pingreq/0.0.1"
const pingResponse = "/ping/pingresp/0.0.1"

type PingProtocol struct {
	node     *Node
	mu       sync.Mutex
	requests map[string]*p2p.PingRequest
	done     chan bool
}

func NewPingProtocol(node *Node, done chan bool) *PingProtocol {
	_ = "STUB: not implemented"
	return nil
}

func (p *PingProtocol) onPingRequest(s network.Stream) { _ = "STUB: not implemented"; return }

func (p *PingProtocol) onPingResponse(s network.Stream) { _ = "STUB: not implemented"; return }

func (p *PingProtocol) Ping(host host.Host) bool { _ = "STUB: not implemented"; return false }
