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

func (p *PatrolMap) copy() PatrolMap {
	newMap := make(PatrolMap, len(*p))
	for y, row := range *p {
		newMap[y] = make([]bool, len(row))
		copy(newMap[y], row)
	}
	return newMap
}

type Point struct {
	x, y int
}

type Turn struct {
	pos Point
	dir int
}

type Guard struct {
	startPos  Point
	curPos    Point
	turns     []Turn
	direction int
	onMap     bool
	inLoop    bool
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

func (g *Guard) copy() Guard {
	newGuard := Guard{
		startPos:  g.startPos,
		curPos:    g.curPos,
		turns:     make([]Turn, len(g.turns)),
		direction: g.direction,
		onMap:     g.onMap,
		inLoop:    g.inLoop,
		visited:   make(map[Point]bool),
	}
	for k, v := range g.visited {
		newGuard.visited[k] = v
	}
	copy(newGuard.turns, g.turns)

	return newGuard
}

func (g *Guard) checkRepeatTurn(t Turn) bool {
	for _, turn := range g.turns {
		if turn == t {
			return true
		}
	}
	return false
}

func (g *Guard) patrol(m PatrolMap) {
	dirOffset := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	freeSpace := true
	for freeSpace {

		nextPos := Point{g.curPos.x + dirOffset[g.direction][0], g.curPos.y + dirOffset[g.direction][1]}

		if !m.onMap(nextPos) {
			g.onMap = false
			g.visited[g.curPos] = true
			freeSpace = false
			break
		}

		if m.isBlocked(nextPos) {
			g.direction = (g.direction + 1) % 4
			newTurn := Turn{g.curPos, g.direction}
			if g.checkRepeatTurn(newTurn) {
				g.inLoop = true
				g.onMap = false
				freeSpace = false
				break
			}
			g.turns = append(g.turns, newTurn)
			freeSpace = false
			break
		}

		g.visited[g.curPos] = true
		g.curPos = nextPos
	}
}

func parseMapAndGuard(input string) (PatrolMap, Guard) {
	var patrolMap PatrolMap
	guard := Guard{visited: make(map[Point]bool), onMap: true, inLoop: false, direction: 0}
	for y, line := range strings.Split(input, "\n") {
		var row []bool
		for x, c := range line {
			if c == '^' {
				guard.startPos = Point{x, y}
				guard.curPos = Point{x, y}
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
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	baseMap, baseGuard := parseMapAndGuard(input)

	count := 0
	for y, row := range baseMap {
		for x := range row {
			if baseMap.isBlocked(Point{x, y}) || (x == baseGuard.startPos.x && y == baseGuard.startPos.y) {
				continue
			}

			newMap := baseMap.copy()
			newMap[y][x] = true
			newGuard := baseGuard.copy()

			for newGuard.onMap {
				newGuard.patrol(newMap)
			}
			if newGuard.inLoop {
				count++
			}

		}
	}

	return count, nil
}
