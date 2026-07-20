package libp2pwebtransport

import (
	"net"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
)

var webtransportMA = ma.StringCast("/quic-v1/webtransport")

func toWebtransportMultiaddr(na net.Addr) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func stringToWebtransportMultiaddr(str string) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func extractCertHashes(addr ma.Multiaddr) ([]multihash.DecodedMultihash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addrComponentForCert(hash []byte) (*ma.Component, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IsWebtransportMultiaddr(multiaddr ma.Multiaddr) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}
