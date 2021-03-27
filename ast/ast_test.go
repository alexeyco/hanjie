package ast_test

import (
	goerrors "errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
)

func TestChar_MarshalText(t *testing.T) {
	t.Parallel()

	ch := ast.Char('x')
	actual, err := ch.MarshalText()

	if err != nil {
		t.Errorf(`Error should be nil, "%v" given`, err)
	}

	if string(actual) != "x" {
		t.Errorf(`Should be "x", "%s" given`, string(actual))
	}
}

func TestChar_UnmarshalText(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		if err := ch.UnmarshalText([]byte("x")); err != nil {
			t.Errorf(`Error should be nil, "%v" given`, err)
		}

		if rune(ch) != 'x' {
			t.Errorf(`Char should be 'x', '%s' given`, string(ch))
		}
	})

	t.Run("ErrorCauseEmpty", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte{})

		if err == nil {
			t.Error(`Error should not be nil`)
		}

		if !goerrors.Is(err, errors.ErrSyntax) {
			t.Errorf(`Error should be errors.ErrSyntax, "%v" given`, err)
		}

		if ch != 0 {
			t.Errorf(`Char should be blank, '%s' given`, string(ch))
		}
	})

	t.Run("ErrorCauseTooManyCharacters", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte("xx"))

		if err == nil {
			t.Error(`Error should not be nil`)
		}

		if !goerrors.Is(err, errors.ErrSyntax) {
			t.Errorf(`Error should be errors.ErrSyntax, "%v" given`, err)
		}

		if ch != 0 {
			t.Errorf(`Char should be blank, '%s' given`, string(ch))
		}
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
			if err != nil {
				t.Errorf(`Should be nil, "%v" given`, err)
			}

			actual := string(b)
			if actual != testDatum.expected {
				t.Errorf(`Should be "%s", "%s" given`, testDatum.expected, actual)
			}
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
			switch {
			case testDatum.err == nil && err != nil:
				t.Errorf(`Should be nil, "%v" given`, err)
			case testDatum.err != nil && err == nil:
				t.Errorf(`Should be "%v", nil given`, testDatum.err)
			case testDatum.err != nil && err != nil && !goerrors.Is(err, testDatum.err):
				t.Errorf(`Should be "%v", "%s" given`, testDatum.err, err)
			}

			if !reflect.DeepEqual(testDatum.expected, actual) {
				t.Errorf(`Should be {R: %d, G: %d, B: %d}, {R: %d, G: %d, B: %d} given`,
					testDatum.expected.R,
					testDatum.expected.G,
					testDatum.expected.B,
					actual.R,
					actual.G,
					actual.B)
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
			if r != testDatum.r || g != testDatum.g || b != testDatum.b || a != testDatum.a {
				t.Errorf(`Should be (%d, %d, %d, %d), (%d, %d, %d, %d) given`,
					testDatum.r,
					testDatum.g,
					testDatum.b,
					testDatum.a,
					r, g, b, a)
			}
		})
	}
}
