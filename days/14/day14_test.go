package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string, [2]int) (int, error)
		size   [2]int
		output int
	}{
		{"Part 1: example", "data/example.txt", Part1, [2]int{11, 7}, 12},
		// {"Part 1: input", "data/input.txt", Part1, 2358},
		// {"Part 2: example", "data/example.txt", Part2, 9},
		// {"Part 2: input", "data/input.txt", Part2, 1737},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.part(tt.input, tt.size)
			if err != nil {
				t.Errorf("%s error = %v", tt.name, err)
				return
			}
			assert.Equal(t, tt.output, result)
		})
	}
}
