package hanjie_test

import (
	"reflect"
	"testing"

	"github.com/alexeyco/hanjie"
	"github.com/alexeyco/hanjie/ast"
)

type validator struct{}

func (v *validator) Validate(puzzle ast.Puzzle) error {
	return nil
}

func TestWithValidator(t *testing.T) {
	t.Parallel()

	v := &validator{}
	o := hanjie.Options{}

	hanjie.WithValidator(v)(&o)

	if !reflect.DeepEqual(v, o.Validator) {
		t.Error(`Should be equal`)
	}
}

func TestSkipValidation(t *testing.T) {
	t.Parallel()

	o := hanjie.Options{}
	hanjie.SkipValidation(&o)

	if o.SkipValidation == false {
		t.Error(`Should be true`)
	}
}
