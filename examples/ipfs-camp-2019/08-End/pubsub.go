package main

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

var handles = map[string]string{}

const pubsubTopic = "/libp2p/example/chat/1.0.0"

func pubsubMessageHandler(id peer.ID, msg *SendMessage) { _ = "STUB: not implemented"; return }

func pubsubUpdateHandler(id peer.ID, msg *UpdatePeer) { _ = "STUB: not implemented"; return }

func pubsubHandler(ctx context.Context, sub *pubsub.Subscription) {
	_ = "STUB: not implemented"
	return
}
