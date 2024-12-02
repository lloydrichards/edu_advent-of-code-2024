package utils

import (
	"os"
)

func LoadInputFile(path string) (string, error) {
	buff, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}
