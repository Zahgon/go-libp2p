package mocknet

import (
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

type printer struct {
	w io.Writer
}

func (p *printer) MocknetLinks(mn Mocknet) { _ = "STUB: not implemented"; return }

func (p *printer) NetworkConns(ni network.Network) { _ = "STUB: not implemented"; return }
