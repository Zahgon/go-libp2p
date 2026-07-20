package rcmgr

import (
	"math"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type baseLimitConfig struct {
	BaseLimit         BaseLimit
	BaseLimitIncrease BaseLimitIncrease
}

type ScalingLimitConfig struct {
	SystemBaseLimit     BaseLimit
	SystemLimitIncrease BaseLimitIncrease

	TransientBaseLimit     BaseLimit
	TransientLimitIncrease BaseLimitIncrease

	AllowlistedSystemBaseLimit     BaseLimit
	AllowlistedSystemLimitIncrease BaseLimitIncrease

	AllowlistedTransientBaseLimit     BaseLimit
	AllowlistedTransientLimitIncrease BaseLimitIncrease

	ServiceBaseLimit     BaseLimit
	ServiceLimitIncrease BaseLimitIncrease
	ServiceLimits        map[string]baseLimitConfig

	ServicePeerBaseLimit     BaseLimit
	ServicePeerLimitIncrease BaseLimitIncrease
	ServicePeerLimits        map[string]baseLimitConfig

	ProtocolBaseLimit     BaseLimit
	ProtocolLimitIncrease BaseLimitIncrease
	ProtocolLimits        map[protocol.ID]baseLimitConfig

	ProtocolPeerBaseLimit     BaseLimit
	ProtocolPeerLimitIncrease BaseLimitIncrease
	ProtocolPeerLimits        map[protocol.ID]baseLimitConfig

	PeerBaseLimit     BaseLimit
	PeerLimitIncrease BaseLimitIncrease
	PeerLimits        map[peer.ID]baseLimitConfig

	ConnBaseLimit     BaseLimit
	ConnLimitIncrease BaseLimitIncrease

	StreamBaseLimit     BaseLimit
	StreamLimitIncrease BaseLimitIncrease
}

func (cfg *ScalingLimitConfig) AddServiceLimit(svc string, base BaseLimit, inc BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}

func (cfg *ScalingLimitConfig) AddProtocolLimit(proto protocol.ID, base BaseLimit, inc BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}

func (cfg *ScalingLimitConfig) AddPeerLimit(p peer.ID, base BaseLimit, inc BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}

func (cfg *ScalingLimitConfig) AddServicePeerLimit(svc string, base BaseLimit, inc BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}

func (cfg *ScalingLimitConfig) AddProtocolPeerLimit(proto protocol.ID, base BaseLimit, inc BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}

type LimitVal int

const (
	DefaultLimit LimitVal = 0

	Unlimited LimitVal = -1

	BlockAllLimit LimitVal = -2
)

func (l LimitVal) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *LimitVal) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (l LimitVal) Build(defaultVal int) int { _ = "STUB: not implemented"; return 0 }

type LimitVal64 int64

const (
	DefaultLimit64 LimitVal64 = 0

	Unlimited64 LimitVal64 = -1

	BlockAllLimit64 LimitVal64 = -2
)

func (l LimitVal64) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *LimitVal64) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (l LimitVal64) Build(defaultVal int64) int64 { _ = "STUB: not implemented"; return 0 }

type ResourceLimits struct {
	Streams         LimitVal   `json:",omitempty"`
	StreamsInbound  LimitVal   `json:",omitempty"`
	StreamsOutbound LimitVal   `json:",omitempty"`
	Conns           LimitVal   `json:",omitempty"`
	ConnsInbound    LimitVal   `json:",omitempty"`
	ConnsOutbound   LimitVal   `json:",omitempty"`
	FD              LimitVal   `json:",omitempty"`
	Memory          LimitVal64 `json:",omitempty"`
}

func (l *ResourceLimits) IsDefault() bool { _ = "STUB: not implemented"; return false }

func (l *ResourceLimits) ToMaybeNilPtr() *ResourceLimits { _ = "STUB: not implemented"; return nil }

func (l *ResourceLimits) Apply(l2 ResourceLimits) { _ = "STUB: not implemented"; return }

func (l *ResourceLimits) Build(defaults Limit) BaseLimit {
	_ = "STUB: not implemented"
	return *new(BaseLimit)
}

