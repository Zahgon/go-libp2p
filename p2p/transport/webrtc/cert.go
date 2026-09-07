package libp2pwebrtc

import (
	"crypto"
	"crypto/ecdsa"
	"io"
	"time"

	ic "github.com/libp2p/go-libp2p/core/crypto"

	"github.com/pion/webrtc/v4"
)

const deterministicCertHKDFInfo = "libp2p webrtc-direct deterministic cert"

var (
	deterministicCertNotBefore = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	deterministicCertNotAfter  = time.Date(2120, 1, 1, 0, 0, 0, 0, time.UTC)
)

func newDeterministicCertificate(key ic.PrivKey) (*webrtc.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type deterministicSigner struct {
	priv *ecdsa.PrivateKey
}

var _ crypto.Signer = deterministicSigner{}

func (ds deterministicSigner) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (ds deterministicSigner) Sign(_ io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
