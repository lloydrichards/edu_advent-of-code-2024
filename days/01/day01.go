package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(scanner *bufio.Scanner) ([][]int, error) {
	elfList := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into two parts
		strList := strings.Fields(line)

		for idx, idStr := range strList {
			locationId, err := strconv.Atoi(idStr)
			if err != nil {
				return nil, err
			}
			if len(elfList) < idx+1 {
				elfList = append(elfList, []int{})
			}
			elfList[idx] = append(elfList[idx], locationId)
		}
	}
	return elfList, nil
}

func sortNumbers(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

func findDistance(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sumNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Part1(dir string) (int, error) {
	log.Println("--Day 1: Historian Hysteria--")
	log.Println("----Part 1----")
	// Read the input file
	file, err := os.Open(dir)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfList, err := parseInput(scanner)
	if err != nil {
		return -1, err
	}

	// Sort the numbers
	for idx, elf := range elfList {
		elfList[idx] = sortNumbers(elf)
	}

	distances := []int{}
	for idx := range elfList[0] {
		distances = append(distances, findDistance(elfList[0][idx], elfList[1][idx]))
	}

	return sumNumbers(distances), nil
}