type PartialLimitConfig struct {
	System    ResourceLimits
	Transient ResourceLimits

	AllowlistedSystem    ResourceLimits
	AllowlistedTransient ResourceLimits

	ServiceDefault ResourceLimits
	Service        map[string]ResourceLimits `json:",omitempty"`

	ServicePeerDefault ResourceLimits
	ServicePeer        map[string]ResourceLimits `json:",omitempty"`

	ProtocolDefault ResourceLimits
	Protocol        map[protocol.ID]ResourceLimits `json:",omitempty"`

	ProtocolPeerDefault ResourceLimits
	ProtocolPeer        map[protocol.ID]ResourceLimits `json:",omitempty"`

	PeerDefault ResourceLimits
	Peer        map[peer.ID]ResourceLimits `json:",omitempty"`

	Conn   ResourceLimits
	Stream ResourceLimits
}

func (cfg *PartialLimitConfig) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyResourceLimitsMap[K comparable](this *map[K]ResourceLimits, other map[K]ResourceLimits, fallbackDefault ResourceLimits) {
	_ = "STUB: not implemented"
	return
}

func (cfg *PartialLimitConfig) Apply(c PartialLimitConfig) { _ = "STUB: not implemented"; return }

func (cfg PartialLimitConfig) Build(defaults ConcreteLimitConfig) ConcreteLimitConfig {
	_ = "STUB: not implemented"
	return *new(ConcreteLimitConfig)
}

func buildMapWithDefault[K comparable](definedLimits map[K]ResourceLimits, defaults map[K]BaseLimit, fallbackDefault BaseLimit) map[K]BaseLimit {
	_ = "STUB: not implemented"
	return nil
}

type ConcreteLimitConfig struct {
	system    BaseLimit
	transient BaseLimit

	allowlistedSystem    BaseLimit
	allowlistedTransient BaseLimit

	serviceDefault BaseLimit
	service        map[string]BaseLimit

	servicePeerDefault BaseLimit
	servicePeer        map[string]BaseLimit

	protocolDefault BaseLimit
	protocol        map[protocol.ID]BaseLimit

	protocolPeerDefault BaseLimit
	protocolPeer        map[protocol.ID]BaseLimit

	peerDefault BaseLimit
	peer        map[peer.ID]BaseLimit

	conn   BaseLimit
	stream BaseLimit
}

func resourceLimitsMapFromBaseLimitMap[K comparable](baseLimitMap map[K]BaseLimit) map[K]ResourceLimits {
	_ = "STUB: not implemented"
	return nil
}

func (cfg ConcreteLimitConfig) ToPartialLimitConfig() PartialLimitConfig {
	_ = "STUB: not implemented"
	return *new(PartialLimitConfig)
}

func (cfg *ScalingLimitConfig) Scale(memory int64, numFD int) ConcreteLimitConfig {
	_ = "STUB: not implemented"
	return *new(ConcreteLimitConfig)
}

func (cfg *ScalingLimitConfig) AutoScale() ConcreteLimitConfig {
	_ = "STUB: not implemented"
	return *new(ConcreteLimitConfig)
}

func scale(base BaseLimit, inc BaseLimitIncrease, memory int64, numFD int) BaseLimit {
	_ = "STUB: not implemented"
	return *new(BaseLimit)
}

