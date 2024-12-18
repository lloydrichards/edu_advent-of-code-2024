package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example", "data/example_1.txt", Part1, 2028},
		{"Part 1: example", "data/example_2.txt", Part1, 10092},
		{"Part 1: input", "data/input.txt", Part1, 1505963},
		// {"Part 2: example", "data/example_3.txt", Part2, 9021},
		{"Part 2: example", "data/example_2.txt", Part2, 9021},
		{"Part 2: input", "data/input.txt", Part2, 1543141},
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
