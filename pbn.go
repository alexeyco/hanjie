package pbn

import (
	"errors"
	"fmt"
	"io"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

// ErrUnmarshal unmarshalling error.
var ErrUnmarshal = errors.New("can't unmarshal")

type PuzzleSet struct {
	Puzzles []Puzzle `yaml:"puzzles"`
}

type Puzzle struct {
	ID              string  `yaml:"id,omitempty"`
	Source          string  `yaml:"source,omitempty"`
	Author          *Author `yaml:"author,omitempty"`
	Copyright       string  `yaml:"copyright,omitempty"`
	Title           string  `yaml:"title"`
	Description     string  `yaml:"description,omitempty"`
	BackgroundColor Char    `yaml:"background_color"`
	DefaultColor    Char    `yaml:"default_color"`
	Colors          Colors  `yaml:"colors"`
	Clue            Clue    `yaml:"clue"`
	Goal            *Goal   `yaml:"goal,flow,omitempty"`
}

type Author struct {
	Name string `yaml:"name,omitempty"`
	ID   string `yaml:"id,omitempty"`
}

type Colors []Color

type Color struct {
	Char Char   `yaml:"char"`
	Hex  string `yaml:"hex"`
}

type Char rune

// MarshalText encodes the character into UTF-8-encoded text and returns the result.
func (c Char) MarshalText() ([]byte, error) {
	return []byte(string(c)), nil
}

// UnmarshalText decodes the character from UTF-8-encoded text.
func (c *Char) UnmarshalText(b []byte) error {
	cnt := utf8.RuneCount(b)
	if cnt == 0 {
		return fmt.Errorf(`%w empty string: should be a single char`, ErrUnmarshal)
	}

	if cnt > 1 {
		return fmt.Errorf(`%w "%s": should be a single char`, ErrUnmarshal, string(b))
	}

	r, _ := utf8.DecodeRune(b)
	*c = Char(r)

	return nil
}

type Clue struct {
	Columns []Line `yaml:"columns"`
	Rows    []Line `yaml:"rows"`
}

type Line struct {
	Color *Char `yaml:"color,omitempty"`
	Count int   `yaml:"count"`
}

type Goal [][]Char

func Read(r io.Reader) (*PuzzleSet, error) {
	var puzzleSet PuzzleSet
	if err := yaml.NewDecoder(r).Decode(&puzzleSet); err != nil {
		return nil, err
	}

	return &puzzleSet, nil
}

func Write(w io.Writer, puzzleSet *PuzzleSet) error {
	return yaml.NewEncoder(w).Encode(puzzleSet)
}
