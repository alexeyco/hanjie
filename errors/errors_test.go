package errors_test

import (
	goerrors "errors"
	"testing"

	"github.com/alexeyco/hanjie/errors"
)

func TestValidationError_Error(t *testing.T) {
	t.Parallel()

	err := errors.ValidationError{}

	if err.Error() != "validation error" {
		t.Errorf(`Should be "validation error", "%s" given`, err.Error())
	}
}

func TestValidationError_Append(t *testing.T) {
	t.Parallel()

	err := errors.ValidationError{}

	err.Append(goerrors.New("foo"))
	err.Append(goerrors.New("bar"))

	if len(err) != 2 {
		t.Fatalf(`Should be 2, %d given`, len(err))
	}

	if err[0].Error() != "foo" {
		t.Errorf(`Should be "foo", "%v" given`, err[0])
	}

	if err[1].Error() != "bar" {
		t.Errorf(`Should be "bar", "%v" given`, err[1])
	}
}
