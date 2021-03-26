package pbn_test

import (
	"testing"

	"github.com/alexeyco/pbn"
)

func TestSkipValidation(t *testing.T) {
	t.Parallel()

	o := pbn.Options{}
	pbn.SkipValidation(&o)

	if o.SkipValidation == false {
		t.Error(`Should be true`)
	}
}
