package day17_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day17"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day17.Parse("../../inputs/day17/sample-input.txt")

	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", day.Part1())
}

func TestPart2(t *testing.T) {
	day := day17.Parse("../../inputs/day17/sample-input-2.txt")

	assert.Equal(t, 117440, day.Part2())
}
