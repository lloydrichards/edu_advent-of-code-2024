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

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	code, err := parseCode(input)
	if err != nil {
		return -1, err
	}

	total := 0

	for _, c := range code {
		total += c[0] * c[1]
	}

	return total, nil
}

func Part2(dir string) (int, error) {

	return -1, nil
}
