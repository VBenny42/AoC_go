package day15_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day15"
	"github.com/stretchr/testify/assert"
)

func TestPart1Smaller(t *testing.T) {
	day := day15.Parse("../../inputs/day15/sample-input-smaller.txt")

	assert.Equal(t, 2028, day.Part1())
}

func TestPart1Larger(t *testing.T) {
	day := day15.Parse("../../inputs/day15/sample-input-larger.txt")

	assert.Equal(t, 10092, day.Part1())
}

func TestPart2Larger(t *testing.T) {
	day := day15.Parse("../../inputs/day15/sample-input-larger.txt")

	assert.Equal(t, 9021, day.Part2())
}
