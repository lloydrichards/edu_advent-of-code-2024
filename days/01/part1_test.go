package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"data/example.txt", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Part1(tt.input)
			if err != nil {
				t.Errorf("Part1() error = %v", err)
				return
			}
			if result != tt.output {
				t.Errorf("Part1() = %v, want %v", result, tt.output)
			}

		})
	}
}
