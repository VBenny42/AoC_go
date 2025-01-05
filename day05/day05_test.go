package day05_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day05"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day05.Parse("../../inputs/day05/sample-input.txt")

	value, _ := day.Part1and2()

	assert.Equal(t, 143, value)
}

func TestPart2(t *testing.T) {
	day := day05.Parse("../../inputs/day05/sample-input.txt")

	_, value := day.Part1and2()

	assert.Equal(t, 123, value)
}
