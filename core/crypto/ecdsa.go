package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"io"
	"math/big"

	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
)

type ECDSAPrivateKey struct {
	priv *ecdsa.PrivateKey
}

type ECDSAPublicKey struct {
	pub *ecdsa.PublicKey
}

type ECDSASig struct {
	R, S *big.Int
}

var (
	ErrNotECDSAPubKey = errors.New("not an ecdsa public key")

	ErrNilSig = errors.New("sig is nil")

	ErrNilPrivateKey = errors.New("private key is nil")

	ErrNilPublicKey = errors.New("public key is nil")

	ECDSACurve = elliptic.P256()
)

func GenerateECDSAKeyPair(src io.Reader) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func GenerateECDSAKeyPairWithCurve(curve elliptic.Curve, src io.Reader) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func ECDSAKeyPairFromKey(priv *ecdsa.PrivateKey) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func ECDSAPublicKeyFromPubKey(pub ecdsa.PublicKey) (PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}

func MarshalECDSAPrivateKey(ePriv ECDSAPrivateKey) (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalECDSAPublicKey(ePub ECDSAPublicKey) (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalECDSAPrivateKey(data []byte) (res PrivKey, err error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), nil
}

func UnmarshalECDSAPublicKey(data []byte) (key PubKey, err error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}

func (ePriv *ECDSAPrivateKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (ePriv *ECDSAPrivateKey) Raw() (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ePriv *ECDSAPrivateKey) Equals(o Key) bool { _ = "STUB: not implemented"; return false }

func (ePriv *ECDSAPrivateKey) Sign(data []byte) (sig []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ePriv *ECDSAPrivateKey) GetPublic() PubKey { _ = "STUB: not implemented"; return *new(PubKey) }

func (ePub *ECDSAPublicKey) Type() pb.KeyType { _ = "STUB: not implemented"; return *new(pb.KeyType) }

func (ePub *ECDSAPublicKey) Raw() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ePub *ECDSAPublicKey) Equals(o Key) bool { _ = "STUB: not implemented"; return false }

func (ePub *ECDSAPublicKey) Verify(data, sigBytes []byte) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
