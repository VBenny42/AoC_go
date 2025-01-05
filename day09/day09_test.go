package day09_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day09"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day09.Parse("../../inputs/day09/sample-input.txt")

	assert.Equal(t, 1928, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day09.Parse("../../inputs/day09/sample-input.txt")

	assert.Equal(t, 2858, day.Part2())
}
