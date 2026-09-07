package protocol

type ID string

const (
	TestingID ID = "/p2p/_testing"
)

func ConvertFromStrings(ids []string) (res []ID) { _ = "STUB: not implemented"; return nil }

func ConvertToStrings(ids []ID) (res []string) { _ = "STUB: not implemented"; return nil }
