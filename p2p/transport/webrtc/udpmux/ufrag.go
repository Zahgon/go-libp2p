package udpmux

const (
	UfragPrefixV1 = "libp2p+webrtc+v1/"
	UfragPrefixV2 = "libp2p+webrtc+v2/"
)

const (
	iceUfragMinLen   = 4
	icePwdMinLen     = 22
	iceCredentialMax = 256
)

func isICEChar(b byte) bool { _ = "STUB: not implemented"; return false }

func isICECharString(s string) bool { _ = "STUB: not implemented"; return false }

func isICEUfrag(s string) bool { _ = "STUB: not implemented"; return false }

func isICEPwd(s string) bool { _ = "STUB: not implemented"; return false }
