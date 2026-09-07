package peer

type addrInfoJson struct {
	ID    ID
	Addrs []string
}

func (pi AddrInfo) MarshalJSON() (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pi *AddrInfo) UnmarshalJSON(b []byte) (err error) { _ = "STUB: not implemented"; return nil }
