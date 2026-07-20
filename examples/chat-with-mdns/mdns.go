package main

import (
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) { _ = "STUB: not implemented"; return }

func initMDNS(peerhost host.Host, rendezvous string) chan peer.AddrInfo {
	_ = "STUB: not implemented"
	return nil
}
