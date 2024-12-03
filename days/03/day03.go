package day03

import (
	U "advent-of-code-2024/internal/utils"
	"regexp"
	"strconv"
	"strings"
)

func parseCode(input string) ([][]int, error) {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	found := re.FindAllString(input, -1)

	result := [][]int{}
	for _, f := range found {
		f = strings.ReplaceAll(f, "mul(", "")
		f = strings.ReplaceAll(f, ")", "")
		parts := strings.Split(f, ",")
		nums := []int{}
		for _, p := range parts {
			num, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		result = append(result, nums)

	}

	return result, nil
}

func totalCodes(code [][]int) int {
	total := 0

	for _, c := range code {
		total += c[0] * c[1]
	}

	return total
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	codes, err := parseCode(input)
	if err != nil {
		return -1, err
	}

	total := totalCodes(codes)

	return total, nil
}

func removeCode(input string) string {
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	input = strings.ReplaceAll(input, "\n", "")
	return re.ReplaceAllString(input, "")
}

func Part2(dir string) (int, error) {

	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	cleaned := removeCode(input)
	codes, err := parseCode(cleaned)
	if err != nil {
		return -1, err
	}

	total := totalCodes(codes)

	return total, nil
}
