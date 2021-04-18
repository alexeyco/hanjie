package ast_test

import (
	"fmt"
	"testing"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
	"github.com/stretchr/testify/assert"
)

func TestChar_MarshalText(t *testing.T) {
	t.Parallel()

	ch := ast.Char('x')
	actual, err := ch.MarshalText()

	assert.NoError(t, err)
	assert.Equal(t, []byte("x"), actual)
}

func TestChar_UnmarshalText(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte("x"))

		assert.NoError(t, err)
		assert.Equal(t, ast.Char('x'), ch)
	})

	t.Run("ErrorCauseEmpty", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte{})

		assert.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrSyntax)
		assert.Equal(t, int32(0), int32(ch))
	})

	t.Run("ErrorCauseTooManyCharacters", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte("xx"))

		assert.Error(t, err)
		assert.ErrorIs(t, err, errors.ErrSyntax)
		assert.Equal(t, int32(0), int32(ch))
	})
}

func TestColor_MarshalText(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		color    ast.Color
		expected string
	}{
		{
			color:    ast.Color{},
			expected: "#000000",
		},
		{
			color:    ast.Color{R: 255},
			expected: "#ff0000",
		},
		{
			color:    ast.Color{G: 255},
			expected: "#00ff00",
		},
		{
			color:    ast.Color{B: 255},
			expected: "#0000ff",
		},
		{
			color:    ast.Color{R: 255, G: 255, B: 255},
			expected: "#ffffff",
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		name := fmt.Sprintf("%d %d %d", testDatum.color.R, testDatum.color.G, testDatum.color.R)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			b, err := testDatum.color.MarshalText()

			assert.NoError(t, err)
			assert.Equal(t, testDatum.expected, string(b))
		})
	}
}

func TestColor_UnmarshalText(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		b        []byte
		expected *ast.Color
		err      error
	}{
		{
			b:        []byte("#000"),
			expected: &ast.Color{},
		},
		{
			b:        []byte("#000000"),
			expected: &ast.Color{},
		},
		{
			b:        []byte("#fff"),
			expected: &ast.Color{R: 255, G: 255, B: 255},
		},
		{
			b:        []byte("#FFF"),
			expected: &ast.Color{R: 255, G: 255, B: 255},
		},
		{
			b:        []byte("#ffffff"),
			expected: &ast.Color{R: 255, G: 255, B: 255},
		},
		{
			b:        []byte("#FFFFFF"),
			expected: &ast.Color{R: 255, G: 255, B: 255},
		},
		{
			b:        []byte("#zzz"),
			expected: &ast.Color{},
			err:      errors.ErrSyntax,
		},
		{
			b:        []byte("#z"),
			expected: &ast.Color{},
			err:      errors.ErrSyntax,
		},
		{
			b:        []byte("wrong"),
			expected: &ast.Color{},
			err:      errors.ErrSyntax,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(string(testDatum.b), func(t *testing.T) {
			t.Parallel()

			actual := &ast.Color{}
			err := actual.UnmarshalText(testDatum.b)

			assert.Equal(t, testDatum.expected, actual)
			if testDatum.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, testDatum.err)
			}
		})
	}
}

func TestColor_RGBA(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		color      ast.Color
		r, g, b, a uint32
	}{
		{
			color: ast.Color{},
			a:     65535,
		},
		{
			color: ast.Color{R: 255, G: 255, B: 255},
			r:     65535,
			g:     65535,
			b:     65535,
			a:     65535,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		name := fmt.Sprintf("%d %d %d", testDatum.color.R, testDatum.color.G, testDatum.color.R)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r, g, b, a := testDatum.color.RGBA()

			assert.Equal(t, testDatum.r, r)
			assert.Equal(t, testDatum.g, g)
			assert.Equal(t, testDatum.b, b)
			assert.Equal(t, testDatum.a, a)
		})
	}
}
