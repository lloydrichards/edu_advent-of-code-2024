package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example", "data/example_1.txt", Part1, 140},
		{"Part 1: example", "data/example_2.txt", Part1, 772},
		{"Part 1: example", "data/example_3.txt", Part1, 1930},
		// {"Part 1: input", "data/input.txt", Part1, 1363682}, <- expensive
		{"Part 2: example", "data/example_1.txt", Part2, 80},
		{"Part 2: example", "data/example_2.txt", Part2, 436},
		{"Part 2: example", "data/example_3.txt", Part2, 1206},
		{"Part 2: example", "data/example_4.txt", Part2, 236},
		{"Part 2: example", "data/example_5.txt", Part2, 368},
		// {"Part 2: input", "data/input.txt", Part2, 787680}, <- expensive
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
