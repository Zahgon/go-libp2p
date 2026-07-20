package network

import (
	"errors"
	"net"
)

type temporaryError string

func (e temporaryError) Error() string   { _ = "STUB: not implemented"; return "" }
func (e temporaryError) Temporary() bool { _ = "STUB: not implemented"; return false }
func (e temporaryError) Timeout() bool   { _ = "STUB: not implemented"; return false }

var _ net.Error = temporaryError("")

var ErrNoRemoteAddrs = errors.New("no remote addresses")

var ErrNoConn = errors.New("no usable connection to peer")

var ErrTransientConn = ErrLimitedConn

var ErrLimitedConn = errors.New("limited connection to peer")

var ErrResourceLimitExceeded = temporaryError("resource limit exceeded")

var ErrResourceScopeClosed = errors.New("resource scope closed")
