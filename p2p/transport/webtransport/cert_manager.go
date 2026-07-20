package libp2pwebtransport

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	ma "github.com/multiformats/go-multiaddr"
)

const clockSkewAllowance = time.Hour
const validityMinusTwoSkew = certValidity - (2 * clockSkewAllowance)

type certConfig struct {
	tlsConf *tls.Config
	sha256  [32]byte
}

func (c *certConfig) Start() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (c *certConfig) End() time.Time   { _ = "STUB: not implemented"; return *new(time.Time) }

func newCertConfig(key ic.PrivKey, start, end time.Time) (*certConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type certManager struct {
	clock     clock.Clock
	ctx       context.Context
	ctxCancel context.CancelFunc
	refCount  sync.WaitGroup

	mx            sync.RWMutex
	lastConfig    *certConfig
	currentConfig *certConfig
	nextConfig    *certConfig
	addrComp      ma.Multiaddr

	serializedCertHashes [][]byte
}

func newCertManager(hostKey ic.PrivKey, clock clock.Clock) (*certManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCurrentBucketStartTime(now time.Time, offset time.Duration) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (m *certManager) init(hostKey ic.PrivKey) error {
	start := m.clock.Now()
	pubkeyBytes, err := hostKey.GetPublic().Raw()
	if err != nil {
		return err
	}

	offset := (time.Duration(binary.LittleEndian.Uint16(pubkeyBytes)) * time.Minute) % certValidity

	start = start.Add(-clockSkewAllowance)
	startTime := getCurrentBucketStartTime(start, offset)
	m.nextConfig, err = newCertConfig(hostKey, startTime, startTime.Add(certValidity))
	if err != nil {
		return err
	}
	return m.rollConfig(hostKey)
}

func (m *certManager) rollConfig(hostKey ic.PrivKey) error { _ = "STUB: not implemented"; return nil }

func (m *certManager) background(hostKey ic.PrivKey) { _ = "STUB: not implemented"; return }

func (m *certManager) GetConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

func (m *certManager) AddrComponent() ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

func (m *certManager) SerializedCertHashes() [][]byte { _ = "STUB: not implemented"; return nil }

func (m *certManager) cacheSerializedCertHashes() error { _ = "STUB: not implemented"; return nil }

func (m *certManager) cacheAddrComponent() error { _ = "STUB: not implemented"; return nil }

func (m *certManager) Close() error { _ = "STUB: not implemented"; return nil }
