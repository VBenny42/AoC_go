package day14_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day14"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day14.Parse("../../inputs/day14/sample-input.txt")

	assert.Equal(t, 21, day.Part1())
}
