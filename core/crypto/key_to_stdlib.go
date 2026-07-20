package crypto

import (
	"crypto"
)

func KeyPairFromStdKey(priv crypto.PrivateKey) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func PrivKeyToStdKey(priv PrivKey) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func PubKeyToStdKey(pub PubKey) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}
