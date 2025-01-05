package day03_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day03"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day03.Parse("../../inputs/day03/sample-input.txt")

	assert.Equal(t, 161, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day03.Parse("../../inputs/day03/sample-input-1.txt")
	day.Part1()

	assert.Equal(t, 48, day.Part2())
}
