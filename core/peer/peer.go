package peer

import (
	"errors"

	"github.com/ipfs/go-cid"
	ic "github.com/libp2p/go-libp2p/core/crypto"
)

var (
	ErrEmptyPeerID = errors.New("empty peer ID")

	ErrNoPublicKey = errors.New("public key is not embedded in peer ID")
)

var AdvancedEnableInlining = true

const maxInlineKeyLength = 42

type ID string

func (id ID) Loggable() map[string]any { _ = "STUB: not implemented"; return nil }

func (id ID) String() string { _ = "STUB: not implemented"; return "" }

func (id ID) ShortString() string { _ = "STUB: not implemented"; return "" }

func (id ID) MatchesPrivateKey(sk ic.PrivKey) bool { _ = "STUB: not implemented"; return false }

func (id ID) MatchesPublicKey(pk ic.PubKey) bool { _ = "STUB: not implemented"; return false }

func (id ID) ExtractPublicKey() (ic.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(ic.PubKey), nil
}

func (id ID) Validate() error { _ = "STUB: not implemented"; return nil }

func IDFromBytes(b []byte) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func Decode(s string) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func FromCid(c cid.Cid) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func ToCid(id ID) cid.Cid { _ = "STUB: not implemented"; return *new(cid.Cid) }

func IDFromPublicKey(pk ic.PubKey) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

func IDFromPrivateKey(sk ic.PrivKey) (ID, error) { _ = "STUB: not implemented"; return *new(ID), nil }

type IDSlice []ID

func (es IDSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (es IDSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (es IDSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (es IDSlice) String() string { _ = "STUB: not implemented"; return "" }
