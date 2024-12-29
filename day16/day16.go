package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type day16 struct {
	grid
	height int
	width  int
}

func (d *day16) neighborsFn(cell state) []state {
	neighbors := make([]state, 0)
	for _, d := range []direction{UP, LEFT, DOWN, RIGHT} {
		if d == cell.d {
			continue
		}
		neighbors = append(neighbors, state{x: cell.x, y: cell.y, d: d})
	}

	switch cell.d {
	case UP:
		if cell.y > 0 && d.grid[cell.y-1][cell.x] != '#' {
			neighbors = append(neighbors, state{x: cell.x, y: cell.y - 1, d: cell.d})
		}
	case LEFT:
		if cell.x > 0 && d.grid[cell.y][cell.x-1] != '#' {
			neighbors = append(neighbors, state{x: cell.x - 1, y: cell.y, d: cell.d})
		}
	case DOWN:
		if cell.y < d.height-1 && d.grid[cell.y+1][cell.x] != '#' {
			neighbors = append(neighbors, state{x: cell.x, y: cell.y + 1, d: cell.d})
		}
	case RIGHT:
		if cell.x < d.width-1 && d.grid[cell.y][cell.x+1] != '#' {
			neighbors = append(neighbors, state{x: cell.x + 1, y: cell.y, d: cell.d})
		}
	}
	return neighbors
}

func costFn(a, b state) int {
	if a.x != b.x || a.y != b.y {
		return 1
	}
	if a.d != b.d {
		if (a.d == UP || a.d == DOWN) && (b.d == UP || b.d == DOWN) {
			return 2000
		}
		if (a.d == LEFT || a.d == RIGHT) && (b.d == LEFT || b.d == RIGHT) {
			return 2000
		}
		return 1000
	}
	return 0
}

func (d *day16) part1and2() {
	var start state
	ends := make([]state, 0)

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			if d.grid[y][x] == 'S' {
				start = state{x: x, y: y, d: UP}
			}
			if d.grid[y][x] == 'E' {
				ends = append(ends, state{x: x, y: y, d: UP})
				ends = append(ends, state{x: x, y: y, d: RIGHT})
				ends = append(ends, state{x: x, y: y, d: DOWN})
				ends = append(ends, state{x: x, y: y, d: LEFT})
			}
		}
	}

	if start == (state{}) || len(ends) == 0 {
		fmt.Println("Invalid input")
		return
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

	fmt.Println("ANSWER1: Least cost path:", minCost)
	fmt.Println("ANSWER2: All tiles on paths", len(allTiles))
}

func solve() *day16 {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return nil
	}
	defer file.Close()

	var grid grid

	s := bufio.NewScanner(file)
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}

	height, width := len(grid), len(grid[0])

	return &day16{grid, height, width}
}

func main() {
	solve().part1and2()
}
