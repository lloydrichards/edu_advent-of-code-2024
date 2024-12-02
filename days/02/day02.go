package day02

import (
	U "advent-of-code-2024/internal/utils"
	"math"
	"strconv"
	"strings"
)

func parseReports(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	reports := [][]int{}
	for _, line := range lines {
		report := []int{}
		for _, levelStr := range strings.Fields(line) {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				return nil, err
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func isValidReport(report []int) bool {
	dir := "asc"

	for i := 0; i < len(report); i++ {
		if i == 0 {
			if report[i] > report[i+1] {
				dir = "desc"
			} else {
				dir = "asc"
			}
			continue
		}

		if dir == "asc" {
			if report[i] < report[i-1] {
				return false
			}
		} else {
			if report[i] > report[i-1] {
				return false
			}
		}
		if report[i] == report[i-1] {
			return false
		}
		if math.Abs(float64(report[i]-report[i-1])) > 3 {
			return false
		}

	}
	return true
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	reports, err := parseReports(input)
	if err != nil {
		return -1, err
	}

	validReports := 0
	for _, report := range reports {
		if isValidReport(report) {
			validReports++
		}
	}
	return validReports, nil
}

func Part2(dir string) (int, error) {
	return -1, nil
}
