package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example", "data/example.txt", Part1, 161},
		{"Part 1: input", "data/input.txt", Part1, 190604937},
		{"Part 1: example", "data/example-2.txt", Part2, 48},
		{"Part 1: input", "data/input.txt", Part2, 82857512},
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
