package libp2p

import (
	"github.com/libp2p/go-libp2p/core/protocol"
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"
)

func SetDefaultServiceLimits(config *rcmgr.ScalingLimitConfig) { _ = "STUB: not implemented"; return }

func addServiceAndProtocolLimit(config *rcmgr.ScalingLimitConfig, service string, proto protocol.ID, limit rcmgr.BaseLimit, increase rcmgr.BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}

func addServicePeerAndProtocolPeerLimit(config *rcmgr.ScalingLimitConfig, service string, proto protocol.ID, limit rcmgr.BaseLimit, increase rcmgr.BaseLimitIncrease) {
	_ = "STUB: not implemented"
	return
}
