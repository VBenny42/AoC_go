package day06_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day06"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day06.Parse("../../inputs/day06/sample-input.txt")

	assert.Equal(t, 41, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day06.Parse("../../inputs/day06/sample-input.txt")

	day.Part1()

	assert.Equal(t, 6, day.Part2())
}
