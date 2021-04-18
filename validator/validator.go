// Package validator contains default puzzle validator.
package validator

import (
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
	for _, puzzle := range puzzleSet {
		for _, r := range v.rules {
			err, stop := r(puzzle)
			if err != nil {
				validationError.Append(err)
			}

			if stop {
				break
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
			authorNameRule,
			titleRule,
			backgroundRule,
			uniqueColorRule,
			goalRowsRule,
			goalLinesRule,
			goalClueMatchRule,
		},
	}
}
