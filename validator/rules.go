package validator

import (
	"fmt"
	"strings"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
)

type rule func(puzzle ast.Puzzle) error

func authorNameShouldNotBeEmpty(puzzle ast.Puzzle) (err error) {
	if puzzle.Author != nil && puzzle.Author.Name == "" {
		err = errors.ErrEmptyAuthorName
	}

	return
}

func titleShouldNotBeEmpty(puzzle ast.Puzzle) (err error) {
	if puzzle.Title == "" {
		err = errors.ErrEmptyTitle
	}

	return
}

func backgroundShouldBeCorrect(puzzle ast.Puzzle) error {
	var symbols []string
	for ch := range puzzle.Colors {
		if ch == puzzle.Background {
			return nil
		}

		symbols = append(symbols, string(ch))
	}

	return fmt.Errorf(`%w, should be one of ["%s"]`,
		errors.ErrIncorrectBackground,
		strings.Join(symbols, `", "`))
}

func colorsShouldBeUnique(puzzle ast.Puzzle) (err error) {
	used := map[string]bool{}
	for _, color := range puzzle.Colors {
		id := fmt.Sprintf("%d-%d-%d", color.R, color.G, color.B)
		if _, ok := used[id]; ok {
			return fmt.Errorf(`%w #%02x%02x%02x {R: %d, G: %d, B: %d}`,
				errors.ErrColorHasAlreadyBeenUsed,
				color.R, color.G, color.B,
				color.R, color.G, color.B)
		}

		used[id] = true
	}

	return
}

func goalShouldHaveRows(puzzle ast.Puzzle) (err error) {
	if puzzle.Goal != nil && len(*puzzle.Goal) == 0 {
		err = errors.ErrGoalIsIncorrect
	}

	return
}

func goalLinesShouldHaveTheSameLength(puzzle ast.Puzzle) error {
	if puzzle.Goal == nil {
		return nil
	}

	var expectedLen int
	for _, row := range *puzzle.Goal {
		rowLen := len(row)
		if rowLen == 0 {
			return errors.ErrGoalIsIncorrect
		}

		if expectedLen == 0 {
			expectedLen = rowLen
			continue
		}

		if expectedLen != rowLen {
			return errors.ErrGoalIsIncorrect
		}
	}

	return nil
}

//func goalShouldBeCorrect(puzzle ast.Puzzle) (err error) {
//	if puzzle.Goal == nil {
//		return
//	}
//
//	var (
//		columns [][]ast.Line
//		rows    [][]ast.Line
//	)
//
//	for r, row := range *puzzle.Goal {
//		for c, ch := range row {
//
//		}
//	}
//
//	if !reflect.DeepEqual(columns, puzzle.Clue.Columns) || !reflect.DeepEqual(rows, puzzle.Clue.Rows) {
//		err = errors.ErrGoalIsIncorrect
//	}
//
//	return
//}
