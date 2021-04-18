package hanjie_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/alexeyco/hanjie"
	"github.com/alexeyco/hanjie/ast"
	"github.com/stretchr/testify/assert"
)

var expectedPuzzleSet = &ast.PuzzleSet{
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

var expectedString = `- id: id
  source: https://foo.bar
  author:
    name: John Doe
    id: johnDoe
  copyright: '&copy; John Doe'
  title: Puzzle
  description: Very beautiful puzzle
  background: .
  colors:
    .: '#ffffff'
    x: '#000000'
  clue:
    columns: [[{color: x, count: 2}], [{color: x, count: 1}, {color: x, count: 1}], [{color: x, count: 2}]]
    rows: [[{color: x, count: 2}], [{color: x, count: 1}, {color: x, count: 1}], [{color: x, count: 2}]]
  goal: [[x, x, .], [x, ., x], [., x, x]]
`

func TestRead(t *testing.T) {
	t.Parallel()

	actual, err := hanjie.Read(strings.NewReader(expectedString))

	assert.NoError(t, err)
	assert.Equal(t, expectedPuzzleSet, actual)
}

func TestWrite(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	err := hanjie.Write(&buf, expectedPuzzleSet)

	assert.NoError(t, err)
	assert.Equal(t, expectedString, buf.String())
}
