package day13_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day13"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day13.Parse("../../inputs/day13/sample-input.txt")

	assert.Equal(t, 480, day.Part1())
}
