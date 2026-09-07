package network

import (
	"context"
	"time"
)

var DialPeerTimeout = 60 * time.Second

type noDialCtxKey struct{}
type dialPeerTimeoutCtxKey struct{}
type forceDirectDialCtxKey struct{}
type allowLimitedConnCtxKey struct{}
type simConnectCtxKey struct{ isClient bool }

var noDial = noDialCtxKey{}
var forceDirectDial = forceDirectDialCtxKey{}
var allowLimitedConn = allowLimitedConnCtxKey{}
var simConnectIsServer = simConnectCtxKey{}
var simConnectIsClient = simConnectCtxKey{isClient: true}

func WithForceDirectDial(ctx context.Context, reason string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetForceDirectDial(ctx context.Context) (forceDirect bool, reason string) {
	_ = "STUB: not implemented"
	return false, ""
}

func WithSimultaneousConnect(ctx context.Context, isClient bool, reason string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetSimultaneousConnect(ctx context.Context) (simconnect bool, isClient bool, reason string) {
	_ = "STUB: not implemented"
	return false, false, ""
}

func WithNoDial(ctx context.Context, reason string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetNoDial(ctx context.Context) (nodial bool, reason string) {
	_ = "STUB: not implemented"
	return false, ""
}

func GetDialPeerTimeout(ctx context.Context) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func WithDialPeerTimeout(ctx context.Context, timeout time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithAllowLimitedConn(ctx context.Context, reason string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithUseTransient(ctx context.Context, reason string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetAllowLimitedConn(ctx context.Context) (usetransient bool, reason string) {
	_ = "STUB: not implemented"
	return false, ""
}

func GetUseTransient(ctx context.Context) (usetransient bool, reason string) {
	_ = "STUB: not implemented"
	return false, ""
}
