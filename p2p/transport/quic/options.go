package libp2pquic

import p2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"

type Option func(o *transportConfig) error

type transportConfig struct {
	tlsIdentityOpts []p2ptls.IdentityOption
}

func WithTLSIdentityOption(opt ...p2ptls.IdentityOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
