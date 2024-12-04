package day04

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strings"
)

type coord struct {
	x int
	y int
}

type Puzzle struct {
	board [][]string
	start []coord
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func parsePuzzle(input string) Puzzle {
	lines := strings.Split(input, "\n")
	board := [][]string{}
	start := []coord{}

	for y, l := range lines {
		row := []string{}
		for x, c := range strings.Split(l, "") {
			if c == "X" {
				start = append(start, coord{x, y})
			}
			row = append(row, string(c))
		}
		board = append(board, row)
	}
	return Puzzle{board, start}
}

func (p *Puzzle) printBoard() {
	for _, row := range p.board {
		fmt.Println(row)
	}
}

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
	puzzle := parsePuzzle(input)

	// puzzle.printBoard()
	matches := puzzle.findMatches()

	return len(matches), nil
}

func Part2(dir string) (int, error) {

	return -1, nil
}
