package day13

import (
	"fmt"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type point struct {
	x, y int
}

type machine struct {
	prize point
	a, b  point
}

type day13 struct {
	machines []machine
}

func (m *machine) cheapestCombination() (int, error) {
	determinant := (m.a.x * m.b.y) - (m.a.y * m.b.x)

	if determinant == 0 {
		return 0, fmt.Errorf("No solution")
	}

	a := ((m.prize.x * m.b.y) - (m.prize.y * m.b.x)) / determinant
	b := ((m.a.x * m.prize.y) - (m.a.y * m.prize.x)) / determinant

	if (m.a.x*a)+(m.b.x*b) != m.prize.x || (m.a.y*a)+(m.b.y*b) != m.prize.y {
		return 0, fmt.Errorf("No solution")
	}

	return 3*a + b, nil
}

func (d *day13) Part1() int {
	minTokens := 0
	for _, m := range d.machines {
		tokens, err := m.cheapestCombination()
		if err != nil {
			continue
		}
		minTokens += tokens
	}
	return minTokens
}

func (d *day13) Part2() int {
	minTokens := 0
	const addition = 10000000000000
	for _, m := range d.machines {
		m.prize.x += addition
		m.prize.y += addition
		tokens, err := m.cheapestCombination()
		if err != nil {
			continue
		}
		minTokens += tokens
	}
	return minTokens
}

func Parse(filename string) *day13 {
	data := utils.SplitLines(filename)

	machines := []machine{}

	var m machine

	for i := 0; i < len(data); i = i + 4 {
		fmt.Sscanf(data[i], "Button A: X+%d, Y+%d", &m.a.x, &m.a.y)
		fmt.Sscanf(data[i+1], "Button B: X+%d, Y+%d", &m.b.x, &m.b.y)
		fmt.Sscanf(data[i+2], "Prize: X=%d, Y=%d", &m.prize.x, &m.prize.y)
		machines = append(machines, m)
	}

	return &day13{machines}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: minTokens:", d.Part1())
	fmt.Println("ANSWER2: minTokens:", d.Part2())
}
