package event

import "github.com/libp2p/go-libp2p/core/network"

type EvtNATDeviceTypeChanged struct {
	TransportProtocol network.NATTransportProtocol

	NatDeviceType network.NATDeviceType
}
