package pbn

type Options struct {
	SkipValidation bool
}

type Option func(*Options)

func SkipValidation(o *Options) {
	o.SkipValidation = true
}
