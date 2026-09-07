package peer

import (
	"encoding"
	"encoding/json"
)

var _ json.Marshaler = (*ID)(nil)
var _ json.Unmarshaler = (*ID)(nil)

var _ encoding.BinaryMarshaler = (*ID)(nil)
var _ encoding.BinaryUnmarshaler = (*ID)(nil)
var _ encoding.TextMarshaler = (*ID)(nil)
var _ encoding.TextUnmarshaler = (*ID)(nil)

func (id ID) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id ID) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id ID) MarshalTo(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (id *ID) Unmarshal(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (id *ID) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (id ID) Size() int { _ = "STUB: not implemented"; return 0 }

func (id ID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *ID) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (id ID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *ID) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }
