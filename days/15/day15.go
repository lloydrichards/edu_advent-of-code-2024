package day15

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strings"
)

type Pos struct {
	x, y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Robot struct {
	pos   Pos
	queue []Direction
}

type Warehouse struct {
	bounds [2]int
	boxes  []Pos
	walls  []Pos
	robot  Robot
}

func newPos(pos Pos, dir Direction) Pos {
	switch dir {
	case Up:
		pos.y--
	case Down:
		pos.y++
	case Left:
		pos.x--
	case Right:
		pos.x++
	}
	return pos
}

func (w *Warehouse) pushBox(box Pos, dir Direction) bool {
	new := newPos(box, dir)

	for _, wall := range w.walls {
		if wall == new {
			return false
		}
	}

	for i, b := range w.boxes {
		if b == new {
			canPush := w.pushBox(b, dir)
			if !canPush {
				return false
			}
			w.boxes[i] = newPos(b, dir)
		}
	}

	return true
}

func parseWarehouse(input string) Warehouse {

	parts := strings.Split(input, "\n\n")

	warehouse := Warehouse{}
	for y, line := range strings.Split(parts[0], "\n") {
		warehouse.bounds[1] = y
		for x, char := range line {
			warehouse.bounds[0] = x
			switch char {
			case '#':
				warehouse.walls = append(warehouse.walls, Pos{x, y})
			case 'O':
				warehouse.boxes = append(warehouse.boxes, Pos{x, y})
			case '@':
				warehouse.robot = Robot{Pos{x, y}, []Direction{}}
			}
		}
	}
	for _, line := range strings.Split(parts[1], "\n") {
		for _, char := range line {
			switch char {
			case '^':
				warehouse.robot.queue = append(warehouse.robot.queue, Up)
			case 'v':
				warehouse.robot.queue = append(warehouse.robot.queue, Down)
			case '<':
				warehouse.robot.queue = append(warehouse.robot.queue, Left)
			case '>':
				warehouse.robot.queue = append(warehouse.robot.queue, Right)
			}
		}
	}
	return warehouse
}

func (w *Warehouse) print() {
	for y := 0; y <= w.bounds[1]; y++ {
		for x := 0; x <= w.bounds[0]; x++ {
			found := false
			for _, box := range w.boxes {
				if box.x == x && box.y == y {
					fmt.Print("O")
					found = true
					break
				}
			}
			if !found {
				for _, wall := range w.walls {
					if wall.x == x && wall.y == y {
						fmt.Print("#")
						found = true
						break
					}
				}
			}
			if !found {
				if w.robot.pos.x == x && w.robot.pos.y == y {
					fmt.Print("@")
					found = true
				}
			}
			if !found {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (w *Warehouse) simulate() {
	for _, dir := range w.robot.queue {
		// fmt.Printf("Move: %d\n", dir)
		if w.pushBox(w.robot.pos, dir) {
			w.robot.pos = newPos(w.robot.pos, dir)
		}
		// w.print()
	}
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	warehouse := parseWarehouse(input)

	warehouse.simulate()

	total := 0
	for _, box := range warehouse.boxes {
		gps := box.x + box.y*100
		total += gps
	}

	return total, nil
}

func Part2(dir string) (int, error) {
	// input, err := U.LoadInputFile(dir)
	// if err != nil {
	// 	return -1, err
	// }

	return -1, nil
}
