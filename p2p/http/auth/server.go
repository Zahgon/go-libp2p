package httppeeridauth

import (
	"hash"
	"net/http"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

type hmacPool struct {
	p sync.Pool
}

func newHmacPool(key []byte) *hmacPool { _ = "STUB: not implemented"; return nil }

func (p *hmacPool) Get() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (p *hmacPool) Put(h hash.Hash) { _ = "STUB: not implemented"; return }

type ServerPeerIDAuth struct {
	PrivKey  crypto.PrivKey
	TokenTTL time.Duration
	Next     func(peer peer.ID, w http.ResponseWriter, r *http.Request)

	NoTLS bool

	ValidHostnameFn func(hostname string) bool

	HmacKey  []byte
	initHmac sync.Once
	hmacPool *hmacPool
}

func (a *ServerPeerIDAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *ServerPeerIDAuth) ServeHTTPWithNextHandler(w http.ResponseWriter, r *http.Request, next func(peer.ID, http.ResponseWriter, *http.Request)) {
	_ = "STUB: not implemented"
	return
}

func HasAuthHeader(r *http.Request) bool { _ = "STUB: not implemented"; return false }
