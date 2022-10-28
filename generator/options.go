package generator

type Options struct {
	Service string

	Directory string

	Namespace string
}

type Option func(o *Options)

func Service(s string) Option {
	return func(o *Options) {
		o.Service = s
	}
}

func Directory(s string) Option {
	return func(o *Options) {
		o.Directory = s
	}
}

func Namespace(s string) Option {
	return func(o *Options) {
		o.Namespace = s
	}
}
