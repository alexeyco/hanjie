package validator_test

import (
	goerrors "errors"
	"testing"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
	"github.com/alexeyco/hanjie/validator"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	v := validator.New()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		if err := v.Validate(puzzleSet); err != nil {
			t.Errorf(`Should be nil`)
		}
	})

	t.Run("OkWithoutGoal", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = nil

		if err := v.Validate(puzzleSet); err != nil {
			t.Errorf(`Should be nil`)
		}
	})

	t.Run("OkWithoutAuthor", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Author = nil

		if err := v.Validate(puzzleSet); err != nil {
			t.Errorf(`Should be nil`)
		}
	})

	t.Run("ErrorCauseEmptyAuthorName", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Author.Name = ""

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrEmptyAuthorName) {
			t.Errorf(`Should be errors.ErrEmptyAuthorName, "%v" given`, validationError[0])
		}
	})

	t.Run("ErrorCauseEmptyTitle", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Title = ""

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrEmptyTitle) {
			t.Errorf(`Should be errors.ErrEmptyTitle, "%v" given`, validationError[0])
		}
	})

	t.Run("ErrorCauseBackgroundIsIncorrect", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Background = ast.Char('-')

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrIncorrectBackground) {
			t.Errorf(`Should be errors.ErrIncorrectBackground, "%v" given`, validationError[0])
		}
	})

	t.Run("ErrorCauseColorHasAlreadyBeenUsed", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Colors = ast.Colors{
			ast.Char('.'): ast.Color{R: 255, G: 255, B: 255},
			ast.Char('x'): ast.Color{},
			ast.Char('y'): ast.Color{},
		}

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrColorHasAlreadyBeenUsed) {
			t.Errorf(`Should be errors.ErrColorHasAlreadyBeenUsed, "%v" given`, validationError[0])
		}
	})

	t.Run("ErrorCauseGoalHaveNoLines", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{}

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrGoalIsIncorrect) {
			t.Errorf(`Should be errors.ErrColorHasAlreadyBeenUsed, "%v" given`, validationError[0])
		}
	})

	t.Run("ErrorCauseGoalHaveEmptyLine", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{
			{},
		}

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrGoalIsIncorrect) {
			t.Errorf(`Should be errors.ErrColorHasAlreadyBeenUsed, "%v" given`, validationError[0])
		}
	})

	t.Run("ErrorCauseGoalLinesHaveDifferentLength", func(t *testing.T) {
		t.Parallel()

		puzzleSet := newPuzzleSet()
		puzzleSet[0].Goal = &ast.Goal{
			{ast.Char('x'), ast.Char('x')},
			{ast.Char('x'), ast.Char('.'), ast.Char('x')},
			{ast.Char('.'), ast.Char('x'), ast.Char('x')},
		}

		err := v.Validate(puzzleSet)
		if err == nil {
			t.Fatalf(`Should not be nil`)
		}

		validationError := err.(errors.ValidationError)
		if len(validationError) != 1 {
			t.Errorf(`Should be 1, %d given`, len(validationError))
		}

		if !goerrors.Is(validationError[0], errors.ErrGoalIsIncorrect) {
			t.Errorf(`Should be errors.ErrColorHasAlreadyBeenUsed, "%v" given`, validationError[0])
		}
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
				Columns: [][]ast.Line{
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
				Rows: [][]ast.Line{
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
