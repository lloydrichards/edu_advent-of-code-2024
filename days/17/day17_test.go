package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		part   func(string) (string, error)
		output string
	}{
		{"Part 1: example", "data/example.txt", Part1, "4,6,3,5,6,3,5,2,1,0"},
		{"Part 1: input", "data/input.txt", Part1, "7,5,4,3,4,5,3,4,6"},
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

	t.Run("Part 2: input", func(t *testing.T) {
		result, err := Part2("data/input.txt")
		if err != nil {
			t.Errorf("%s error = %v", "Part 2: input", err)
			return
		}
		assert.Equal(t, 164278899142333, result)
	})
}
