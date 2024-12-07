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

func generateCombinations(n int, base int) [][]int {
	combinations := [][]int{}
	totalCombinations := 1 // temp^n combinations
	for i := 0; i < n; i++ {
		totalCombinations *= base
	}

	for i := 0; i < totalCombinations; i++ {
		combination := make([]int, n)
		temp := i
		for j := 0; j < n; j++ {
			combination[j] = temp % base
			temp /= base
		}
		combinations = append(combinations, combination)
	}

	return combinations
}

func concatenateNumbers(a, b int) (int, error) {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)
	concatenatedStr := strA + strB
	concatenatedInt, err := strconv.Atoi(concatenatedStr)
	if err != nil {
		return 0, err
	}
	return concatenatedInt, nil
}

func solvablePredicate(target int, values []int, base int) bool {
	for _, combination := range generateCombinations(len(values)-1, base) {
		comboTotal := values[0]
		for i, operation := range combination {
			if operation == 0 {
				comboTotal += values[i+1]
			} else if operation == 1 {
				comboTotal *= values[i+1]
			} else if operation == 2 {
				newTotal, _ := concatenateNumbers(comboTotal, values[i+1])
				comboTotal = newTotal
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
		isSolvable := solvablePredicate(test.name, test.values, 2)
		if isSolvable {
			total += test.name
		}
	}

	return total, nil
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	tests := parseTests(input)

	total := 0

	for _, test := range tests {
		isSolvable := solvablePredicate(test.name, test.values, 3)
		if isSolvable {
			total += test.name
		}
	}

	return total, nil
}
