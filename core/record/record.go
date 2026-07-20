package record

import (
	"errors"
	"reflect"
)

var (
	ErrPayloadTypeNotRegistered = errors.New("payload type is not registered")

	payloadTypeRegistry = make(map[string]reflect.Type)
)

type Record interface {
	Domain() string

	Codec() []byte

	MarshalRecord() ([]byte, error)

	UnmarshalRecord([]byte) error
}

func RegisterType(prototype Record) { _ = "STUB: not implemented"; return }

func unmarshalRecordPayload(payloadType []byte, payloadBytes []byte) (_rec Record, err error) {
	_ = "STUB: not implemented"
	return *new(Record), nil
}

func blankRecordForPayloadType(payloadType []byte) (Record, error) {
	_ = "STUB: not implemented"
	return *new(Record), nil
}

func getValueType(i any) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
