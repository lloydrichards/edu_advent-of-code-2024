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
	bounds   [2]int
	boxes    []Pos
	bigBoxes [][2]Pos
	walls    []Pos
	robot    Robot
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
func parseBigWarehouse(input string) Warehouse {

	parts := strings.Split(input, "\n\n")

	warehouse := Warehouse{}
	for y, line := range strings.Split(parts[0], "\n") {
		warehouse.bounds[1] = y
		for x := 0; x < len(line)*2; x += 2 {
			warehouse.bounds[0] = x
			switch line[x/2] {
			case '#':
				warehouse.walls = append(warehouse.walls, Pos{x, y})
				warehouse.walls = append(warehouse.walls, Pos{x + 1, y})
			case 'O':
				warehouse.bigBoxes = append(warehouse.bigBoxes, [2]Pos{{x, y}, {x + 1, y}})
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

func (w *Warehouse) printBig() {
	for y := 0; y <= w.bounds[1]; y++ {
		for x := 0; x <= w.bounds[0]; x++ {
			found := false
			for _, box := range w.bigBoxes {
				if box[0].x == x && box[0].y == y {
					fmt.Print("[")
					found = true
					break
				}
				if box[1].x == x && box[1].y == y {
					fmt.Print("]")
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

func (w *Warehouse) pushBigBox(box Pos, dir Direction) ([][2]Pos, error) {
	movableBoxes := [][2]Pos{}
	nextPos := newPos(box, dir)

	for _, wall := range w.walls {
		if wall == nextPos {
			return nil, fmt.Errorf("can't push box to wall")
		}
	}

	switch dir {
	case Right:
		for _, b := range w.bigBoxes {
			if b[0] == nextPos {
				more := [][2]Pos{}
				for _, p := range b {
					m, err := w.pushBigBox(p, dir)
					if err != nil {
						return [][2]Pos{}, err
					}
					more = append(more, m...)
				}
				movableBoxes = append(movableBoxes, b)
				movableBoxes = append(movableBoxes, more...)

			}
		}
	case Left:
		for _, b := range w.bigBoxes {
			if b[1] == nextPos {
				more := [][2]Pos{}
				for _, p := range b {
					m, err := w.pushBigBox(p, dir)
					if err != nil {
						return [][2]Pos{}, err
					}
					more = append(more, m...)
				}
				movableBoxes = append(movableBoxes, b)
				movableBoxes = append(movableBoxes, more...)

			}
		}
	case Up, Down:
		for _, b := range w.bigBoxes {
			if b[0] == nextPos {
				more := [][2]Pos{}
				for _, p := range b {
					m, err := w.pushBigBox(p, dir)
					if err != nil {
						return [][2]Pos{}, err
					}
					more = append(more, m...)
				}
				movableBoxes = append(movableBoxes, b)
				movableBoxes = append(movableBoxes, more...)

			}
			if b[1] == nextPos {
				more := [][2]Pos{}
				for _, p := range b {
					m, err := w.pushBigBox(p, dir)
					if err != nil {
						return [][2]Pos{}, err
					}
					more = append(more, m...)
				}
				movableBoxes = append(movableBoxes, b)
				movableBoxes = append(movableBoxes, more...)

			}
		}

	}

	return movableBoxes, nil
}

func (w *Warehouse) simulateBig() {
	for _, dir := range w.robot.queue {
		// dirStr := ""
		// switch dir {
		// case Up:
		// 	dirStr = "^"
		// case Down:
		// 	dirStr = "v"
		// case Left:
		// 	dirStr = "<"
		// case Right:
		// 	dirStr = ">"
		// }
		// fmt.Printf("Move: %s\n", dirStr)
		moving, err := w.pushBigBox(w.robot.pos, dir)
		if err == nil {
			w.robot.pos = newPos(w.robot.pos, dir)
		}
		for i, b := range w.bigBoxes {
			for _, box := range moving {
				if b == box {
					w.bigBoxes[i] = [2]Pos{newPos(b[0], dir), newPos(b[1], dir)}
				}
			}
		}
		// w.printBig()
	}
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	warehouse := parseBigWarehouse(input)

	// warehouse.printBig()

	warehouse.simulateBig()

	total := 0

	for _, box := range warehouse.bigBoxes {
		gps := box[0].x + box[0].y*100
		total += gps
	}

	return total, nil
}
