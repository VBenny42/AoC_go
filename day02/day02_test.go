package day02_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day02"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day02.Parse("../../inputs/day02/sample-input.txt")

	assert.Equal(t, 2, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day02.Parse("../../inputs/day02/sample-input.txt")

	assert.Equal(t, 4, day.Part2())
}
