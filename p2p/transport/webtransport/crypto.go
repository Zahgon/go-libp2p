package libp2pwebtransport

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"io"
	"time"

	ic "github.com/libp2p/go-libp2p/core/crypto"

	"github.com/multiformats/go-multihash"
)

const deterministicCertInfo = "determinisitic cert"

func getTLSConf(key ic.PrivKey, start, end time.Time) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateCert(key ic.PrivKey, start, end time.Time) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ErrCertHashMismatch struct {
	Expected []byte
	Actual   [][]byte
}

func (e ErrCertHashMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func verifyRawCerts(rawCerts [][]byte, certHashes []multihash.DecodedMultihash) error {
	_ = "STUB: not implemented"
	return nil
}

func newDeterministicReader(seed []byte, salt []byte, info string) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

type deterministicSigner struct {
	priv *ecdsa.PrivateKey
}

var _ crypto.Signer = deterministicSigner{}

func (ds deterministicSigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ds deterministicSigner) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}
