package pbn_test

import (
	"reflect"
	"testing"

	"github.com/alexeyco/pbn"
	"github.com/alexeyco/pbn/ast"
)

type validator struct{}

func (v *validator) Validate(puzzle ast.Puzzle) error {
	return nil
}

func TestWithValidator(t *testing.T) {
	t.Parallel()

	v := &validator{}
	o := pbn.Options{}

	pbn.WithValidator(v)(&o)

	if !reflect.DeepEqual(v, o.Validator) {
		t.Error(`Should be equal`)
	}
}

func TestSkipValidation(t *testing.T) {
	t.Parallel()

	o := pbn.Options{}
	pbn.SkipValidation(&o)

	if o.SkipValidation == false {
		t.Error(`Should be true`)
	}
}
