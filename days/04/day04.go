package day04

import (
	U "advent-of-code-2024/internal/utils"
	"strings"
)

type coord struct {
	x int
	y int
}

type Puzzle struct {
	startChar string
	board     [][]string
	start     []coord
}

func parsePuzzle(input string, startChar string) Puzzle {
	lines := strings.Split(input, "\n")
	board := [][]string{}
	start := []coord{}

	for y, l := range lines {
		row := []string{}
		for x, c := range strings.Split(l, "") {
			if c == startChar {
				start = append(start, coord{x, y})
			}
			row = append(row, string(c))
		}
		board = append(board, row)
	}
	return Puzzle{startChar, board, start}
}

// func (p *Puzzle) printBoard() {
// 	for _, row := range p.board {
// 		for _, c := range row {
// 			if c == p.startChar {
// 				fmt.Print(c)
// 				continue
// 			}
// 			fmt.Print(strings.ToLower(c))
// 		}
// 		fmt.Println()
// 	}
// }

func (p *Puzzle) validPath(start coord, xOffset int, yOffset int) bool {
	validChars := []string{"X", "M", "A", "S"}
	valid := false
	for idx, c := range validChars {
		if start.x+xOffset*idx < 0 || start.x+xOffset*idx >= len(p.board[0]) {
			valid = false
			break
		}
		if start.y+yOffset*idx < 0 || start.y+yOffset*idx >= len(p.board) {
			valid = false
			break
		}
		if p.board[start.y+yOffset*idx][start.x+xOffset*idx] == c {
			valid = true
		} else {
			valid = false
			break
		}
	}
	return valid

}

func (p *Puzzle) findMatches() []coord {
	matched := []coord{}
	for _, s := range p.start {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if p.validPath(s, x, y) {
					matched = append(matched, s)
				}
			}
		}

	}
	return matched
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	puzzle := parsePuzzle(input, "X")

	// puzzle.printBoard()
	matches := puzzle.findMatches()

	return len(matches), nil
}

func (p *Puzzle) validX(start coord) bool {

	validWord := 0

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 || y == 0 {
				continue
			}
			if start.x < 1 || start.x >= len(p.board[0])-1 {
				continue
			}
			if start.y < 1 || start.y >= len(p.board)-1 {
				continue
			}
			if p.board[start.y+y][start.x+x] == "M" {
				if p.board[start.y-y][start.x-x] == "S" {
					validWord++

				}
			}

		}
	}

	return validWord == 2

}

func (p *Puzzle) findXMAS() []coord {
	matched := []coord{}
	for _, s := range p.start {
		if p.validX(s) {
			matched = append(matched, s)
		}

	}
	return matched
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	puzzle := parsePuzzle(input, "A")

	// puzzle.printBoard()

	matches := puzzle.findXMAS()

	return len(matches), nil
}
