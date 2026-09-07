package libp2ptls

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"time"

	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

const certValidityPeriod = 100 * 365 * 24 * time.Hour
const certificatePrefix = "libp2p-tls-handshake:"
const alpn string = "libp2p"

var extensionID = getPrefixedExtensionID([]int{1, 1})
var extensionCritical bool

type signedKey struct {
	PubKey    []byte
	Signature []byte
}

type Identity struct {
	config tls.Config
}

type IdentityConfig struct {
	CertTemplate *x509.Certificate
	KeyLogWriter io.Writer
}

type IdentityOption func(r *IdentityConfig)

func WithCertTemplate(template *x509.Certificate) IdentityOption {
	_ = "STUB: not implemented"
	return *new(IdentityOption)
}

func WithKeyLogWriter(w io.Writer) IdentityOption {
	_ = "STUB: not implemented"
	return *new(IdentityOption)
}

func NewIdentity(privKey ic.PrivKey, opts ...IdentityOption) (*Identity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Identity) ConfigForPeer(remote peer.ID) (*tls.Config, <-chan ic.PubKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PubKeyFromCertChain(chain []*x509.Certificate) (ic.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(ic.PubKey), nil
}

func GenerateSignedExtension(sk ic.PrivKey, pubKey crypto.PublicKey) (pkix.Extension, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Extension), nil
}

func keyToCertificate(sk ic.PrivKey, certTmpl *x509.Certificate) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certTemplate() (*x509.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }
