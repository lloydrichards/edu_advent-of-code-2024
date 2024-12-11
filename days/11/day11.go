package day11

import (
	U "advent-of-code-2024/internal/utils"
	"strconv"
	"strings"
)

type Stone int

func (s *Stone) blink() []Stone {
	strValue := strconv.Itoa(int(*s))

	if *s == 0 {
		return []Stone{1}
	}

	if len(strValue)%2 == 0 {
		firstHalf, _ := strconv.Atoi(strValue[:len(strValue)/2])
		secondHalf, _ := strconv.Atoi(strValue[len(strValue)/2:])
		return []Stone{Stone((firstHalf)), Stone(secondHalf)}
	}

	return []Stone{Stone(*s * 2024)}
}

func parseInput(input string) []Stone {
	stones := []Stone{}
	for _, char := range strings.Fields(input) {
		num, _ := strconv.Atoi(char)
		stones = append(stones, Stone(num))
	}
	return stones
}

func flatten(stones [][]Stone) []Stone {
	flat := []Stone{}
	for _, row := range stones {
		flat = append(flat, row...)
	}
	return flat
}

func Part1(dir string) (int, error) {
	input, _ := U.LoadInputFile(dir)
	stones := parseInput(input)

	for i := 1; i <= 25; i++ {
		newStones := [][]Stone{}
		for _, stone := range stones {
			blinked := stone.blink()
			newStones = append(newStones, blinked)
		}

		stones = flatten(newStones)
	}

	return len(stones), nil
}

func Part2(dir string) (int, error) {
	// input, err := U.LoadInputFile(dir)
	// if err != nil {
	// 	return -1, err
	// }

	return -1, nil
}
