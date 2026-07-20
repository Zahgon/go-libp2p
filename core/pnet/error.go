package pnet

var ErrNotInPrivateNetwork = NewError("private network was not configured but" +
	" is enforced by the environment")

type Error interface {
	IsPNetError() bool
}

func NewError(err string) error { _ = "STUB: not implemented"; return nil }

func IsPNetError(err error) bool { _ = "STUB: not implemented"; return false }

type pnetErr string

var _ Error = (*pnetErr)(nil)

func (p pnetErr) Error() string { _ = "STUB: not implemented"; return "" }

func (pnetErr) IsPNetError() bool { _ = "STUB: not implemented"; return false }
