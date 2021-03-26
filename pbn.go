package pbn

type PuzzleSet struct {
	Puzzle []Puzzle
}

type Puzzle struct {
	BackgroundColor *Char
	DefaultColor    *Char
	Source          string
	ID              string
	Title           string
	Author          *Author
	Copyright       string
	Description     string
	Colors          Colors
	Clue            Clue
	Goal            *Goal
}

type Author struct {
	Name string
	ID   string
}

type Colors []Color

type Color struct {
	Char Char
	Hex  string
}

type Char rune

type Clue struct {
	Columns []Line
	Rows    []Line
}

type Line struct {
	Color *Char
	Count int
}

type Goal [][]Char
