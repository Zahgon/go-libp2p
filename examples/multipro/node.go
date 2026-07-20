package main

import (
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	p2p "github.com/libp2p/go-libp2p/examples/multipro/pb"

	"github.com/gogo/protobuf/proto"
)

const clientVersion = "go-p2p-node/0.0.1"

type Node struct {
	host.Host
	*PingProtocol
	*EchoProtocol
}

func NewNode(host host.Host, done chan bool) *Node { _ = "STUB: not implemented"; return nil }

func (n *Node) authenticateMessage(message proto.Message, data *p2p.MessageData) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *Node) signProtoMessage(message proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Node) signData(data []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *Node) verifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *Node) NewMessageData(messageId string, gossip bool) *p2p.MessageData {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) sendProtoMessage(id peer.ID, p protocol.ID, data proto.Message) bool {
	_ = "STUB: not implemented"
	return false
}
