package day07_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day07"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day07.Parse("../../inputs/day07/sample-input.txt")

	assert.Equal(t, 3749, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day07.Parse("../../inputs/day07/sample-input.txt")

	assert.Equal(t, 11387, day.Part2())
}

func BenchmarkPart1(b *testing.B) {
	day := day07.Parse("../../inputs/day07/input.txt")

	// Reset timer after setup
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		day.Part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	day := day07.Parse("../../inputs/day07/input.txt")

	// Reset timer after setup
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		day.Part2()
	}
}
