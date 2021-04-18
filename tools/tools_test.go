package tools_test

import (
	"testing"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/tools"
	"github.com/stretchr/testify/assert"
)

func TestGoalToClue(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name       string
		goal       ast.Goal
		background ast.Char
		expected   ast.Clue
	}{
		{
			name: "SingleColored",
			goal: ast.Goal{
				{ast.Char('x'), ast.Char('.'), ast.Char('x'), ast.Char('.')},
				{ast.Char('x'), ast.Char('x'), ast.Char('.'), ast.Char('x')},
				{ast.Char('x'), ast.Char('.'), ast.Char('x'), ast.Char('x')},
			},
			background: ast.Char('.'),
			expected: ast.Clue{
				Columns: []ast.Line{
					{
						{Color: ast.Char('x'), Count: 3},
					},
					{
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 1},
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 2},
					},
				},
				Rows: []ast.Line{
					{
						{Color: ast.Char('x'), Count: 1},
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 2},
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 1},
						{Color: ast.Char('x'), Count: 2},
					},
				},
			},
		},
		{
			name: "MultiColored",
			goal: ast.Goal{
				{ast.Char('x'), ast.Char('.'), ast.Char('y'), ast.Char('x')},
				{ast.Char('x'), ast.Char('y'), ast.Char('y'), ast.Char('.')},
				{ast.Char('y'), ast.Char('x'), ast.Char('x'), ast.Char('x')},
			},
			background: ast.Char('.'),
			expected: ast.Clue{
				Columns: []ast.Line{
					{
						{Color: ast.Char('x'), Count: 2},
						{Color: ast.Char('y'), Count: 1},
					},
					{
						{Color: ast.Char('y'), Count: 1},
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('y'), Count: 2},
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 1},
						{Color: ast.Char('x'), Count: 1},
					},
				},
				Rows: []ast.Line{
					{
						{Color: ast.Char('x'), Count: 1},
						{Color: ast.Char('y'), Count: 1},
						{Color: ast.Char('x'), Count: 1},
					},
					{
						{Color: ast.Char('x'), Count: 1},
						{Color: ast.Char('y'), Count: 2},
					},
					{
						{Color: ast.Char('y'), Count: 1},
						{Color: ast.Char('x'), Count: 3},
					},
				},
			},
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			actual := tools.GoalToClue(testDatum.goal, testDatum.background)
			assert.Equal(t, testDatum.expected, actual)
		})
	}
}

func TestTransposeGoal(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name     string
		goal     ast.Goal
		expected ast.Goal
	}{
		{
			name: "Ok",
			goal: ast.Goal{
				{ast.Char('a'), ast.Char('b'), ast.Char('c')},
				{ast.Char('d'), ast.Char('e'), ast.Char('f')},
				{ast.Char('g'), ast.Char('h'), ast.Char('i')},
			},
			expected: ast.Goal{
				{ast.Char('a'), ast.Char('d'), ast.Char('g')},
				{ast.Char('b'), ast.Char('e'), ast.Char('h')},
				{ast.Char('c'), ast.Char('f'), ast.Char('i')},
			},
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			actual := tools.TransposeGoal(testDatum.goal)
			assert.Equal(t, testDatum.expected, actual)
		})
	}
}
