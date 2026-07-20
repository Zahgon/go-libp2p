package crypto

import (
	"io"

	pb "github.com/libp2p/go-libp2p/core/crypto/pb"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type Secp256k1PrivateKey secp256k1.PrivateKey

type Secp256k1PublicKey secp256k1.PublicKey

func GenerateSecp256k1Key(_ io.Reader) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func UnmarshalSecp256k1PrivateKey(data []byte) (k PrivKey, err error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), nil
}

func UnmarshalSecp256k1PublicKey(data []byte) (_k PubKey, err error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}

func (k *Secp256k1PrivateKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (k *Secp256k1PrivateKey) Raw() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *Secp256k1PrivateKey) Equals(o Key) bool { _ = "STUB: not implemented"; return false }

func (k *Secp256k1PrivateKey) Sign(data []byte) (_sig []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Secp256k1PrivateKey) GetPublic() PubKey { _ = "STUB: not implemented"; return *new(PubKey) }

func (k *Secp256k1PublicKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (k *Secp256k1PublicKey) Raw() (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Secp256k1PublicKey) Equals(o Key) bool { _ = "STUB: not implemented"; return false }

func (k *Secp256k1PublicKey) Verify(data []byte, sigStr []byte) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
