package main

import (
	"fmt"
	"os"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type (
	grid  [][]rune
	coord struct {
		x, y int
	}
)

type day08 struct {
	grid        grid
	frequencies map[rune]map[coord]bool
	antinodes   map[coord]bool
}

func (g grid) inBounds(c coord) bool {
	return 0 <= c.x && c.x < len((g)[0]) && 0 <= c.y && c.y < len(g)
}

func (d *day08) findCloseAntinodes(frequency map[coord]bool) {
	kIndices := make([]coord, len(frequency))
	i := 0
	for k := range frequency {
		kIndices[i] = k
		i++
	}

	combinations := make([]int, 2)

	gen := combin.NewCombinationGenerator(len(frequency), 2)
	for gen.Next() {
		gen.Combination(combinations)

		f1, f2 := kIndices[combinations[0]], kIndices[combinations[1]]

		dX, dY := f2.x-f1.x, f2.y-f1.y

		posX, posY := f1.x-dX, f1.y-dY
		negX, negY := f2.x+dX, f2.y+dY

		if d.grid.inBounds(coord{posX, posY}) {
			d.antinodes[coord{posX, posY}] = true
		}
		if d.grid.inBounds(coord{negX, negY}) {
			d.antinodes[coord{negX, negY}] = true
		}
	}
}

func (d *day08) findAllAntinodes(frequency map[coord]bool) {
	kIndices := make([]coord, len(frequency))
	i := 0
	for k := range frequency {
		kIndices[i] = k
		i++
	}

	combinations := make([]int, 2)

	gen := combin.NewCombinationGenerator(len(frequency), 2)
	for gen.Next() {
		gen.Combination(combinations)

		f1, f2 := kIndices[combinations[0]], kIndices[combinations[1]]

		dX, dY := f2.x-f1.x, f2.y-f1.y

		startX, startY := f1.x, f1.y
		for d.grid.inBounds(coord{startX, startY}) {
			d.antinodes[coord{startX, startY}] = true
			startX += dX
			startY += dY
		}

		startX, startY = f2.x, f2.y
		for d.grid.inBounds(coord{startX, startY}) {
			d.antinodes[coord{startX, startY}] = true
			startX -= dX
			startY -= dY
		}

	}
}

func (d *day08) part1() {
	for _, v := range d.frequencies {
		d.findCloseAntinodes(v)
	}

	fmt.Println("ANSWER1: unique antinodes:", len(d.antinodes))
}

func (d *day08) part2() {
	for _, v := range d.frequencies {
		d.findAllAntinodes(v)
	}

	fmt.Println("ANSWER2: unique antinodes:", len(d.antinodes))
}

func parse() *day08 {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := grid{}
	frequencies := map[rune]map[coord]bool{}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	for y, line := range lines {
		grid = append(grid, []rune(line))
		for x, r := range line {
			if r != '.' {
				if frequencies[r] == nil {
					frequencies[r] = map[coord]bool{}
				}
				frequencies[r][coord{x, y}] = true
			}
		}
	}

	antinodes := map[coord]bool{}

	return &day08{grid, frequencies, antinodes}
}

func main() {
	d := parse()
	d.part1()
	d.part2()
}
