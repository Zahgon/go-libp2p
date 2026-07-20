package tlsdiag

import (
	ic "github.com/libp2p/go-libp2p/core/crypto"
)

func generateKey(keyType string) (priv ic.PrivKey, err error) {
	_ = "STUB: not implemented"
	return *new(ic.PrivKey), nil
}
