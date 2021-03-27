package hanjie

import "github.com/alexeyco/hanjie/validator"

// Options defines puzzle read- and write options.
type Options struct {
	Validator      Validator
	SkipValidation bool
}

// Option setter.
type Option func(*Options)

// WithValidator to set a custom validator.
func WithValidator(v Validator) Option {
	return func(o *Options) {
		o.Validator = v
	}
}

// SkipValidation skips validation.
func SkipValidation(o *Options) {
	o.SkipValidation = true
}

func newOptions() Options {
	return Options{
		Validator: validator.New(),
	}
}
