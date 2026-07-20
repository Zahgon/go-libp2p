package httppeeridauth

import (
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/http/auth/internal/handshake"
)

type ClientPeerIDAuth struct {
	PrivKey  crypto.PrivKey
	TokenTTL time.Duration

	tm tokenMap
}

type clientAsRoundTripper struct {
	*http.Client
}

func (c clientAsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ClientPeerIDAuth) AuthenticatedDo(client *http.Client, req *http.Request) (peer.ID, *http.Response, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil, nil
}

func (a *ClientPeerIDAuth) AuthenticateWithRoundTripper(rt http.RoundTripper, req *http.Request) (peer.ID, *http.Response, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil, nil
}

func (a *ClientPeerIDAuth) HasToken(hostname string) bool { _ = "STUB: not implemented"; return false }

func (a *ClientPeerIDAuth) runHandshake(rt http.RoundTripper, req *http.Request, b bodyMeta, hs *handshake.PeerIDAuthHandshakeClient) (peer.ID, *http.Response, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil, nil
}

var errTokenRejected = errors.New("token rejected")

func (a *ClientPeerIDAuth) doWithToken(rt http.RoundTripper, req *http.Request, ti tokenInfo) (peer.ID, *http.Response, error) {
	_ = "STUB: not implemented"
	return *new(peer.ID), nil, nil
}

type bodyMeta struct {
	body          io.ReadCloser
	contentLength int64
	getBody       func() (io.ReadCloser, error)
}

func clearBody(req *http.Request) bodyMeta { _ = "STUB: not implemented"; return *new(bodyMeta) }

func (b *bodyMeta) setBody(req *http.Request) { _ = "STUB: not implemented"; return }

type tokenInfo struct {
	token      string
	insertedAt time.Time
	peerID     peer.ID
}

type tokenMap struct {
	tokenMapMu sync.Mutex
	tokenMap   map[string]tokenInfo
}

func (tm *tokenMap) get(hostname string, ttl time.Duration) (tokenInfo, bool) {
	_ = "STUB: not implemented"
	return *new(tokenInfo), false
}

func (tm *tokenMap) set(hostname string, ti tokenInfo) { _ = "STUB: not implemented"; return }
