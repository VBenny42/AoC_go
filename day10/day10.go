package day10

import (
	"fmt"
	"strconv"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type (
	grid  [][]int
	coord struct {
		x, y int
	}
)

type day10 struct {
	grid          grid
	trailheads    map[coord]int
	ninePositions []coord
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d *day10) getNextPosition(c coord, dir direction) (coord, error) {
	m, n := len(d.grid[0]), len(d.grid)
	value := d.grid[c.y][c.x]
	var next coord
	switch dir {
	case up:
		if c.y == 0 {
			return next, fmt.Errorf("Cannot move up from %v", c)
		}
		next = coord{c.x, c.y - 1}
	case down:
		if c.y == n-1 {
			return next, fmt.Errorf("Cannot move down from %v", c)
		}
		next = coord{c.x, c.y + 1}
	case left:
		if c.x == 0 {
			return next, fmt.Errorf("Cannot move left from %v", c)
		}
		next = coord{c.x - 1, c.y}
	case right:
		if c.x == m-1 {
			return next, fmt.Errorf("Cannot move right from %v", c)
		}
		next = coord{c.x + 1, c.y}
	}
	if value-d.grid[next.y][next.x] != 1 {
		return next, fmt.Errorf("Cannot move from %v to %v", c, next)
	}
	return next, nil
}

func (d *day10) findPathsToZeroOne(c coord, visited map[coord]bool) {
	if _, ok := visited[c]; ok {
		return
	}
	visited[c] = true

	if _, ok := d.trailheads[c]; ok {
		d.trailheads[c]++
		return
	}

	for _, dir := range []direction{up, down, left, right} {
		next, err := d.getNextPosition(c, dir)
		if err != nil {
			continue
		}
		d.findPathsToZeroOne(next, visited)
	}
}

func (d *day10) findPathsToZeroAll(c coord) {
	if _, ok := d.trailheads[c]; ok {
		d.trailheads[c]++
		return
	}

	for _, dir := range []direction{up, down, left, right} {
		next, err := d.getNextPosition(c, dir)
		if err != nil {
			continue
		}
		d.findPathsToZeroAll(next)
	}
}

func (d *day10) Part1() int {
	for _, nine := range d.ninePositions {
		visited := make(map[coord]bool)
		d.findPathsToZeroOne(nine, visited)
	}

	score := 0
	for _, v := range d.trailheads {
		score += v
	}

	return score
}

func (d *day10) Part2() int {
	for k := range d.trailheads {
		d.trailheads[k] = 0
	}

	for _, nine := range d.ninePositions {
		d.findPathsToZeroAll(nine)
	}

	score := 0
	for _, v := range d.trailheads {
		score += v
	}

	return score
}

func Parse(filename string) *day10 {
	data := utils.SplitLines(filename)

	grid := grid{}

	for _, line := range data {
		row := []int{}
		for _, c := range line {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	trailheads := make(map[coord]int)
	ninePositions := []coord{}

	for y, row := range grid {
		for x, n := range row {
			if n == 9 {
				ninePositions = append(ninePositions, coord{x, y})
			}
			if n == 0 {
				trailheads[coord{x, y}] = 0
			}
		}
	}

	return &day10{grid, trailheads, ninePositions}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: score:", d.Part1())
	fmt.Println("ANSWER2: score:", d.Part2())
}