var DefaultLimits = ScalingLimitConfig{
	SystemBaseLimit: BaseLimit{
		ConnsInbound:    64,
		ConnsOutbound:   128,
		Conns:           128,
		StreamsInbound:  64 * 16,
		StreamsOutbound: 128 * 16,
		Streams:         128 * 16,
		Memory:          128 << 20,
		FD:              256,
	},

	SystemLimitIncrease: BaseLimitIncrease{
		ConnsInbound:    64,
		ConnsOutbound:   128,
		Conns:           128,
		StreamsInbound:  64 * 16,
		StreamsOutbound: 128 * 16,
		Streams:         128 * 16,
		Memory:          1 << 30,
		FDFraction:      1,
	},

	TransientBaseLimit: BaseLimit{
		ConnsInbound:    32,
		ConnsOutbound:   64,
		Conns:           64,
		StreamsInbound:  128,
		StreamsOutbound: 256,
		Streams:         256,
		Memory:          32 << 20,
		FD:              64,
	},

	TransientLimitIncrease: BaseLimitIncrease{
		ConnsInbound:    16,
		ConnsOutbound:   32,
		Conns:           32,
		StreamsInbound:  128,
		StreamsOutbound: 256,
		Streams:         256,
		Memory:          128 << 20,
		FDFraction:      0.25,
	},

	AllowlistedSystemBaseLimit: BaseLimit{
		ConnsInbound:    64,
		ConnsOutbound:   128,
		Conns:           128,
		StreamsInbound:  64 * 16,
		StreamsOutbound: 128 * 16,
		Streams:         128 * 16,
		Memory:          128 << 20,
		FD:              256,
	},

	AllowlistedSystemLimitIncrease: BaseLimitIncrease{
		ConnsInbound:    64,
		ConnsOutbound:   128,
		Conns:           128,
		StreamsInbound:  64 * 16,
		StreamsOutbound: 128 * 16,
		Streams:         128 * 16,
		Memory:          1 << 30,
		FDFraction:      1,
	},

	AllowlistedTransientBaseLimit: BaseLimit{
		ConnsInbound:    32,
		ConnsOutbound:   64,
		Conns:           64,
		StreamsInbound:  128,
		StreamsOutbound: 256,
		Streams:         256,
		Memory:          32 << 20,
		FD:              64,
	},

	AllowlistedTransientLimitIncrease: BaseLimitIncrease{
		ConnsInbound:    16,
		ConnsOutbound:   32,
		Conns:           32,
		StreamsInbound:  128,
		StreamsOutbound: 256,
		Streams:         256,
		Memory:          128 << 20,
		FDFraction:      0.25,
	},

	ServiceBaseLimit: BaseLimit{
		StreamsInbound:  1024,
		StreamsOutbound: 4096,
		Streams:         4096,
		Memory:          64 << 20,
	},

	ServiceLimitIncrease: BaseLimitIncrease{
		StreamsInbound:  512,
		StreamsOutbound: 2048,
		Streams:         2048,
		Memory:          128 << 20,
	},

	ServicePeerBaseLimit: BaseLimit{
		StreamsInbound:  128,
		StreamsOutbound: 256,
		Streams:         256,
		Memory:          16 << 20,
	},

	ServicePeerLimitIncrease: BaseLimitIncrease{
		StreamsInbound:  4,
		StreamsOutbound: 8,
		Streams:         8,
		Memory:          4 << 20,
	},

	ProtocolBaseLimit: BaseLimit{
		StreamsInbound:  512,
		StreamsOutbound: 2048,
		Streams:         2048,
		Memory:          64 << 20,
	},

	ProtocolLimitIncrease: BaseLimitIncrease{
		StreamsInbound:  256,
		StreamsOutbound: 512,
		Streams:         512,
		Memory:          164 << 20,
	},

	ProtocolPeerBaseLimit: BaseLimit{
		StreamsInbound:  64,
		StreamsOutbound: 128,
		Streams:         256,
		Memory:          16 << 20,
	},

	ProtocolPeerLimitIncrease: BaseLimitIncrease{
		StreamsInbound:  4,
		StreamsOutbound: 8,
		Streams:         16,
		Memory:          4,
	},

	PeerBaseLimit: BaseLimit{

		ConnsInbound:    8,
		ConnsOutbound:   8,
		Conns:           8,
		StreamsInbound:  256,
		StreamsOutbound: 512,
		Streams:         512,
		Memory:          64 << 20,
		FD:              4,
	},

	PeerLimitIncrease: BaseLimitIncrease{
		StreamsInbound:  128,
		StreamsOutbound: 256,
		Streams:         256,
		Memory:          128 << 20,
		FDFraction:      1.0 / 64,
	},

	ConnBaseLimit: BaseLimit{
		ConnsInbound:  1,
		ConnsOutbound: 1,
		Conns:         1,
		FD:            1,
		Memory:        32 << 20,
	},

	StreamBaseLimit: BaseLimit{
		StreamsInbound:  1,
		StreamsOutbound: 1,
		Streams:         1,
		Memory:          16 << 20,
	},
}

var infiniteBaseLimit = BaseLimit{
	Streams:         math.MaxInt,
	StreamsInbound:  math.MaxInt,
	StreamsOutbound: math.MaxInt,
	Conns:           math.MaxInt,
	ConnsInbound:    math.MaxInt,
	ConnsOutbound:   math.MaxInt,
	FD:              math.MaxInt,
	Memory:          math.MaxInt64,
}

var InfiniteLimits = ConcreteLimitConfig{
	system:               infiniteBaseLimit,
	transient:            infiniteBaseLimit,
	allowlistedSystem:    infiniteBaseLimit,
	allowlistedTransient: infiniteBaseLimit,
	serviceDefault:       infiniteBaseLimit,
	servicePeerDefault:   infiniteBaseLimit,
	protocolDefault:      infiniteBaseLimit,
	protocolPeerDefault:  infiniteBaseLimit,
	peerDefault:          infiniteBaseLimit,
	conn:                 infiniteBaseLimit,
	stream:               infiniteBaseLimit,
}
