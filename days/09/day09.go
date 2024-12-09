package day09

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

func blockDecompress(s string) [][]string {
	output := [][]string{}
	id := 0
	for i, part := range strings.Split(s, "") {
		num, _ := strconv.Atoi(part)
		isEven := i%2 == 0
		block := []string{}
		for j := 0; j < num; j++ {
			if isEven {
				block = append(block, strconv.Itoa(id))
			} else {
				block = append(block, ".")
			}

		}
		if isEven {
			id++
		}
		output = append(output, block)
	}
	return output
}
func findIdxOfFirstDotBlock(filesystem [][]string, size int) int {
	for i, part := range filesystem {
		if len(part) >= size && part[0] == "." {
			return i
		}
	}
	return -1
}

func groupBlocks(fs []string) [][]string {
	output := [][]string{}
	blockIdx := 0
	curChar := fs[0]
	for i := 0; i < len(fs); i++ {
		if fs[i] != curChar {
			blockIdx++
			curChar = fs[i]
		}
		if len(output) <= blockIdx {
			output = append(output, []string{})
		}
		output[blockIdx] = append(output[blockIdx], fs[i])
	}

	return output

}

func swapBlocks(fs [][]string, a int, b int) [][]string {
	if len(fs[a]) >= len(fs[b]) {
		for i := 0; i < len(fs[b]); i++ {
			fs[a][i], fs[b][i] = fs[b][i], fs[a][i]
		}
	}
	return fs
}

func flatten(fs [][]string) []string {
	output := []string{}
	for _, block := range fs {
		output = append(output, block...)
	}
	return output
}

func findIdx(fs [][]string, id string) int {
	for i, part := range fs {
		if len(part) > 0 && part[0] == id {
			return i
		}
	}
	return -1
}

func findLastNum(fs [][]string) int {
	for i := len(fs) - 1; i >= 0; i-- {
		if len(fs[i]) > 0 && fs[i][0] != "." {
			num, _ := strconv.Atoi(fs[i][0])
			return num
		}
	}
	return 0
}

func blockCompress(fs [][]string) [][]string {
	for i := findLastNum(fs); i >= 0; i-- {
		curPartIdx := findIdx(fs, strconv.Itoa(i))
		firstIdx := findIdxOfFirstDotBlock(fs, len(fs[curPartIdx]))
		if firstIdx == -1 {
			continue
		}

		if firstIdx > curPartIdx {
			continue
		}

		fs = swapBlocks(fs, firstIdx, curPartIdx)
		fs = groupBlocks(flatten(fs))
	}
	return fs
}

func Part2(dir string) (int, error) {
	input, err := U.LoadInputFile(dir)
	if err != nil {
		return -1, err
	}

	fs := blockDecompress(input)
	compressedFs := blockCompress(fs)
	flatFs := flatten(compressedFs)
	checksum := calcChecksum(flatFs)

	return sum(checksum), nil
}
