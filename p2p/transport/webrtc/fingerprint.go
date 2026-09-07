package libp2pwebrtc

import (
	"crypto"
	"crypto/x509"
	"errors"

	ma "github.com/multiformats/go-multiaddr"
	mh "github.com/multiformats/go-multihash"
	"github.com/pion/webrtc/v4"
)

var errHashUnavailable = errors.New("fingerprint: hash algorithm is not linked into the binary")

func parseFingerprint(cert *x509.Certificate, algo crypto.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeRemoteFingerprint(maddr ma.Multiaddr) (*mh.DecodedMultihash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeDTLSFingerprint(fp webrtc.DTLSFingerprint) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
