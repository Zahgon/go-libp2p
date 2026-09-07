package crypto

import (
	"errors"
	"io"

	"github.com/libp2p/go-libp2p/core/crypto/pb"
)

const (
	RSA = iota

	Ed25519

	Secp256k1

	ECDSA
)

var (
	ErrBadKeyType = errors.New("invalid or unsupported key type")

	KeyTypes = []int{
		RSA,
		Ed25519,
		Secp256k1,
		ECDSA,
	}
)

type PubKeyUnmarshaller func(data []byte) (PubKey, error)

type PrivKeyUnmarshaller func(data []byte) (PrivKey, error)

var PubKeyUnmarshallers = map[pb.KeyType]PubKeyUnmarshaller{
	pb.KeyType_RSA:       UnmarshalRsaPublicKey,
	pb.KeyType_Ed25519:   UnmarshalEd25519PublicKey,
	pb.KeyType_Secp256k1: UnmarshalSecp256k1PublicKey,
	pb.KeyType_ECDSA:     UnmarshalECDSAPublicKey,
}

var PrivKeyUnmarshallers = map[pb.KeyType]PrivKeyUnmarshaller{
	pb.KeyType_RSA:       UnmarshalRsaPrivateKey,
	pb.KeyType_Ed25519:   UnmarshalEd25519PrivateKey,
	pb.KeyType_Secp256k1: UnmarshalSecp256k1PrivateKey,
	pb.KeyType_ECDSA:     UnmarshalECDSAPrivateKey,
}

type Key interface {
	Equals(Key) bool

	Raw() ([]byte, error)

	Type() pb.KeyType
}

type PrivKey interface {
	Key

	Sign([]byte) ([]byte, error)

	GetPublic() PubKey
}

type PubKey interface {
	Key

	Verify(data []byte, sig []byte) (bool, error)
}

type GenSharedKey func([]byte) ([]byte, error)

func GenerateKeyPair(typ, bits int) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func GenerateKeyPairWithReader(typ, bits int, src io.Reader) (PrivKey, PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), *new(PubKey), nil
}

func UnmarshalPublicKey(data []byte) (PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}

func PublicKeyFromProto(pmes *pb.PublicKey) (PubKey, error) {
	_ = "STUB: not implemented"
	return *new(PubKey), nil
}

func MarshalPublicKey(k PubKey) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func PublicKeyToProto(k PubKey) (*pb.PublicKey, error) { _ = "STUB: not implemented"; return nil, nil }

func UnmarshalPrivateKey(data []byte) (PrivKey, error) {
	_ = "STUB: not implemented"
	return *new(PrivKey), nil
}

func MarshalPrivateKey(k PrivKey) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ConfigDecodeKey(b string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ConfigEncodeKey(b []byte) string { _ = "STUB: not implemented"; return "" }

func KeyEqual(k1, k2 Key) bool { _ = "STUB: not implemented"; return false }

func basicEquals(k1, k2 Key) bool { _ = "STUB: not implemented"; return false }
