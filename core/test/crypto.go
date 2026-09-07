package test

import (
	"sync/atomic"

	ci "github.com/libp2p/go-libp2p/core/crypto"
)

var globalSeed atomic.Int64

func RandTestKeyPair(typ, bits int) (ci.PrivKey, ci.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(ci.PrivKey), *new(ci.PubKey), nil
}

func SeededTestKeyPair(typ, bits int, seed int64) (ci.PrivKey, ci.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(ci.PrivKey), *new(ci.PubKey), nil
}
