package main

import (
	"github.com/libp2p/go-libp2p/core/crypto"
)

func LoadIdentity(keyPath string) (crypto.PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey), nil
}

func ReadIdentity(path string) (crypto.PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey), nil
}

func GenerateIdentity(path string) (crypto.PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivKey), nil
}
