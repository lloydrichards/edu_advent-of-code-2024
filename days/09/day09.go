package day00

import (
	U "advent-of-code-2024/internal/utils"
	"strconv"
	"strings"
)

func decompress(s string) []string {
	output := []string{}
	id := 0
	for i, part := range strings.Split(s, "") {
		num, _ := strconv.Atoi(part)
		isEven := i%2 == 0
		for j := 0; j < num; j++ {
			if isEven {
				output = append(output, strconv.Itoa(id))
			} else {
				output = append(output, ".")
			}

		}
		if isEven {
			id++
		}
	}
	return output
}

func findIdxOfFirstDot(filesystem []string) int {
	for i, part := range filesystem {
		if part == "." {
			return i
		}
	}
	return -1
}

func compress(fs []string) []string {
	for i := len(fs) - 1; i > -1; i-- {
		if fs[i] != "." {
			firstIdx := findIdxOfFirstDot(fs)
			if firstIdx > i {
				break
			}
			fs[firstIdx], fs[i] = fs[i], fs[firstIdx]
		}
	}
	return fs
}

func calcChecksum(fs []string) []int {
	checksum := []int{}
	for i := 0; i < len(fs); i++ {

		if fs[i] != "." {
			num, _ := strconv.Atoi(fs[i])
			checksum = append(checksum, i*num)
		}
	}
	return checksum
}

func sum(checksum []int) int {
	total := 0
	for _, num := range checksum {
		total += num
	}
	return total
}

func Part1(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	fs := decompress(input)
	compressedFs := compress(fs)
	checksum := calcChecksum(compressedFs)

	return sum(checksum), nil
}

func Part2(dir string) (int, error) {
	// input, err := U.LoadInputFile(dir)
	// if err != nil {
	// 	return -1, err
	// }

	return -1, nil
}
