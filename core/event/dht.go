package event

type RawJSON string

type GenericDHTEvent struct {
	Type string

	Raw RawJSON
}
