package day14

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Robot struct {
	x, y   int
	v      [2]int
	bounds [2]int
}

func (r *Robot) step(steps int) {
	// fmt.Println("Starting at", r.x, r.y)
	// fmt.Println("Moving", r.v)

	for i := 0; i < steps; i++ {
		newX := r.x + r.v[0]
		newY := r.y + r.v[1]

		// loop to the other side
		if newX < 0 {
			r.x = r.bounds[0] + newX
		} else if newX >= r.bounds[0] {
			r.x = newX - r.bounds[0]
		} else {
			r.x = newX
		}

		if newY < 0 {
			r.y = r.bounds[1] + newY
		} else if newY >= r.bounds[1] {
			r.y = newY - r.bounds[1]
		} else {
			r.y = newY
		}
		// r.print()
	}

}

func (r *Robot) print() {

	fmt.Println("At", r.x, r.y)
	for i := 0; i < r.bounds[1]; i++ {
		for j := 0; j < r.bounds[0]; j++ {
			if r.y == i && r.x == j {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (r *Robot) getQuad() int {
	meanX := r.bounds[0] / 2
	meanY := r.bounds[1] / 2

	if r.x < meanX && r.y < meanY {
		return 1
	} else if r.x > meanX && r.y < meanY {
		return 2
	} else if r.x > meanX && r.y > meanY {
		return 3
	} else if r.x < meanX && r.y > meanY {
		return 4
	}
	return -1
}

func printRobots(robots []Robot, bounds [2]int) {
	for y := 0; y < bounds[1]; y++ {
		for x := 0; x < bounds[0]; x++ {
			found := false
			for _, robot := range robots {
				if robot.x == x && robot.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseRobots(input string, bounds [2]int) []Robot {

	robots := []Robot{}
	for _, line := range U.SplitByNewline(input) {
		re := regexp.MustCompile(`[+-]?\d+`)
		nums := re.FindAllString(line, -1)

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		vx, _ := strconv.Atoi(nums[2])
		vy, _ := strconv.Atoi(nums[3])

		robots = append(robots, Robot{x, y, [2]int{vx, vy}, bounds})
	}
	return robots
}

func Part1(dir string, bounds [2]int) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	robots := parseRobots(input, bounds)

	afterRobots := []Robot{}
	for _, robot := range robots {
		robot.step(100)
		afterRobots = append(afterRobots, robot)

	}

	quads := make(map[int]int)
	for _, robot := range afterRobots {
		quad := robot.getQuad()
		if quad == -1 {
			continue
		}
		quads[quad]++
	}

	// printRobots(afterRobots, bounds)
	total := 1

	for _, v := range quads {
		total *= v
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
