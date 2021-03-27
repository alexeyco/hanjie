// Package validator contains default puzzle validator.
package validator

import (
	"fmt"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
)

// Validator of the puzzle.
type Validator struct {
	rules []rule
}

// Validate the puzzle.
func (v *Validator) Validate(puzzleSet ast.PuzzleSet) error {
	var validationError errors.ValidationError
	for n, puzzle := range puzzleSet {
		for _, r := range v.rules {
			if err := r(puzzle); err != nil {
				validationError.Append(fmt.Errorf(`#%d: %w`, n, err))
			}
		}
	}

	if len(validationError) == 0 {
		return nil
	}

	return validationError
}

// New returns new validator instance.
func New() *Validator {
	return &Validator{
		rules: []rule{
			authorNameShouldNotBeEmpty,
			titleShouldNotBeEmpty,
			backgroundShouldBeCorrect,
			colorsShouldBeUnique,
			goalShouldHaveRows,
			goalLinesShouldHaveTheSameLength,
			//goalShouldBeCorrect,
		},
	}
}
