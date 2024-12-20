package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example", "data/example.txt", Part1, 14},
		{"Part 1: input", "data/input.txt", Part1, 271},
		{"Part 2: example", "data/example.txt", Part2, 34},
		{"Part 2: input", "data/input.txt", Part2, 994},
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
