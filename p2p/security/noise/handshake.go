package noise

import (
	"context"

	"github.com/libp2p/go-libp2p/p2p/security/noise/pb"

	"github.com/flynn/noise"
)

const payloadSigPrefix = "noise-libp2p-static-key:"

var cipherSuite = noise.NewCipherSuite(noise.DH25519, noise.CipherChaChaPoly, noise.HashSHA256)

func (s *secureSession) runHandshake(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *secureSession) setCipherStates(cs1, cs2 *noise.CipherState) {
	_ = "STUB: not implemented"
	return
}

func (s *secureSession) sendHandshakeMessage(hs *noise.HandshakeState, payload []byte, hbuf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *secureSession) readHandshakeMessage(hs *noise.HandshakeState) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *secureSession) generateHandshakePayload(localStatic noise.DHKey, ext *pb.NoiseExtensions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *secureSession) handleRemoteHandshakePayload(payload []byte, remoteStatic []byte) (*pb.NoiseExtensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
