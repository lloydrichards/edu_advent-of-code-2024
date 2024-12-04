package main

import (
	"fmt"

	day01 "advent-of-code-2024/days/01"
	day02 "advent-of-code-2024/days/02"
	day03 "advent-of-code-2024/days/03"
	day04 "advent-of-code-2024/days/04"
)

func main() {
	fmt.Printf("🎄🎄🎄 Welcome to the Advent of Code 2024 🎄🎄🎄\n\n")

	fmt.Println("--- Day 1: Historian Hysteria ---")
	day1Part1, _ := day01.Part1("days/01/data/input.txt")
	fmt.Printf("⭐️ Part 1: distance = %d\n", day1Part1)
	day1Part2, _ := day01.Part2("days/01/data/input.txt")
	fmt.Printf("⭐️ Part 2: score = %d\n\n", day1Part2)

	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	day2Part1, _ := day02.Part1("days/02/data/input.txt")
	fmt.Printf("⭐️ Part 1: reports = %d\n", day2Part1)
	day2Part2, _ := day02.Part2("days/02/data/input.txt")
	fmt.Printf("⭐️ Part 2: dampened reports = %d\n\n", day2Part2)

	fmt.Println("--- Day 3: Mull It Over ---")
	day3Part1, _ := day03.Part1("days/03/data/input.txt")
	fmt.Printf("⭐️ Part 1: result = %d\n", day3Part1)
	day3Part2, _ := day03.Part2("days/03/data/input.txt")
	fmt.Printf("⭐️ Part 2: result = %d\n\n", day3Part2)

	fmt.Println("--- Day 4: Ceres Search ---")
	day4Part1, _ := day04.Part1("days/04/data/input.txt")
	fmt.Printf("⭐️ Part 1: matches = %d\n", day4Part1)
	day4Part2, _ := day04.Part2("days/04/data/input.txt")
	fmt.Printf("⭐️ Part 2: matches = %d\n\n", day4Part2)

}
