package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example_1", "data/example_1.txt", Part1, 7036},
		{"Part 1: example_2", "data/example_2.txt", Part1, 11048},
		// {"Part 1: input", "data/input.txt", Part1, 2358},
		// {"Part 2: example", "data/example.txt", Part2, 9},
		// {"Part 2: input", "data/input.txt", Part2, 1737},
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
