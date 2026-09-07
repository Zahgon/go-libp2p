package websocket

import (
	"net"
	"net/url"

	ma "github.com/multiformats/go-multiaddr"
)

type Addr struct {
	*url.URL
}

var _ net.Addr = (*Addr)(nil)

func (addr *Addr) Network() string { _ = "STUB: not implemented"; return "" }

func NewAddr(host string) *Addr { _ = "STUB: not implemented"; return nil }

func NewAddrWithScheme(host string, isSecure bool) *Addr { _ = "STUB: not implemented"; return nil }

func ConvertWebsocketMultiaddrToNetAddr(maddr ma.Multiaddr) (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func ParseWebsocketNetAddr(a net.Addr) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func parseMultiaddr(maddr ma.Multiaddr) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type parsedWebsocketMultiaddr struct {
	isWSS bool

	sni *ma.Component

	restMultiaddr ma.Multiaddr
}

func parseWebsocketMultiaddr(a ma.Multiaddr) (parsedWebsocketMultiaddr, error) {
	_ = "STUB: not implemented"
	return *new(parsedWebsocketMultiaddr), nil
}
