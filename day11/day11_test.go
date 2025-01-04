package day11_test

import (
	"testing"

	"github.com/VBenny42/AoC_go/day11"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day11.Parse("../inputs/day11/sample-input.txt")

	assert.Equal(t, 55312, day.Part1())
}
