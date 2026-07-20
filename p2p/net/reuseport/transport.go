package reuseport

import (
	"errors"
	"sync"

	logging "github.com/libp2p/go-libp2p/gologshim"
)

var log = logging.Logger("reuseport-transport")

var ErrWrongProto = errors.New("can only dial TCP over IPv4 or IPv6")

type Transport struct {
	v4 network
	v6 network
}

type network struct {
	mu        sync.RWMutex
	listeners map[*listener]struct{}
	dialer    *dialer
}
