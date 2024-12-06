package main

import (
	"fmt"

	day01 "advent-of-code-2024/days/01"
	day02 "advent-of-code-2024/days/02"
	day03 "advent-of-code-2024/days/03"
	day04 "advent-of-code-2024/days/04"
	day05 "advent-of-code-2024/days/05"
	day06 "advent-of-code-2024/days/06"
)

func main() {
	fmt.Printf("🎄🎄🎄 Welcome to the Advent of Code 2024 🎄🎄🎄\n\n")

	fmt.Println("--- Day 1: Historian Hysteria ---")
	day01Part1, _ := day01.Part1("days/01/data/input.txt")
	fmt.Printf("⭐️ Part 1: distance = %d\n", day01Part1)
	day01Part2, _ := day01.Part2("days/01/data/input.txt")
	fmt.Printf("⭐️ Part 2: score = %d\n\n", day01Part2)

	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	day02Part1, _ := day02.Part1("days/02/data/input.txt")
	fmt.Printf("⭐️ Part 1: reports = %d\n", day02Part1)
	day02Part2, _ := day02.Part2("days/02/data/input.txt")
	fmt.Printf("⭐️ Part 2: dampened reports = %d\n\n", day02Part2)

	fmt.Println("--- Day 3: Mull It Over ---")
	day03Part1, _ := day03.Part1("days/03/data/input.txt")
	fmt.Printf("⭐️ Part 1: result = %d\n", day03Part1)
	day03Part2, _ := day03.Part2("days/03/data/input.txt")
	fmt.Printf("⭐️ Part 2: result = %d\n\n", day03Part2)

	fmt.Println("--- Day 4: Ceres Search ---")
	day04Part1, _ := day04.Part1("days/04/data/input.txt")
	fmt.Printf("⭐️ Part 1: matches = %d\n", day04Part1)
	day04Part2, _ := day04.Part2("days/04/data/input.txt")
	fmt.Printf("⭐️ Part 2: matches = %d\n\n", day04Part2)
	
	fmt.Println("--- Day 5: Print Queue ---")
	day05Part1, _ := day05.Part1("days/05/data/input.txt")
	fmt.Printf("⭐️ Part 1: matches = %d\n", day05Part1)
	day05Part2, _ := day05.Part2("days/05/data/input.txt")
	fmt.Printf("⭐️ Part 2: matches = %d\n\n", day05Part2)

	fmt.Println("--- Day 6: Guard Gallivant ---")
	day06Part1, _ := day06.Part1("days/06/data/input.txt")
	fmt.Printf("⭐️ Part 1: positions = %d\n", day06Part1)

}
