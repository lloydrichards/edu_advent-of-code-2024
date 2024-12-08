package day08

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

func (p *Point) offset(other Point) Point {
	return Point{p.x - other.x, p.y - other.y}
}

func (p *Point) add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func (p *Point) inverse() Point {
	return Point{-p.x, -p.y}
}

type AntennaMap struct {
	grid     [][]string
	antennas map[string][]Point
}

func (m *AntennaMap) print() {
	for _, row := range m.grid {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *AntennaMap) printAntennas(p []Point) {
	for y, row := range m.grid {
		for x, c := range row {
			found := false
			for _, antenna := range p {
				if antenna.x == x && antenna.y == y {
					fmt.Print("X")
					found = true
					break
				}
			}
			if !found {
				fmt.Print(c)
			}

		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *AntennaMap) findPairs() [][]Point {
	pairs := [][]Point{}
	for _, requency := range m.antennas {
		for firstIdx, firstAntenna := range requency {
			for secondIdx, secondAntenna := range requency {
				if firstIdx != secondIdx {
					pairs = append(pairs, []Point{firstAntenna, secondAntenna})
				}
			}
		}
	}
	return pairs
}

func (m *AntennaMap) isInMap(a Point) bool {
	return a.x >= 0 && a.x < len(m.grid[0]) && a.y >= 0 && a.y < len(m.grid)
}

func parseMap(input string) AntennaMap {
	mapParts := U.SplitByNewline(input)

	m := AntennaMap{grid: make([][]string, len(mapParts)), antennas: make(map[string][]Point)}
	for y, mapPart := range mapParts {
		for x, c := range strings.Split(mapPart, "") {
			m.grid[y] = append(m.grid[y], c)
			if c != "." {
				m.antennas[c] = append(m.antennas[c], Point{x, y})
			}
		}

	}
	return m
}

func findAntiNodes(a Point, b Point) []Point {
	offset := a.offset(b)
	antiNodes := []Point{}
	antiNodes = append(antiNodes, a.add(offset))
	antiNodes = append(antiNodes, b.add(offset.inverse()))

	return antiNodes

}

func (m *AntennaMap) findHarmonicAntiNodes(a Point, b Point) []Point {
	offset := a.offset(b)
	antiNodes := []Point{}
	newLeftAntiNode := a.add(offset)
	leftOnMap := true
	for leftOnMap {
		if !m.isInMap(newLeftAntiNode) {
			leftOnMap = false
			break
		}
		antiNodes = append(antiNodes, newLeftAntiNode)
		newLeftAntiNode = newLeftAntiNode.add(offset)
	}

	inverseOffset := offset.inverse()
	newRightAntiNode := b.add(inverseOffset)
	rightOnMap := true
	for rightOnMap {
		if !m.isInMap(newRightAntiNode) {
			rightOnMap = false
			break
		}
		antiNodes = append(antiNodes, newRightAntiNode)
		newRightAntiNode = newRightAntiNode.add(inverseOffset)
	}
	return antiNodes

}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	antennaMap := parseMap(input)

	// antennaMap.print()

	allPairs := antennaMap.findPairs()

	uniqueAntiNodes := map[Point]bool{}
	for _, pair := range allPairs {
		antiNodes := findAntiNodes(pair[0], pair[1])
		for _, antiNode := range antiNodes {
			if antennaMap.isInMap(antiNode) {
				uniqueAntiNodes[antiNode] = true

			}
		}

	}

	return len(uniqueAntiNodes), nil
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	antennaMap := parseMap(input)

	// antennaMap.print()

	allPairs := antennaMap.findPairs()

	uniqueAntiNodes := map[Point]bool{}
	// allAntiNodes := []Point{}
	for _, pair := range allPairs {
		antiNodes := antennaMap.findHarmonicAntiNodes(pair[0], pair[1])
		uniqueAntiNodes[pair[0]] = true
		uniqueAntiNodes[pair[1]] = true
		for _, antiNode := range antiNodes {
			uniqueAntiNodes[antiNode] = true
			// allAntiNodes = append(allAntiNodes, antiNode)

		}

	}

	// antennaMap.printAntennas(allAntiNodes)

	return len(uniqueAntiNodes), nil
}
