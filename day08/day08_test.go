package day08_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day08"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day08.Parse("../../inputs/day08/sample-input.txt")

	assert.Equal(t, 14, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day08.Parse("../../inputs/day08/sample-input.txt")

	assert.Equal(t, 34, day.Part2())
}
