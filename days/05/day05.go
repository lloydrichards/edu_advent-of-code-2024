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

func eitherPages(rules []Rules, pages []Pages) ([]Pages, []Pages) {

	validPages := []Pages{}
	invalidPages := []Pages{}

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
		if !isValid {
			invalidPages = append(invalidPages, list)
		}
	}

	return validPages, invalidPages
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	rules, pages := parseInput(input)

	valid, _ := eitherPages(rules, pages)

	total := 0
	for _, page := range valid {
		total += middleArray(page)
	}

	return total, nil
}

func shouldSwap(i, j int, rules []Rules) bool {
	for _, rule := range rules {
		if rule.Min == j && rule.Max == i {
			return true
		}
	}
	return false
}

func contains(pages Pages, page int) bool {
	for _, p := range pages {
		if p == page {
			return true
		}
	}
	return false
}

func sortPages(pages Pages, rules []Rules) Pages {
	filteredRules := []Rules{}
	for _, rule := range rules {
		if contains(pages, rule.Min) && contains(pages, rule.Max) {
			filteredRules = append(filteredRules, rule)
		}
	}

	sorted := make([]int, len(pages))
	copy(sorted, pages)

	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if shouldSwap(sorted[j], sorted[j+1], filteredRules) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}
	rules, pages := parseInput(input)

	_, brokenPages := eitherPages(rules, pages)

	sortedPages := []Pages{}
	for _, page := range brokenPages {
		sorted := sortPages(page, rules)
		sortedPages = append(sortedPages, sorted)
	}

	total := 0
	for _, page := range sortedPages {
		total += middleArray(page)
	}

	return total, nil
}
