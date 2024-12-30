package main

import (
	"fmt"
	"os"
	"strings"
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

func (d *day13) part1() {
	minTokens := 0
	for _, m := range d.machines {
		tokens, err := m.cheapestCombination()
		if err != nil {
			continue
		}
		minTokens += tokens
	}
	fmt.Println("ANSWER1: minTokens:", minTokens)
}

func (d *day13) part2() {
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
	fmt.Println("ANSWER2: minTokens:", minTokens)
}

func parse() *day13 {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	lines := strings.Split(strings.Trim(string(file), "\n"), "\n")

	machines := []machine{}

	var m machine

	for i := 0; i < len(lines); i = i + 4 {
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &m.a.x, &m.a.y)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &m.b.x, &m.b.y)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &m.prize.x, &m.prize.y)
		machines = append(machines, m)
	}

	return &day13{machines}
}

func main() {
	d := parse()
	d.part1()
	d.part2()
}
