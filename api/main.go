package main

import (
	"log"

	day01 "advent-of-code-2024/days/01"
	day02 "advent-of-code-2024/days/02"
)

func main() {
	log.Println("Welcome to the Advent of Code 2024")

	log.Println("--- Day 1: Historian Hysteria ---")
	day1Part1, _ := day01.Part1("days/01/data/input.txt")
	log.Printf("⭐️ Part 1: distance = %d", day1Part1)
	day1Part2, _ := day01.Part2("days/01/data/input.txt")
	log.Printf("⭐️ Part 2: score = %d", day1Part2)

	log.Println("--- Day 2: Red-Nosed Reports --- ")
	day2Part1, _ := day02.Part1("days/02/data/input.txt")
	log.Printf("⭐️ Part 1: valid reports = %d", day2Part1)

}
