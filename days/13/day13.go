package day00

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	x, y int
}

type ClawMachine struct {
	a     Button
	b     Button
	prize [2]int
}

func (c *ClawMachine) Play() ([2]int, error) {

	for aPress := 0; aPress < 100; aPress++ {
		for bPress := 0; bPress < 100; bPress++ {
			totalX := aPress*c.a.x + bPress*c.b.x
			totalY := aPress*c.a.y + bPress*c.b.y

			if totalX == c.prize[0] && totalY == c.prize[1] {
				return [2]int{aPress, bPress}, nil
			}
		}
	}

	return [2]int{-1, -1}, fmt.Errorf("no solution found")
}

func (c *ClawMachine) PlayMath() ([2]int, error) {

	// px = ax * a + bx * b
	// py = ay * a + by * b

	// a = (px - bx * b) / ax

	// py = ay * ((px - bx * b) / ax) + by * b
	// b = (py - ay * ((px - bx * b) / ax)) / by

	// b = (py * ax - ay * px) / (by * ax - ay * bx)

	bPress := (c.prize[1]*c.a.x - c.a.y*c.prize[0]) / (c.b.y*c.a.x - c.a.y*c.b.x)
	aPress := (c.prize[0] - bPress*c.b.x) / c.a.x

	totalX := aPress*c.a.x + bPress*c.b.x
	totalY := aPress*c.a.y + bPress*c.b.y

	if totalX == c.prize[0] && totalY == c.prize[1] {
		return [2]int{aPress, bPress}, nil
	}

	return [2]int{-1, -1}, fmt.Errorf("no solution found")
}

func parseGames(input string, offset int) []ClawMachine {

	gamesStr := strings.Split(input, "\n\n")
	games := []ClawMachine{}

	for _, game := range gamesStr {
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllString(game, -1)

		a := Button{}
		b := Button{}
		prize := [2]int{}

		a.x, _ = strconv.Atoi(nums[0])
		a.y, _ = strconv.Atoi(nums[1])
		b.x, _ = strconv.Atoi(nums[2])
		b.y, _ = strconv.Atoi(nums[3])
		prize[0], _ = strconv.Atoi(nums[4])
		prize[1], _ = strconv.Atoi(nums[5])

		prize[0] += offset
		prize[1] += offset

		games = append(games, ClawMachine{a, b, prize})
	}

	return games
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	games := parseGames(input, 0)

	tokens := 0
	for _, game := range games {
		result, err := game.Play()
		if err != nil {
			continue
		}
		tokens += result[0]*3 + result[1]

	}

	return tokens, nil
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	games := parseGames(input, 10000000000000)

	tokens := 0
	for _, game := range games {
		result, err := game.PlayMath()
		if err != nil {
			continue
		}
		tokens += result[0]*3 + result[1]

	}

	return tokens, nil
}
