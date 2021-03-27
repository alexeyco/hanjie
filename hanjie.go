// Package hanjie contains any entities and instruments to read, write and validate puzzles.
package hanjie

import (
	"io"

	"gopkg.in/yaml.v3"

	"github.com/alexeyco/hanjie/ast"
)

// Validator interface.
type Validator interface {
	Validate(ast.PuzzleSet) error
}

// Read set of puzzles from io.Reader.
func Read(r io.Reader, options ...Option) (*ast.PuzzleSet, error) {
	var puzzleSet ast.PuzzleSet
	if err := yaml.NewDecoder(r).Decode(&puzzleSet); err != nil {
		return nil, err
	}

	o := newOptions()
	for _, opt := range options {
		opt(&o)
	}

	return &puzzleSet, nil
}

// Write set of puzzles to io.Writer.
func Write(w io.Writer, puzzleSet *ast.PuzzleSet, options ...Option) error {
	o := newOptions()
	for _, opt := range options {
		opt(&o)
	}

	return yaml.NewEncoder(w).Encode(puzzleSet)
}
