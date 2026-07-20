package config

import (
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/quic-go/quic-go"
)

const (
	statelessResetKeyInfo = "libp2p quic stateless reset key"
	tokenGeneratorKeyInfo = "libp2p quic token generator key"
)

func PrivKeyToStatelessResetKey(key crypto.PrivKey) (quic.StatelessResetKey, error) {
	_ = "STUB: not implemented"
	return *new(quic.StatelessResetKey), nil
}

func PrivKeyToTokenGeneratorKey(key crypto.PrivKey) (quic.TokenGeneratorKey, error) {
	_ = "STUB: not implemented"
	return *new(quic.TokenGeneratorKey), nil
}
