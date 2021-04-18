package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alexeyco/hanjie/ast"
	"github.com/alexeyco/hanjie/errors"
	"github.com/alexeyco/hanjie/tools"
)

type rule func(puzzle ast.Puzzle) (err error, stop bool)

func authorNameRule(puzzle ast.Puzzle) (err error, stop bool) {
	if puzzle.Author != nil && puzzle.Author.Name == "" {
		err = errors.ErrEmptyAuthorName
	}

	return
}

func titleRule(puzzle ast.Puzzle) (err error, stop bool) {
	if puzzle.Title == "" {
		err = errors.ErrEmptyTitle
	}

	return
}

func backgroundRule(puzzle ast.Puzzle) (error, bool) {
	var symbols []string
	for ch := range puzzle.Colors {
		if ch == puzzle.Background {
			return nil, false
		}

		symbols = append(symbols, string(ch))
	}

	return fmt.Errorf(`%w, should be one of ["%s"]`,
		errors.ErrIncorrectBackground,
		strings.Join(symbols, `", "`)), true
}

func uniqueColorRule(puzzle ast.Puzzle) (err error, stop bool) {
	used := map[string]bool{}
	for _, color := range puzzle.Colors {
		id := fmt.Sprintf("%d-%d-%d", color.R, color.G, color.B)
		if _, ok := used[id]; ok {
			return fmt.Errorf(`%w #%02x%02x%02x {R: %d, G: %d, B: %d}`,
				errors.ErrColorHasAlreadyBeenUsed,
				color.R, color.G, color.B,
				color.R, color.G, color.B), false
		}

		used[id] = true
	}

	return
}

func goalRowsRule(puzzle ast.Puzzle) (err error, stop bool) {
	if puzzle.Goal != nil && len(*puzzle.Goal) == 0 {
		err = errors.ErrGoalIsIncorrect
		stop = true
	}

	return
}

func goalLinesRule(puzzle ast.Puzzle) (error, bool) {
	if puzzle.Goal == nil {
		return nil, false
	}

	var expectedLen int
	for _, row := range *puzzle.Goal {
		rowLen := len(row)
		if rowLen == 0 {
			return errors.ErrGoalIsIncorrect, true
		}

		if expectedLen == 0 {
			expectedLen = rowLen
			continue
		}

		if expectedLen != rowLen {
			return errors.ErrGoalIsIncorrect, true
		}
	}

	return nil, false
}

func goalClueMatchRule(puzzle ast.Puzzle) (err error, stop bool) {
	if puzzle.Goal == nil {
		return
	}

	clue := tools.GoalToClue(*puzzle.Goal, puzzle.Background)
	if !reflect.DeepEqual(clue, puzzle.Clue) {
		return errors.ErrGoalDoesNotMatchTheClue, true
	}

	return
}
