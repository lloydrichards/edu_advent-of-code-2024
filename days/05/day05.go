package day05

import (
	U "advent-of-code-2024/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type Rules struct {
	Min int
	Max int
}

type Pages []int

func parseInput(input string) ([]Rules, []Pages) {
	split := strings.Split(input, "\n\n")
	rules := []Rules{}
	pages := []Pages{}

	for _, rule := range strings.Split(split[0], "\n") {
		parts := strings.Split(rule, "|")

		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		rules = append(rules, Rules{Min: min, Max: max})
	}

	for _, page := range strings.Split(split[1], "\n") {
		parts := strings.Split(page, ",")
		p := []int{}
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			p = append(p, num)
		}
		pages = append(pages, p)
	}

	return rules, pages
}

func mapPageToIdx(pages Pages) map[int]int {
	m := map[int]int{}
	for i, page := range pages {
		m[page] = i
	}
	return m
}

func mapIdxToRules(rule Rules, mapper map[int]int) (Rules, error) {
	newMin, okMin := mapper[rule.Min]
	if !okMin {
		return Rules{}, fmt.Errorf("missing")
	}
	newMax, okMax := mapper[rule.Max]
	if !okMax {
		return Rules{}, fmt.Errorf("missing")
	}

	return Rules{Min: newMin, Max: newMax}, nil
}

func middleArray(arr []int) int {
	return arr[len(arr)/2]
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	rules, pages := parseInput(input)

	validPages := []Pages{}

	for _, list := range pages {
		pageIdxs := mapPageToIdx(list)
		isValid := true
		for _, rule := range rules {
			mappedRule, err := mapIdxToRules(rule, pageIdxs)
			if err != nil {
				continue
			}
			if mappedRule.Min > mappedRule.Max {
				isValid = false
				break
			}
		}
		if isValid {
			validPages = append(validPages, list)
		}
	}
	total := 0
	for _, page := range validPages {
		total += middleArray(page)
	}

	return total, nil
}

func Part2(dir string) (int, error) {
	_, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	return -1, nil
}
