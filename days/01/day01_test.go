package day01

import (
	"testing"
)

func TestDay1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (int, error)
		output int
	}{
		{"Part 1: example", "data/example.txt", Part1, 11},
		{"Part 2: example", "data/example.txt", Part2, 31},
		{"Part 1: input", "data/input.txt", Part1, 2970687},
		{"Part 2: input", "data/input.txt", Part2, 23963899},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.part(tt.input)
			if err != nil {
				t.Errorf("%s error = %v", tt.name, err)
				return
			}
			if result != tt.output {
				t.Errorf("%s = %v, want %v", tt.name, result, tt.output)
			}

		})
	}
}
