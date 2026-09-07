package connmgr

import (
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
)

type Decayer interface {
	io.Closer

	RegisterDecayingTag(name string, interval time.Duration, decayFn DecayFn, bumpFn BumpFn) (DecayingTag, error)
}

type DecayFn func(value DecayingValue) (after int, rm bool)

type BumpFn func(value DecayingValue, delta int) (after int)

type DecayingTag interface {
	Name() string

	Interval() time.Duration

	Bump(peer peer.ID, delta int) error

	Remove(peer peer.ID) error

	Close() error
}

type DecayingValue struct {
	Tag DecayingTag

	Peer peer.ID

	Added time.Time

	LastVisit time.Time

	Value int
}
