package main

import (
	"context"
	"flag"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
)

var (
	topicNameFlag = flag.String("topicName", "applesauce", "name of topic to join")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	go discoverPeers(ctx, h)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	topic, err := ps.Join(*topicNameFlag)
	if err != nil {
		panic(err)
	}
	go streamConsoleTo(ctx, topic)

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	printMessagesFrom(ctx, sub)
}

func initDHT(ctx context.Context, h host.Host) *dht.IpfsDHT { _ = "STUB: not implemented"; return nil }

func discoverPeers(ctx context.Context, h host.Host) { _ = "STUB: not implemented"; return }

func streamConsoleTo(ctx context.Context, topic *pubsub.Topic) { _ = "STUB: not implemented"; return }

func printMessagesFrom(ctx context.Context, sub *pubsub.Subscription) {
	_ = "STUB: not implemented"
	return
}
