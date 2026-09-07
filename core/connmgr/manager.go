package connmgr

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

func SupportsDecay(mgr ConnManager) (Decayer, bool) {
	_ = "STUB: not implemented"
	return *new(Decayer), false
}

type ConnManager interface {
	TagPeer(peer.ID, string, int)

	UntagPeer(p peer.ID, tag string)

	UpsertTag(p peer.ID, tag string, upsert func(int) int)

	GetTagInfo(p peer.ID) *TagInfo

	TrimOpenConns(ctx context.Context)

	Notifee() network.Notifiee

	Protect(id peer.ID, tag string)

	Unprotect(id peer.ID, tag string) (protected bool)

	IsProtected(id peer.ID, tag string) (protected bool)

	CheckLimit(l GetConnLimiter) error

	Close() error
}

type TagInfo struct {
	FirstSeen time.Time
	Value     int

	Tags map[string]int

	Conns map[string]time.Time
}

type GetConnLimiter interface {
	GetConnLimit() int
}
