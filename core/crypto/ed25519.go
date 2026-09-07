package crypto

import (
	"crypto/ed25519"
	"io"

	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
)

type Ed25519PrivateKey struct {
	k ed25519.PrivateKey
}

type Ed25519PublicKey struct {
	k ed25519.PublicKey
}

func GenerateEd25519Key(src io.Reader) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func (k *Ed25519PrivateKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (k *Ed25519PrivateKey) Raw() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *Ed25519PrivateKey) pubKeyBytes() []byte { _ = "STUB: not implemented"; return nil }

func (k *Ed25519PrivateKey) Equals(o Key) bool { _ = "STUB: not implemented"; return false }

func (k *Ed25519PrivateKey) GetPublic() PubKey { _ = "STUB: not implemented"; return *new(PubKey) }

func (k *Ed25519PrivateKey) Sign(msg []byte) (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Ed25519PublicKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (k *Ed25519PublicKey) Raw() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *Ed25519PublicKey) Equals(o Key) bool { _ = "STUB: not implemented"; return false }

func (k *Ed25519PublicKey) Verify(data []byte, sig []byte) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func UnmarshalEd25519PublicKey(data []byte) (PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}

func UnmarshalEd25519PrivateKey(data []byte) (PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), nil
}
