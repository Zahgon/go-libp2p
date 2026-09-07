package test

import (
	"testing"

	"github.com/libp2p/go-libp2p/core/peer"
)

func RandPeerID() (peer.ID, error) { _ = "STUB: not implemented"; return *new(peer.ID), nil }

func RandPeerIDFatal(t testing.TB) peer.ID { _ = "STUB: not implemented"; return *new(peer.ID) }
