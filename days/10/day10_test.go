package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example", "data/example.txt", Part1, 36},
		{"Part 1: example", "data/example_2.txt", Part1, 2},
		{"Part 1: input", "data/input.txt", Part1, 501},
		{"Part 2: example", "data/example.txt", Part2, 81},
		{"Part 2: input", "data/input.txt", Part2, 1017},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.part(tt.input)
			if err != nil {
				t.Errorf("%s error = %v", tt.name, err)
				return
			}
			assert.Equal(t, tt.output, result)
		})
	}
}
