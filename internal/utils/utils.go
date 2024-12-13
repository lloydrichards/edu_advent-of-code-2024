package utils

import (
	"os"
	"strings"
)

func LoadInputFile(path string) (string, error) {
	buff, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func SplitByNewline(input string) []string {
	return strings.Split(input, "\n")
}

func GetGridSize(input string) (int, int) {
	lines := SplitByNewline(input)
	return len(strings.Split(lines[0], "")), len(lines)
}
