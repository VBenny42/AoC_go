package day04_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day04"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day04.Parse("../../inputs/day04/sample-input.txt")

	assert.Equal(t, 18, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day04.Parse("../../inputs/day04/sample-input.txt")

	assert.Equal(t, 9, day.Part2())
}
