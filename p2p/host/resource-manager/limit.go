package rcmgr

import (
	"io"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type Limit interface {
	GetMemoryLimit() int64

	GetStreamLimit(network.Direction) int

	GetStreamTotalLimit() int

	GetConnLimit(network.Direction) int

	GetConnTotalLimit() int

	GetFDLimit() int
}

type Limiter interface {
	GetSystemLimits() Limit
	GetTransientLimits() Limit
	GetAllowlistedSystemLimits() Limit
	GetAllowlistedTransientLimits() Limit
	GetServiceLimits(svc string) Limit
	GetServicePeerLimits(svc string) Limit
	GetProtocolLimits(proto protocol.ID) Limit
	GetProtocolPeerLimits(proto protocol.ID) Limit
	GetPeerLimits(p peer.ID) Limit
	GetStreamLimits(p peer.ID) Limit
	GetConnLimits() Limit
}

func NewDefaultLimiterFromJSON(in io.Reader) (Limiter, error) {
	_ = "STUB: not implemented"
	return *new(Limiter), nil
}

func NewLimiterFromJSON(in io.Reader, defaults ConcreteLimitConfig) (Limiter, error) {
	_ = "STUB: not implemented"
	return *new(Limiter), nil
}

func readLimiterConfigFromJSON(in io.Reader, defaults ConcreteLimitConfig) (ConcreteLimitConfig, error) {
	_ = "STUB: not implemented"
	return *new(ConcreteLimitConfig), nil
}

type fixedLimiter struct {
	ConcreteLimitConfig
}

var _ Limiter = (*fixedLimiter)(nil)

func NewFixedLimiter(conf ConcreteLimitConfig) Limiter {
	_ = "STUB: not implemented"
	return *new(Limiter)
}

type BaseLimit struct {
	Streams         int   `json:",omitempty"`
	StreamsInbound  int   `json:",omitempty"`
	StreamsOutbound int   `json:",omitempty"`
	Conns           int   `json:",omitempty"`
	ConnsInbound    int   `json:",omitempty"`
	ConnsOutbound   int   `json:",omitempty"`
	FD              int   `json:",omitempty"`
	Memory          int64 `json:",omitempty"`
}

func valueOrBlockAll(n int) LimitVal { _ = "STUB: not implemented"; return *new(LimitVal) }

func valueOrBlockAll64(n int64) LimitVal64 { _ = "STUB: not implemented"; return *new(LimitVal64) }

func (l BaseLimit) ToResourceLimits() ResourceLimits {
	_ = "STUB: not implemented"
	return *new(ResourceLimits)
}

func (l *BaseLimit) Apply(l2 BaseLimit) { _ = "STUB: not implemented"; return }

type BaseLimitIncrease struct {
	Streams         int `json:",omitempty"`
	StreamsInbound  int `json:",omitempty"`
	StreamsOutbound int `json:",omitempty"`
	Conns           int `json:",omitempty"`
	ConnsInbound    int `json:",omitempty"`
	ConnsOutbound   int `json:",omitempty"`

	Memory int64 `json:",omitempty"`

	FDFraction float64 `json:",omitempty"`
}

func (l *BaseLimitIncrease) Apply(l2 BaseLimitIncrease) { _ = "STUB: not implemented"; return }

func (l BaseLimit) GetStreamLimit(dir network.Direction) int { _ = "STUB: not implemented"; return 0 }

func (l BaseLimit) GetStreamTotalLimit() int { _ = "STUB: not implemented"; return 0 }

func (l BaseLimit) GetConnLimit(dir network.Direction) int { _ = "STUB: not implemented"; return 0 }

func (l BaseLimit) GetConnTotalLimit() int { _ = "STUB: not implemented"; return 0 }

func (l BaseLimit) GetFDLimit() int { _ = "STUB: not implemented"; return 0 }

func (l BaseLimit) GetMemoryLimit() int64 { _ = "STUB: not implemented"; return 0 }

func (l *fixedLimiter) GetSystemLimits() Limit { _ = "STUB: not implemented"; return *new(Limit) }

func (l *fixedLimiter) GetTransientLimits() Limit { _ = "STUB: not implemented"; return *new(Limit) }

func (l *fixedLimiter) GetAllowlistedSystemLimits() Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetAllowlistedTransientLimits() Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetServiceLimits(svc string) Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetServicePeerLimits(svc string) Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetProtocolLimits(proto protocol.ID) Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetProtocolPeerLimits(proto protocol.ID) Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetPeerLimits(p peer.ID) Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetStreamLimits(_ peer.ID) Limit {
	_ = "STUB: not implemented"
	return *new(Limit)
}

func (l *fixedLimiter) GetConnLimits() Limit { _ = "STUB: not implemented"; return *new(Limit) }
