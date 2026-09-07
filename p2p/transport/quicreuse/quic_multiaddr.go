package quicreuse

import (
	"net"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/quic-go/quic-go"
)

var (
	quicV1MA = ma.StringCast("/quic-v1")
)

func ToQuicMultiaddr(na net.Addr, version quic.Version) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func FromQuicMultiaddr(addr ma.Multiaddr) (*net.UDPAddr, quic.Version, error) {
	_ = "STUB: not implemented"
	return nil, *new(quic.Version), nil
}
