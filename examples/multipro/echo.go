package main

import (
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"

	pb "github.com/libp2p/go-libp2p/examples/multipro/pb"
)

const echoRequest = "/echo/echoreq/0.0.1"
const echoResponse = "/echo/echoresp/0.0.1"

type EchoProtocol struct {
	node     *Node
	requests map[string]*pb.EchoRequest
	done     chan bool
}

func NewEchoProtocol(node *Node, done chan bool) *EchoProtocol {
	_ = "STUB: not implemented"
	return nil
}

func (e *EchoProtocol) onEchoRequest(s network.Stream) { _ = "STUB: not implemented"; return }

func (e *EchoProtocol) onEchoResponse(s network.Stream) { _ = "STUB: not implemented"; return }

func (e *EchoProtocol) Echo(host host.Host) bool { _ = "STUB: not implemented"; return false }
