package autonatv2

import "time"

type autoNATSettings struct {
	allowPrivateAddrs                    bool
	serverRPM                            int
	serverPerPeerRPM                     int
	serverDialDataRPM                    int
	maxConcurrentRequestsPerPeer         int
	dataRequestPolicy                    dataRequestPolicyFunc
	now                                  func() time.Time
	amplificatonAttackPreventionDialWait time.Duration
	metricsTracer                        MetricsTracer
	throttlePeerDuration                 time.Duration
}

func defaultSettings() *autoNATSettings { _ = "STUB: not implemented"; return nil }

type AutoNATOption func(s *autoNATSettings) error

func WithServerRateLimit(rpm, perPeerRPM, dialDataRPM int, maxConcurrentRequestsPerPeer int) AutoNATOption {
	_ = "STUB: not implemented"
	return *new(AutoNATOption)
}

func WithMetricsTracer(m MetricsTracer) AutoNATOption {
	_ = "STUB: not implemented"
	return *new(AutoNATOption)
}

func withDataRequestPolicy(drp dataRequestPolicyFunc) AutoNATOption {
	_ = "STUB: not implemented"
	return *new(AutoNATOption)
}

func AllowPrivateAddrs(s *autoNATSettings) error { _ = "STUB: not implemented"; return nil }

func withAmplificationAttackPreventionDialWait(d time.Duration) AutoNATOption {
	_ = "STUB: not implemented"
	return *new(AutoNATOption)
}

func withThrottlePeerDuration(d time.Duration) AutoNATOption {
	_ = "STUB: not implemented"
	return *new(AutoNATOption)
}
