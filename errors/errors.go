// Package errors contains errors.
package errors

import "errors"

var (
	// ErrSyntax syntax error.
	ErrSyntax = errors.New("syntax problems")

	// ErrEmptyAuthorName validation error, reports that puzzle author name shouldn't be empty.
	ErrEmptyAuthorName = errors.New("author name shouldn't be empty")
	// ErrEmptyTitle validation error, reports that puzzle title shouldn't be empty.
	ErrEmptyTitle = errors.New("title shouldn't be empty")
	// ErrIncorrectBackground validation error, reports that puzzle title shouldn't be empty.
	ErrIncorrectBackground = errors.New("incorrect background")
	// ErrColorHasAlreadyBeenUsed validation error, reports if color has already been used.
	ErrColorHasAlreadyBeenUsed = errors.New("color has already been used")
	// ErrGoalIsIncorrect validation error, reports if goal is incorrect.
	ErrGoalIsIncorrect = errors.New("goal is incorrect")
)

// ValidationError puzzle validation errors batch.
type ValidationError []error

// Error returns validation error message.
func (e ValidationError) Error() string {
	return "validation error"
}

// Merge errors.
func (e *ValidationError) Append(err error) {
	*e = append(*e, err)
}
