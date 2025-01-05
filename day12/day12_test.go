package day12_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day12"
	"github.com/stretchr/testify/assert"
)

func TestPart1Smaller(t *testing.T) {
	day := day12.Parse("../../inputs/day12/sample-input-smaller.txt")

	assert.Equal(t, 140, day.Part1())
}

func TestPart1Larger(t *testing.T) {
	day := day12.Parse("../../inputs/day12/sample-input-larger.txt")

	assert.Equal(t, 1930, day.Part1())
}

func TestPart2Smaller(t *testing.T) {
	day := day12.Parse("../../inputs/day12/sample-input-smaller.txt")

	assert.Equal(t, 80, day.Part2())
}

func TestPart2Larger(t *testing.T) {
	day := day12.Parse("../../inputs/day12/sample-input-larger.txt")

	assert.Equal(t, 1206, day.Part2())
}
