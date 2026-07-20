package metricshelper

import ma "github.com/multiformats/go-multiaddr"

var transports = [...]int{ma.P_CIRCUIT, ma.P_WEBRTC, ma.P_WEBRTC_DIRECT, ma.P_WEBTRANSPORT, ma.P_QUIC, ma.P_QUIC_V1, ma.P_WSS, ma.P_WS, ma.P_TCP}

func GetTransport(a ma.Multiaddr) string { _ = "STUB: not implemented"; return "" }

func GetIPVersion(addr ma.Multiaddr) string { _ = "STUB: not implemented"; return "" }
