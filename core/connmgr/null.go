package connmgr

import (
	"context"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

type NullConnMgr struct{}

var _ ConnManager = (*NullConnMgr)(nil)

func (NullConnMgr) TagPeer(peer.ID, string, int)             { _ = "STUB: not implemented"; return }
func (NullConnMgr) UntagPeer(peer.ID, string)                { _ = "STUB: not implemented"; return }
func (NullConnMgr) UpsertTag(peer.ID, string, func(int) int) { _ = "STUB: not implemented"; return }
func (NullConnMgr) GetTagInfo(peer.ID) *TagInfo              { _ = "STUB: not implemented"; return nil }
func (NullConnMgr) TrimOpenConns(_ context.Context)          { _ = "STUB: not implemented"; return }
func (NullConnMgr) Notifee() network.Notifiee {
	_ = "STUB: not implemented"
	return *new(network.Notifiee)
}
func (NullConnMgr) Protect(peer.ID, string)           { _ = "STUB: not implemented"; return }
func (NullConnMgr) Unprotect(peer.ID, string) bool    { _ = "STUB: not implemented"; return false }
func (NullConnMgr) IsProtected(peer.ID, string) bool  { _ = "STUB: not implemented"; return false }
func (NullConnMgr) CheckLimit(_ GetConnLimiter) error { _ = "STUB: not implemented"; return nil }
func (NullConnMgr) Close() error                      { _ = "STUB: not implemented"; return nil }
