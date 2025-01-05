package day01_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day01"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day01.Parse("../../inputs/day01/sample-input.txt")

	assert.Equal(t, 11, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day01.Parse("../../inputs/day01/sample-input.txt")

	assert.Equal(t, 31, day.Part2())
}
