package routing

type Option func(opts *Options) error

type Options struct {
	Expired bool
	Offline bool

	Other map[any]any
}

func (opts *Options) Apply(options ...Option) error { _ = "STUB: not implemented"; return nil }

func (opts *Options) ToOption() Option { _ = "STUB: not implemented"; return *new(Option) }

var Expired Option = func(opts *Options) error {
	opts.Expired = true
	return nil
}

var Offline Option = func(opts *Options) error {
	opts.Offline = true
	return nil
}
