package hanjie_test

import (
	"testing"

	"github.com/alexeyco/hanjie"
	"github.com/alexeyco/hanjie/ast"
	"github.com/stretchr/testify/assert"
)

type validator struct{}

func (v *validator) Validate(_ ast.PuzzleSet) error {
	return nil
}

func TestWithValidator(t *testing.T) {
	t.Parallel()

	v := &validator{}
	o := hanjie.Options{}

	hanjie.WithValidator(v)(&o)

	assert.Equal(t, v, o.Validator)
}

func TestSkipValidation(t *testing.T) {
	t.Parallel()

	o := hanjie.Options{}
	hanjie.SkipValidation(&o)

	assert.True(t, o.SkipValidation)
}
