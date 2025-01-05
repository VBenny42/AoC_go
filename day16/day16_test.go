package day16_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day16"
	"github.com/stretchr/testify/assert"
)

func TestPart1Smaller(t *testing.T) {
	day := day16.Parse("../../inputs/day16/sample-input-smaller.txt")

	value, _ := day.Part1and2()

	assert.Equal(t, 7036, value)
}

func TestPart1(t *testing.T) {
	day := day16.Parse("../../inputs/day16/sample-input.txt")

	value, _ := day.Part1and2()

	assert.Equal(t, 11048, value)
}

func TestPart2(t *testing.T) {
	day := day16.Parse("../../inputs/day16/sample-input.txt")

	_, value := day.Part1and2()

	assert.Equal(t, 64, value)
}
