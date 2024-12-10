package day10

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strconv"
)

type Point [3]int

type Trail []Point

type TopoMap struct {
	grid   [][]int
	peaks  []Point
	starts []Point
}

func (t *TopoMap) printTrails(trails []Trail) {
	empty := make([][]string, len(t.grid))
	for y := 0; y < len(t.grid); y++ {
		empty[y] = make([]string, len(t.grid[0]))
		for x := 0; x < len(t.grid[0]); x++ {
			empty[y][x] = "."
		}
	}
	for _, trail := range trails {
		for _, point := range trail {
			empty[point[1]][point[0]] = strconv.Itoa(point[2])
		}
	}
	for y := 0; y < len(t.grid); y++ {
		for x := 0; x < len(t.grid[0]); x++ {
			if empty[y][x] == "9" {
				fmt.Print("^")
			} else if empty[y][x] == "0" {
				fmt.Print("S")
			} else {
				fmt.Print(empty[y][x])
			}
		}
		fmt.Println()
	}
}

type Hiker struct {
	curPos Point
	path   Trail
}

type Hikers []Hiker

func (h *Hikers) push(hiker Hiker) {
	*h = append(*h, hiker)
}

func (h *Hikers) pop(idx int) Hiker {
	hiker := (*h)[idx]
	*h = append((*h)[:idx], (*h)[idx+1:]...)
	return hiker
}

func parseTypography(input string) TopoMap {
	lines := U.SplitByNewline(input)
	topo := TopoMap{grid: make([][]int, len(lines))}

	for i, line := range lines {
		topo.grid[i] = make([]int, len(line))

		for j, char := range line {
			num, _ := strconv.Atoi(string(char))
			topo.grid[i][j] = num

			if num == 9 {
				topo.peaks = append(topo.peaks, Point{j, i, num})
			}
			if num == 0 {
				topo.starts = append(topo.starts, Point{j, i, num})
			}
		}
	}

	return topo
}

func hikeTrail(topo TopoMap, hikers Hikers, finished map[Point][]Trail) map[Point][]Trail {
	if len(hikers) == 0 {
		return finished
	}

	hiker := hikers.pop(0)

	// if hiker is at the end, add to finished
	if hiker.curPos[2] == 9 {
		finished[hiker.path[0]] = append(finished[hiker.path[0]], hiker.path)
		return hikeTrail(topo, hikers, finished)
	}

	// check if hiker can move up, down, left, right
	dir := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, d := range dir {
		newX := hiker.curPos[0] + d[0]
		newY := hiker.curPos[1] + d[1]

		if newX < 0 || newX >= len(topo.grid[0]) || newY < 0 || newY >= len(topo.grid) {
			continue
		}

		newPos := Point{newX, newY, topo.grid[newY][newX]}
		if newPos[2]-hiker.curPos[2] != 1 {
			continue
		}

		newPath := make(Trail, len(hiker.path))
		copy(newPath, hiker.path)
		newPath = append(newPath, newPos)

		newHiker := Hiker{curPos: newPos, path: newPath}
		hikers.push(newHiker)

	}
	return hikeTrail(topo, hikers, finished)
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	typography := parseTypography(input)

	hikers := make(Hikers, len(typography.starts))
	for i, start := range typography.starts {
		hikers[i] = Hiker{curPos: start, path: []Point{start}}
	}

	finishedTrails := hikeTrail(typography, hikers, map[Point][]Trail{})

	sum := 0
	for _, paths := range finishedTrails {
		uniquePeaks := map[Point]bool{}
		for _, path := range paths {
			uniquePeaks[path[len(path)-1]] = true
		}
		sum += len(uniquePeaks)
	}
	return sum, nil
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	typography := parseTypography(input)

	hikers := make(Hikers, len(typography.starts))
	for i, start := range typography.starts {
		hikers[i] = Hiker{curPos: start, path: []Point{start}}
	}

	finishedTrails := hikeTrail(typography, hikers, map[Point][]Trail{})

	sum := 0
	for _, paths := range finishedTrails {
		sum += len(paths)
	}
	return sum, nil
}
