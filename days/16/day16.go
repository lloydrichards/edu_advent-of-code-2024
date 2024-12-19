package day16

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
)

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

type Pos struct {
	x, y int
}

type Cell struct {
	x, y     int
	cost     int
	passable bool
	origin   *Cell
	dir      Direction
}

type Maze struct {
	grid   [][]Cell
	start  Pos
	finish Pos
}

func (m Maze) print(runner [2]int) {
	for y := range m.grid {
		for x := range m.grid[y] {
			if x == runner[0] && y == runner[1] {
				fmt.Print("R")
			} else if m.start.x == x && m.start.y == y {
				fmt.Print("S")
			} else if m.finish.x == x && m.finish.y == y {
				fmt.Print("F")
			} else if m.grid[y][x].passable {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (m Maze) printTrail(trail []Cell) {
	for y := range m.grid {
		for x := range m.grid[y] {
			if m.start.x == x && m.start.y == y {
				fmt.Print("S")
			} else if m.finish.x == x && m.finish.y == y {
				fmt.Print("F")
			} else if m.grid[y][x].passable {
				if contain(trail, m.grid[y][x]) {
					for _, t := range trail {
						if t.x == x && t.y == y {
							switch t.dir {
							case N:
								fmt.Print("^")
							case E:
								fmt.Print(">")
							case S:
								fmt.Print("v")
							case W:
								fmt.Print("<")
							}
						}
					}
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

type Queue struct {
	open   []Cell
	closed []Cell
}

func (q *Queue) pop() Cell {
	cell := q.open[0]
	q.open = q.open[1:]
	return cell
}

func (q *Queue) push(cell Cell) {
	for i, c := range q.open {
		if cell.cost < c.cost {
			q.open = append(q.open[:i], append([]Cell{cell}, q.open[i:]...)...)
			q.sort()
			return
		}
	}
	q.open = append(q.open, cell)
}

func (q *Queue) sort() {
	for i := 0; i < len(q.open); i++ {
		for j := i + 1; j < len(q.open); j++ {
			if q.open[i].cost > q.open[j].cost {
				q.open[i], q.open[j] = q.open[j], q.open[i]
			}
		}
	}
}

func (q *Queue) update(cell Cell) {
	for i, c := range q.open {
		if cell == c {
			q.open[i] = cell
			q.sort()
			return
		}
	}
}

func (q *Queue) trail(c Cell) []Cell {
	trail := []Cell{c}
	for c.origin != nil {
		trail = append(trail, *c.origin)
		c = *c.origin
	}
	return trail
}

func contain(cells []Cell, cell Cell) bool {
	for _, c := range cells {
		if c.x == cell.x && c.y == cell.y {
			return true
		}
	}
	return false
}

func (m *Maze) runner() (Cell, Queue) {
	queue := Queue{
		open:   []Cell{m.grid[m.start.y][m.start.x]},
		closed: []Cell{}}

	for len(queue.open) > 0 {
		current := queue.pop()
		queue.closed = append(queue.closed, current)

		// if current.x == m.finish.x && current.y == m.finish.y {
		// 	return current, queue
		// }

		// check neighbors
		neighbors := []Cell{
			m.grid[current.y-1][current.x], // N
			m.grid[current.y][current.x+1], // E
			m.grid[current.y+1][current.x], // S
			m.grid[current.y][current.x-1], // W
		}

		for i, neighbor := range neighbors {
			if !neighbor.passable {
				continue
			}

			if current.origin == &neighbor {
				continue
			}

			if contain(queue.closed, neighbor) {
				continue
			}

			switch i {
			case 0:
				neighbor.dir = N
			case 1:
				neighbor.dir = E
			case 2:
				neighbor.dir = S
			case 3:
				neighbor.dir = W
			}

			temp_cost := current.cost + 1
			if neighbor.dir != current.dir {
				temp_cost = current.cost + 1001
			}

			if !contain(queue.open, neighbor) || temp_cost < neighbor.cost {
				neighbor.cost = temp_cost
				neighbor.origin = &current
				queue.push(neighbor)
			} else {
				queue.update(neighbor)
			}
		}
	}

	lowest := Cell{}
	for _, cell := range queue.closed {
		if cell.x == m.finish.x && cell.y == m.finish.y {
			if lowest.cost == 0 || cell.cost < lowest.cost {
				lowest = cell
			}
		}
	}
	return lowest, queue
}

func newMaze(input string) Maze {
	lines := U.SplitByNewline(input)
	maze := Maze{
		grid: make([][]Cell, len(lines)),
	}

	// find start and finish
	for y, row := range lines {
		for x, cell := range row {
			switch cell {
			case 'S':
				maze.start = Pos{x, y}
			case 'E':
				maze.finish = Pos{x, y}
			}
		}
	}
	// calculate heuristic
	for y, row := range lines {
		maze.grid[y] = make([]Cell, len(row))
		for x, cell := range row {

			maze.grid[y][x] = Cell{x, y, 0, cell != '#', nil, E}
		}
	}
	return maze
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	maze := newMaze(input)

	cell, _ := maze.runner()

	// trail := q.trail(cell)

	// maze.printTrail(trail)

	return cell.cost, nil
}

func Part2(dir string) (int, error) {
	// input, err := U.LoadInputFile(dir)
	// if err != nil {
	// 	return -1, err
	// }

	return -1, nil
}
