package event

import (
	"io"
	"reflect"
)

type SubscriptionOpt = func(any) error

type EmitterOpt = func(any) error

type CancelFunc = func()

type wildcardSubscriptionType any

var WildcardSubscription = new(wildcardSubscriptionType)

type Emitter interface {
	io.Closer

	Emit(evt any) error
}

type Subscription interface {
	io.Closer

	Out() <-chan any

	Name() string
}

type Bus interface {
	Subscribe(eventType any, opts ...SubscriptionOpt) (Subscription, error)

	Emitter(eventType any, opts ...EmitterOpt) (Emitter, error)

	GetAllEventTypes() []reflect.Type
}
