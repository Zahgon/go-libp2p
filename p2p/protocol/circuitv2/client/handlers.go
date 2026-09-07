package client

import (
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

var (
	StreamTimeout = 1 * time.Minute
	AcceptTimeout = 10 * time.Second
)

func (c *Client) handleStreamV2(s network.Stream) { _ = "STUB: not implemented"; return }
