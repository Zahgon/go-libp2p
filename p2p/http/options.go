package libp2phttp

type RoundTripperOption func(o roundTripperOpts) roundTripperOpts

type roundTripperOpts struct {
	preferHTTPTransport          bool
	serverMustAuthenticatePeerID bool
}

func PreferHTTPTransport(o roundTripperOpts) roundTripperOpts {
	_ = "STUB: not implemented"
	return *new(roundTripperOpts)
}

func ServerMustAuthenticatePeerID(o roundTripperOpts) roundTripperOpts {
	_ = "STUB: not implemented"
	return *new(roundTripperOpts)
}
