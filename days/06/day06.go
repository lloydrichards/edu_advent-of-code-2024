package day06

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strings"
)

type PatrolMap [][]bool

func (p *PatrolMap) print() {
	for _, row := range *p {
		for _, cell := range row {
			if cell {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func (p *PatrolMap) isBlocked(pos Point) bool {
	return (*p)[pos.y][pos.x]
}

func (p *PatrolMap) onMap(pos Point) bool {
	return pos.x >= 0 && pos.x < len((*p)[0]) && pos.y >= 0 && pos.y < len(*p)
}

type Point struct {
	x, y int
}

type Guard struct {
	position  Point
	direction int
	onMap     bool
	visited   map[Point]bool
}

func (g *Guard) printPatrol(m PatrolMap) {
	for y, row := range m {
		for x, cell := range row {
			if g.visited[Point{x, y}] {
				fmt.Print("X")
				continue
			}
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g *Guard) patrol(m PatrolMap) {
	dirOffset := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	freeSpace := true
	for freeSpace {
		nextPos := Point{g.position.x + dirOffset[g.direction][0], g.position.y + dirOffset[g.direction][1]}
		if !m.onMap(nextPos) {
			g.onMap = false
			g.visited[g.position] = true
			freeSpace = false
			break
		}
		if m.isBlocked(nextPos) {
			g.direction = (g.direction + 1) % 4
			freeSpace = false
			break
		}
		g.visited[g.position] = true
		g.position = nextPos
	}
}

func parseMapAndGuard(input string) (PatrolMap, Guard) {
	var patrolMap PatrolMap
	guard := Guard{visited: make(map[Point]bool), onMap: true, direction: 0}
	for y, line := range strings.Split(input, "\n") {
		var row []bool
		for x, c := range line {
			if c == '^' {
				guard.position = Point{x, y}
			}
			row = append(row, c == '#')
		}
		patrolMap = append(patrolMap, row)
	}

	return patrolMap, guard
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	patrolMap, guard := parseMapAndGuard(input)
	// patrolMap.print()

	for guard.onMap {
		guard.patrol(patrolMap)
	}

	// fmt.Println("-----------------")
	// guard.printPatrol(patrolMap)

	return len(guard.visited), nil
}

func Part2(dir string) (int, error) {
	// input, err := U.LoadInputFile(dir)
	// if err != nil {
	// 	return -1, err
	// }

	return -1, nil
}
