package day12

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strings"
)

type Plot [2]int

type Region struct {
	veg   string
	plots []Plot
	edges []Plot
}

func (r *Region) addPlot(p Plot) {
	r.plots = append(r.plots, p)
	r.calcEdges()
}

func (r *Region) isEdge(p Plot) bool {
	for _, edge := range r.edges {
		if edge == p {
			return true
		}
	}
	return false
}

func (r *Region) calcEdges() {
	r.edges = []Plot{}
	for _, plot := range r.plots {
		r.edges = append(r.edges, Plot{plot[0] - 1, plot[1]})
		r.edges = append(r.edges, Plot{plot[0] + 1, plot[1]})
		r.edges = append(r.edges, Plot{plot[0], plot[1] - 1})
		r.edges = append(r.edges, Plot{plot[0], plot[1] + 1})
	}
}

func (r *Region) calcFenceCost() int {
	outerFence := 0
	for _, edge := range r.edges {
		isPlot := false
		for _, plot := range r.plots {
			if edge == plot {
				isPlot = true
				break
			}
		}
		if !isPlot {
			outerFence++
		}
	}
	return outerFence * len(r.plots)
}

func (r *Region) print(maxX int, maxY int) {

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			isPlot := false
			for _, plot := range r.plots {
				if plot[0] == x && plot[1] == y {
					isPlot = true
					break
				}
			}
			if isPlot {
				fmt.Print(r.veg)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (r *Region) overlap(other Region) bool {
	for _, plot := range r.plots {
		for _, otherPlot := range other.plots {
			if plot == otherPlot {
				return true
			}
		}
	}
	return false
}

func parseRegions(input string) []Region {

	regions := []Region{}
	for y, line := range U.SplitByNewline(input) {
		for x, char := range strings.Split(line, "") {
			newPlot := Plot{x, y}
			if len(regions) == 0 {
				regions = append(regions, Region{veg: char, plots: []Plot{newPlot}})
				regions[0].calcEdges()
				continue
			}
			isEdge := false
			for i, region := range regions {
				if region.isEdge(newPlot) && region.veg == char {
					region.addPlot(newPlot)
					regions[i] = region
					isEdge = true
					continue
				}
			}
			if !isEdge {
				regions = append(regions, Region{veg: char, plots: []Plot{newPlot}})
				regions[len(regions)-1].calcEdges()
			}
		}
	}
	return regions
}

func uniquePlots(plots []Plot) []Plot {
	unique := []Plot{}
	for _, plot := range plots {
		isUnique := true
		for _, uPlot := range unique {
			if plot == uPlot {
				isUnique = false
				break
			}
		}
		if isUnique {
			unique = append(unique, plot)
		}
	}
	return unique
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	// x, y := U.GetGridSize(input)

	regions := parseRegions(input)

	for i, region := range regions {
		for j, other := range regions {
			if i == j {
				continue
			}
			if region.veg == other.veg {
				if region.overlap(other) {
					region.plots = uniquePlots(append(region.plots, other.plots...))
					region.calcEdges()
					regions[i] = region
					regions[j] = Region{}
				}
			}
		}
	}

	totalCost := 0
	for _, region := range regions {
		if len(region.plots) == 0 {
			continue
		}
		fence := region.calcFenceCost()
		// fmt.Printf("Region %s: %d plots, %d fence\n", region.veg, len(region.plots), fence)
		// region.print(x, y)
		totalCost += fence
	}

	return totalCost, nil

}

func Part2(dir string) (int, error) {
	// input, err := U.LoadInputFile(dir)
	// if err != nil {
	// 	return -1, err
	// }

	return -1, nil
}
