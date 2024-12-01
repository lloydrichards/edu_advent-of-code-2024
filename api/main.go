package main

import (
	"log"

	day01 "advent-of-code-2024/days/01"
)

func main() {

	log.Println("Welcome to the Advent of Code 2024")
	result, err := day01.Part1("days/01/data/input.txt")
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println("Part1: ", result)
}
