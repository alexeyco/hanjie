// Package ast declares puzzle types.
package ast

import (
	"fmt"
	"unicode/utf8"

	"github.com/alexeyco/hanjie/errors"
)

// PuzzleSet the root structure of the document. Includes all puzzles.
type PuzzleSet []Puzzle

// Puzzle a puzzle in the set of puzzles.
type Puzzle struct {
	ID          string  `yaml:"id,omitempty"`
	Source      string  `yaml:"source,omitempty"`
	Author      *Author `yaml:"author,omitempty"`
	Copyright   string  `yaml:"copyright,omitempty"`
	Title       string  `yaml:"title"`
	Description string  `yaml:"description,omitempty"`
	Background  Char    `yaml:"background"`
	Colors      Colors  `yaml:"colors"`
	Clue        Clue    `yaml:"clue"`
	Goal        *Goal   `yaml:"goal,flow,omitempty"`
}

// Author of puzzle.
type Author struct {
	Name string `yaml:"name,omitempty"`
	ID   string `yaml:"id,omitempty"`
}

// Colors used in puzzle.
type Colors map[Char]Color

// Char uniquely identifies a color.
type Char rune

// MarshalText encodes the character into UTF-8-encoded text and returns the result.
func (c Char) MarshalText() ([]byte, error) {
	return []byte(string(c)), nil
}

// UnmarshalText decodes the character from UTF-8-encoded text.
func (c *Char) UnmarshalText(b []byte) error {
	cnt := utf8.RuneCount(b)
	if cnt == 0 {
		return fmt.Errorf(`%w empty string: should be a single char`, errors.ErrUnmarshal)
	}

	if cnt > 1 {
		return fmt.Errorf(`%w "%s": should be a single char`, errors.ErrUnmarshal, string(b))
	}

	r, _ := utf8.DecodeRune(b)
	*c = Char(r)

	return nil
}

// Color in RGB.
type Color struct {
	R, G, B uint8
}

// MarshalText encodes the color into UTF-8-encoded hex string.
func (c Color) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)), nil
}

// UnmarshalText decodes the color from UTF-8-encoded hex string.
func (c *Color) UnmarshalText(b []byte) (err error) {
	s := string(b)
	if s[0] != '#' {
		return errors.ErrUnmarshal
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}

		err = errors.ErrUnmarshal

		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errors.ErrUnmarshal
	}

	return
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values for the color.
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8

	g = uint32(c.G)
	g |= g << 8

	b = uint32(c.B)
	b |= b << 8

	a = uint32(0xff)
	a |= a << 8

	return
}

// Clue defines a clue used in the puzzle.
type Clue struct {
	Columns []Line `yaml:"columns"`
	Rows    []Line `yaml:"rows"`
}

// Line of the clue.
type Line struct {
	Color Char `yaml:"color"`
	Count int  `yaml:"count"`
}

// Goal of the puzzle.
type Goal [][]Char
