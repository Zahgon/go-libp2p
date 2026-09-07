package record

import (
	"errors"
	"sync"

	"github.com/libp2p/go-libp2p/core/crypto"
)

type Envelope struct {
	PublicKey crypto.PubKey

	PayloadType []byte

	RawPayload []byte

	signature []byte

	cached         Record
	unmarshalError error
	unmarshalOnce  sync.Once
}

var ErrEmptyDomain = errors.New("envelope domain must not be empty")
var ErrEmptyPayloadType = errors.New("payloadType must not be empty")
var ErrInvalidSignature = errors.New("invalid signature or incorrect domain")

func Seal(rec Record, privateKey crypto.PrivKey) (*Envelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConsumeEnvelope(data []byte, domain string) (envelope *Envelope, rec Record, err error) {
	_ = "STUB: not implemented"
	return nil, *new(Record), nil
}

func ConsumeTypedEnvelope(data []byte, destRecord Record) (envelope *Envelope, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalEnvelope(data []byte) (*Envelope, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Envelope) Marshal() (res []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Envelope) Equal(other *Envelope) bool { _ = "STUB: not implemented"; return false }

func (e *Envelope) Record() (Record, error) { _ = "STUB: not implemented"; return *new(Record), nil }

func (e *Envelope) TypedRecord(dest Record) error { _ = "STUB: not implemented"; return nil }

func (e *Envelope) validate(domain string) error { _ = "STUB: not implemented"; return nil }

func makeUnsigned(domain string, payloadType []byte, payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
