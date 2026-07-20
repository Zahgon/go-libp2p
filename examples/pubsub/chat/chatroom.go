package main

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const ChatRoomBufSize = 128

type ChatRoom struct {
	Messages chan *ChatMessage

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
	nick     string
}

type ChatMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

func JoinChatRoom(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, nickname string, roomName string) (*ChatRoom, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cr *ChatRoom) Publish(message string) error { _ = "STUB: not implemented"; return nil }

func (cr *ChatRoom) ListPeers() []peer.ID { _ = "STUB: not implemented"; return nil }

func (cr *ChatRoom) readLoop() { _ = "STUB: not implemented"; return }

func topicName(roomName string) string { _ = "STUB: not implemented"; return "" }
