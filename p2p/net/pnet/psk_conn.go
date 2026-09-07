package pnet

import (
	"crypto/cipher"
	"net"

	"github.com/libp2p/go-libp2p/core/pnet"
)

var (
	errShortNonce  = pnet.NewError("could not read full nonce")
	errInsecureNil = pnet.NewError("insecure is nil")
	errPSKNil      = pnet.NewError("pre-shread key is nil")
)

type pskConn struct {
	net.Conn
	psk *[32]byte

	writeS20 cipher.Stream
	readS20  cipher.Stream
}

func (c *pskConn) Read(out []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *pskConn) Write(in []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var _ net.Conn = (*pskConn)(nil)

func newPSKConn(psk *[32]byte, insecure net.Conn) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
