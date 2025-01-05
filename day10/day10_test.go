package day10_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day10"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day10.Parse("../../inputs/day10/sample-input.txt")

	assert.Equal(t, 36, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day10.Parse("../../inputs/day10/sample-input.txt")

	assert.Equal(t, 81, day.Part2())
}
