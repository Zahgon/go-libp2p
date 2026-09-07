package libp2pwebrtc

import (
	"crypto"
	"net"

	"github.com/multiformats/go-multihash"
)

const clientSDP = `v=0
o=- 0 0 IN %[1]s %[2]s
s=-
c=IN %[1]s %[2]s
t=0 0

m=application %[3]d UDP/DTLS/SCTP webrtc-datachannel
a=mid:0
a=ice-options:ice2
a=ice-ufrag:%[4]s
a=ice-pwd:%[5]s
a=fingerprint:sha-256 ba:78:16:bf:8f:01:cf:ea:41:41:40:de:5d:ae:22:23:b0:03:61:a3:96:17:7a:9c:b4:10:ff:61:f2:00:15:ad
a=setup:actpass
a=sctp-port:5000
a=max-message-size:16384
`

func createClientSDP(addr *net.UDPAddr, ufrag, pwd string) string {
	_ = "STUB: not implemented"
	return ""
}

const serverSDP = `v=0
o=- 0 0 IN %[1]s %[2]s
s=-
t=0 0
a=ice-lite
m=application %[3]d UDP/DTLS/SCTP webrtc-datachannel
c=IN %[1]s %[2]s
a=mid:0
a=ice-options:ice2
a=ice-ufrag:%[4]s
a=ice-pwd:%[4]s
a=fingerprint:%[5]s

a=setup:passive
a=sctp-port:5000
a=max-message-size:16384
a=candidate:1 1 UDP 1 %[2]s %[3]d typ host
a=end-of-candidates
`

func createServerSDP(addr *net.UDPAddr, ufrag string, fingerprint multihash.DecodedMultihash) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getSupportedSDPHash(code uint64) (crypto.Hash, bool) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), false
}

func getSupportedSDPString(code uint64) (string, error) { _ = "STUB: not implemented"; return "", nil }
