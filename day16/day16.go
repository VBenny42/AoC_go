package day16

import (
	"fmt"
	"math"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type day16 struct {
	grid
	height int
	width  int
}

var opposites = map[direction]direction{
	up:    down,
	down:  up,
	left:  right,
	right: left,
}

func (d *day16) neighborsFn(cell state) []state {
	neighbors := make([]state, 3)
	i := 0
	for _, d := range []direction{up, left, down, right} {
		if d == cell.d {
			continue
		}
		// Don't ever go backwards
		if d == opposites[cell.d] {
			continue
		}
		neighbors[i] = state{x: cell.x, y: cell.y, d: d}
		i++
	}

	switch cell.d {
	case up:
		if cell.y > 0 && d.grid[cell.y-1][cell.x] != '#' {
			neighbors[2] = state{x: cell.x, y: cell.y - 1, d: cell.d}
		}
	case left:
		if cell.x > 0 && d.grid[cell.y][cell.x-1] != '#' {
			neighbors[2] = state{x: cell.x - 1, y: cell.y, d: cell.d}
		}
	case down:
		if cell.y < d.height-1 && d.grid[cell.y+1][cell.x] != '#' {
			neighbors[2] = state{x: cell.x, y: cell.y + 1, d: cell.d}
		}
	case right:
		if cell.x < d.width-1 && d.grid[cell.y][cell.x+1] != '#' {
			neighbors[2] = state{x: cell.x + 1, y: cell.y, d: cell.d}
		}
	}

	return neighbors
}

func costFn(a, b state) int {
	if a.x != b.x || a.y != b.y {
		return 1
	}
	if a.d != b.d {
		// Shouldn't be passed in, pruning from neighborsFn should prevent this
		// if (a.d == up || a.d == down) && (b.d == up || b.d == down) {
		// 	return 2000
		// }
		// if (a.d == left || a.d == right) && (b.d == left || b.d == right) {
		// 	return 2000
		// }
		return 1000
	}
	return 0
}

func (d *day16) Part1and2() (int, int) {
	var start state
	ends := make([]state, 0)

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			if d.grid[y][x] == 'S' {
				start = state{x: x, y: y, d: right}
			}
			if d.grid[y][x] == 'E' {
				ends = append(ends, state{x: x, y: y, d: up})
				ends = append(ends, state{x: x, y: y, d: right})
				ends = append(ends, state{x: x, y: y, d: down})
				ends = append(ends, state{x: x, y: y, d: left})
			}
		}
	}

	if start == (state{}) || len(ends) == 0 {
		fmt.Println("Invalid input")
		return -1, -1
	}

	dijk := dijkstra{
		neighbors: d.neighborsFn,
		cost:      costFn,
		previous:  make(map[state][]state),
		costs:     make(map[state]int),
		minCost:   0,
		maxCost:   math.MaxInt,
	}

	minCost := dijk.maxCost
	dijk.findPath(start)

	var minEnd state
	for _, end := range ends {
		cost := dijk.getCost(end)
		if cost < minCost {
			minCost = cost
			minEnd = end
		}
	}

	allTiles := make(map[struct{ x, y int }]struct{})
	for _, node := range dijk.getPaths(minEnd) {
		allTiles[struct{ x, y int }{x: node.x, y: node.y}] = struct{}{}
	}

	return minCost, len(allTiles)
}

func Parse(filename string) *day16 {
	data := utils.SplitLines(filename)

	var grid grid

	for _, line := range data {
		grid = append(grid, []rune(line))
	}

	height, width := len(grid), len(grid[0])

	return &day16{grid, height, width}
}

func Solve(filename string) {
	minCost, allTiles := Parse(filename).Part1and2()
	fmt.Println("ANSWER1: Least cost path:", minCost)
	fmt.Println("ANSWER2: All tiles on paths", allTiles)
}
