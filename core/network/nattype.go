package network

type NATDeviceType int

const (
	NATDeviceTypeUnknown NATDeviceType = iota

	NATDeviceTypeEndpointIndependent

	NATDeviceTypeEndpointDependent
)

const (
	NATDeviceTypeCone = NATDeviceTypeEndpointIndependent

	NATDeviceTypeSymmetric = NATDeviceTypeEndpointDependent
)

func (r NATDeviceType) String() string { _ = "STUB: not implemented"; return "" }

type NATTransportProtocol int

const (
	NATTransportUDP NATTransportProtocol = iota

	NATTransportTCP
)

func (n NATTransportProtocol) String() string { _ = "STUB: not implemented"; return "" }
