package day07

import (
	U "advent-of-code-2024/internal/utils"
	"strconv"
	"strings"
)

type Test struct {
	name   int
	values []int
}

func parseTests(input string) []Test {
	testParts := strings.Split(input, "\n")

	tests := []Test{}
	for _, testPart := range testParts {
		parts := strings.Split(testPart, ":")

		testName, _ := strconv.Atoi(parts[0])
		testStrValues := strings.Fields(parts[1])

		testValues := []int{}

		for _, vStr := range testStrValues {
			value, _ := strconv.Atoi(vStr)
			testValues = append(testValues, value)
		}

		tests = append(tests, Test{name: testName, values: testValues})
	}

	return tests
}

func generateCombinations(n int) [][]bool {
	combinations := [][]bool{}
	totalCombinations := 1 << n // 2^n combinations

	for i := 0; i < totalCombinations; i++ {
		combination := make([]bool, n)
		for j := 0; j < n; j++ {
			combination[j] = (i & (1 << j)) != 0
		}
		combinations = append(combinations, combination)
	}

	return combinations
}

func solvablePredicate(target int, values []int) bool {
	for _, combination := range generateCombinations(len(values) - 1) {
		comboTotal := values[0]
		for i, operation := range combination {
			if operation {
				comboTotal += values[i+1]
			} else {
				comboTotal *= values[i+1]
			}
		}

		if target == comboTotal {
			return true
		}
	}
	return false
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	tests := parseTests(input)

	total := 0

	for _, test := range tests {
		isSolvable := solvablePredicate(test.name, test.values)
		if isSolvable {
			total += test.name
		}
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
