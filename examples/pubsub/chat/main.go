package main

import (
	"context"
	"flag"
	"time"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

const DiscoveryInterval = time.Hour

const DiscoveryServiceTag = "pubsub-chat-example"

func main() {

	nickFlag := flag.String("nick", "", "nickname to use in chat. will be generated if empty")
	roomFlag := flag.String("room", "awesome-chat-room", "name of chat room to join")
	flag.Parse()

	ctx := context.Background()

	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	if err := setupDiscovery(h); err != nil {
		panic(err)
	}

	nick := *nickFlag
	if len(nick) == 0 {
		nick = defaultNick(h.ID())
	}

	room := *roomFlag

	cr, err := JoinChatRoom(ctx, ps, h.ID(), nick, room)
	if err != nil {
		panic(err)
	}

	ui := NewChatUI(cr)
	if err = ui.Run(); err != nil {
		printErr("error running text UI: %s", err)
	}
}

func printErr(m string, args ...interface{}) { _ = "STUB: not implemented"; return }

func defaultNick(p peer.ID) string { _ = "STUB: not implemented"; return "" }

func shortID(p peer.ID) string { _ = "STUB: not implemented"; return "" }

type discoveryNotifee struct {
	h host.Host
}

func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) { _ = "STUB: not implemented"; return }

func setupDiscovery(h host.Host) error { _ = "STUB: not implemented"; return nil }
