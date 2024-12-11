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

func memoizedBlink(stoneMap map[Stone]int) map[Stone]int {
	newStoneMap := map[Stone]int{}
	for stone, count := range stoneMap {
		str := strconv.Itoa(int(stone))
		if stone == 0 {
			newStoneMap[1] += count
		} else if len(str)%2 == 0 {
			firstHalf, _ := strconv.Atoi(str[:len(str)/2])
			secondHalf, _ := strconv.Atoi(str[len(str)/2:])
			newStoneMap[Stone(firstHalf)] += count
			newStoneMap[Stone(secondHalf)] += count
		} else {
			newStoneMap[Stone(int(stone)*2024)] += count
		}
	}
	return newStoneMap
}

func Part2(dir string) (int, error) {
	input, _ := U.LoadInputFile(dir)
	stones := parseInput(input)

	stoneMap := map[Stone]int{}
	for _, stone := range stones {
		stoneMap[stone]++
	}

	for i := 1; i <= 75; i++ {
		stoneMap = memoizedBlink(stoneMap)
	}

	total := 0
	for _, count := range stoneMap {
		total += count
	}
	return total, nil
}
