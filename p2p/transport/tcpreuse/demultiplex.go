package tcpreuse

import (
	"time"

	manet "github.com/multiformats/go-multiaddr/net"
)

var identifyConnTimeout = 5 * time.Second

type DemultiplexedConnType int

const (
	DemultiplexedConnType_Unknown DemultiplexedConnType = iota
	DemultiplexedConnType_MultistreamSelect
	DemultiplexedConnType_HTTP
	DemultiplexedConnType_TLS
)

func (t DemultiplexedConnType) String() string { _ = "STUB: not implemented"; return "" }

func (t DemultiplexedConnType) IsKnown() bool { _ = "STUB: not implemented"; return false }

func identifyConnType(c manet.Conn) (DemultiplexedConnType, manet.Conn, error) {
	_ = "STUB: not implemented"
	return *new(DemultiplexedConnType), *new(manet.Conn), nil
}

type Prefix = [3]byte

func IsMultistreamSelect(s Prefix) bool { _ = "STUB: not implemented"; return false }

func IsHTTP(s Prefix) bool { _ = "STUB: not implemented"; return false }

func IsTLS(s Prefix) bool { _ = "STUB: not implemented"; return false }
