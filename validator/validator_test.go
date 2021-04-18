package validator_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
	"github.com/alexeyco/hanjie/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	v := validator.New()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()

		err := v.Validate(puzzleSet)

		assert.NoError(t, err)
	})

	t.Run("OkWithoutGoal", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = nil

		err := v.Validate(puzzleSet)

		assert.NoError(t, err)
	})

	t.Run("OkWithoutAuthor", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Author = nil

		err := v.Validate(puzzleSet)

		assert.NoError(t, err)
	})

	t.Run("ErrorCauseEmptyAuthorName", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Author.Name = ""

		expected := errors.ValidationError{
			errors.ErrEmptyAuthorName,
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseEmptyTitle", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Title = ""

		expected := errors.ValidationError{
			errors.ErrEmptyTitle,
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseBackgroundIsIncorrect", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Background = ast.Char('-')

		expected := errors.ValidationError{
			fmt.Errorf(`%w, should be one of ["%s"]`, errors.ErrIncorrectBackground, strings.Join([]string{".", "x"}, `", "`)),
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseColorHasAlreadyBeenUsed", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Colors = ast.Colors{
			ast.Char('.'): ast.Color{R: 255, G: 255, B: 255},
			ast.Char('x'): ast.Color{},
			ast.Char('y'): ast.Color{},
		}

		expected := errors.ValidationError{
			fmt.Errorf(`%w #%02x%02x%02x {R: %d, G: %d, B: %d}`, errors.ErrColorHasAlreadyBeenUsed, 0, 0, 0, 0, 0, 0),
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseGoalHaveNoLines", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{}

		expected := errors.ValidationError{
			errors.ErrGoalIsIncorrect,
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseGoalHaveEmptyLine", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{
			{},
		}

		expected := errors.ValidationError{
			errors.ErrGoalIsIncorrect,
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseGoalLinesHaveDifferentLength", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{
			{ast.Char('x'), ast.Char('x')},
			{ast.Char('x'), ast.Char('.'), ast.Char('x')},
			{ast.Char('.'), ast.Char('x'), ast.Char('x')},
		}

		expected := errors.ValidationError{
			errors.ErrGoalIsIncorrect,
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorCauseGoalDoesNotMatchTheClue", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{
			{ast.Char('x'), ast.Char('x'), ast.Char('x')},
			{ast.Char('x'), ast.Char('x'), ast.Char('x')},
			{ast.Char('x'), ast.Char('x'), ast.Char('x')},
		}

		expected := errors.ValidationError{
			errors.ErrGoalDoesNotMatchTheClue,
		}

		actual := v.Validate(puzzleSet)

		assert.Error(t, actual)
		assert.Equal(t, expected, actual)
	})
}

func newPuzzleSet() ast.PuzzleSet {
	return ast.PuzzleSet{
		{
			ID:     "id",
			Source: "https://foo.bar",
			Author: &ast.Author{
				Name: "John Doe",
				ID:   "johnDoe",
			},
			Copyright:   "&copy; John Doe",
			Title:       "Puzzle",
			Description: "Very beautiful puzzle",
			Background:  ast.Char('.'),
			Colors: ast.Colors{
				ast.Char('.'): ast.Color{R: 255, G: 255, B: 255},
				ast.Char('x'): ast.Color{},
			},
			Clue: ast.Clue{
				Columns: []ast.Line{
					{
						{Color: ast.Char('x'), Count: 2},
					},
					{
						{Color: ast.Char('x'), Count: 1}, {Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 2},
					},
				},
				Rows: []ast.Line{
					{
						{Color: ast.Char('x'), Count: 2},
					},
					{
						{Color: ast.Char('x'), Count: 1}, {Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 2},
					},
				},
			},
			Goal: &ast.Goal{
				{ast.Char('x'), ast.Char('x'), ast.Char('.')},
				{ast.Char('x'), ast.Char('.'), ast.Char('x')},
				{ast.Char('.'), ast.Char('x'), ast.Char('x')},
			},
		},
	}
}
