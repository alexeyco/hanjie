// Package validator contains default puzzle validator.
package validator

import "github.com/alexeyco/hanjie/ast"

// Validator of the puzzle.
type Validator struct{}

// Validate the puzzle.
func (v Validator) Validate(puzzle ast.Puzzle) error {
	panic("implement me")
}

// New returns new validator instance.
func New() *Validator {
	return &Validator{}
}
