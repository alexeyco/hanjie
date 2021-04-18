// Package tools contains hanjie tools.
package tools

import "github.com/alexeyco/hanjie/ast"

// GoalToClue returns clue by goal.
func GoalToClue(goal ast.Goal, background ast.Char) ast.Clue {
	return ast.Clue{
		Columns: goalToLines(TransposeGoal(goal), background),
		Rows:    goalToLines(goal, background),
	}
}

func goalToLines(goal ast.Goal, background ast.Char) []ast.Line {
	var lines []ast.Line

	for r := range goal {
		currentLine := ast.Line{}
		previousColor := background
		count := 0

		for c := range goal[r] {
			currentColor := goal[r][c]
			if currentColor != previousColor {
				if previousColor != background {
					currentLine = append(currentLine, ast.Item{
						Color: previousColor,
						Count: count,
					})
				}

				count = 0
			}

			if currentColor != background {
				count++
			}

			previousColor = currentColor
		}

		if count > 0 {
			currentLine = append(currentLine, ast.Item{
				Color: previousColor,
				Count: count,
			})
		}

		lines = append(lines, currentLine)
	}

	return lines
}

// TransposeGoal returns transposed goal.
func TransposeGoal(goal ast.Goal) ast.Goal {
	rows := len(goal)
	if rows == 0 {
		return goal
	}

	columns := len(goal[0])
	if columns == 0 {
		return goal
	}

	transposed := make([][]ast.Char, columns)

	for c := 0; c < columns; c++ {
		transposed[c] = make([]ast.Char, rows)

		for r := 0; r < rows; r++ {
			transposed[c][r] = goal[r][c]
		}
	}

	return transposed
}
