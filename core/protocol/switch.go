package protocol

import (
	"io"

	"github.com/multiformats/go-multistream"
)

type HandlerFunc = multistream.HandlerFunc[ID]

type Router interface {
	AddHandler(protocol ID, handler HandlerFunc)

	AddHandlerWithFunc(protocol ID, match func(ID) bool, handler HandlerFunc)

	RemoveHandler(protocol ID)

	Protocols() []ID
}

type Negotiator interface {
	Negotiate(rwc io.ReadWriteCloser) (ID, HandlerFunc, error)

	Handle(rwc io.ReadWriteCloser) error
}

type Switch interface {
	Router
	Negotiator
}
