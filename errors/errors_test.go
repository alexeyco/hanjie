package errors_test

import (
	goerrors "errors"
	"testing"

	"github.com/alexeyco/hanjie/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidationError_Error(t *testing.T) {
	t.Parallel()

	err := errors.ValidationError{}

	assert.Equal(t, err.Error(), "validation error")
}

func TestValidationError_Append(t *testing.T) {
	t.Parallel()

	expected := errors.ValidationError{
		goerrors.New("foo"),
		goerrors.New("bar"),
	}

	actual := errors.ValidationError{}

	actual.Append(goerrors.New("foo"))
	actual.Append(goerrors.New("bar"))

	assert.Len(t, actual, 2)
	assert.Equal(t, expected, actual)
}
