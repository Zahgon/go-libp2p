package crypto

import (
	"crypto/rsa"
	"io"

	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
)

type RsaPrivateKey struct {
	sk rsa.PrivateKey
}

type RsaPublicKey struct {
	k rsa.PublicKey

	cached []byte
}

func GenerateRSAKeyPair(bits int, src io.Reader) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func (pk *RsaPublicKey) Verify(data, sig []byte) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pk *RsaPublicKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (pk *RsaPublicKey) Raw() (res []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (pk *RsaPublicKey) Equals(k Key) bool { _ = "STUB: not implemented"; return false }

func (sk *RsaPrivateKey) Sign(message []byte) (sig []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sk *RsaPrivateKey) GetPublic() PubKey { _ = "STUB: not implemented"; return *new(PubKey) }

func (sk *RsaPrivateKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (sk *RsaPrivateKey) Raw() (res []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (sk *RsaPrivateKey) Equals(k Key) bool { _ = "STUB: not implemented"; return false }

func UnmarshalRsaPrivateKey(b []byte) (key PrivKey, err error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), nil
}

func UnmarshalRsaPublicKey(b []byte) (key PubKey, err error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}
