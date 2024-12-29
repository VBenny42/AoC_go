package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func neighborsFn(cell state, g grid, width, height int) []state {
	neighbors := make([]state, 0)
	for _, d := range []direction{UP, LEFT, DOWN, RIGHT} {
		if d == cell.d {
			continue
		}
		neighbors = append(neighbors, state{x: cell.x, y: cell.y, d: d})
	}

	switch cell.d {
	case UP:
		if cell.y > 0 && g[cell.y-1][cell.x] != '#' {
			neighbors = append(neighbors, state{x: cell.x, y: cell.y - 1, d: cell.d})
		}
	case LEFT:
		if cell.x > 0 && g[cell.y][cell.x-1] != '#' {
			neighbors = append(neighbors, state{x: cell.x - 1, y: cell.y, d: cell.d})
		}
	case DOWN:
		if cell.y < height-1 && g[cell.y+1][cell.x] != '#' {
			neighbors = append(neighbors, state{x: cell.x, y: cell.y + 1, d: cell.d})
		}
	case RIGHT:
		if cell.x < width-1 && g[cell.y][cell.x+1] != '#' {
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

func part1and2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}
	defer file.Close()

	var grid grid

	s := bufio.NewScanner(file)
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}

	height, width := len(grid), len(grid[0])

	var start state
	ends := make([]state, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'S' {
				start = state{x: x, y: y, d: UP}
			}
			if grid[y][x] == 'E' {
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

	n := func(cell state) []state {
		return neighborsFn(cell, grid, width, height)
	}

	d := dijkstra{
		neighbors: n,
		cost:      costFn,
		previous:  make(map[state][]state),
		costs:     make(map[state]int),
		minCost:   0,
		maxCost:   math.MaxInt,
	}

	minCost := d.maxCost
	d.findPath(start)
	var minEnd state
	for _, end := range ends {
		cost := d.getCost(end)
		if cost < minCost {
			minCost = cost
			minEnd = end
		}
	}

	tilesAll := make(map[struct{ x, y int }]struct{})
	for _, node := range d.getPaths(minEnd) {
		tilesAll[struct{ x, y int }{x: node.x, y: node.y}] = struct{}{}
	}

	fmt.Println("ANSWER1: Least cost path:", minCost)
	fmt.Println("ANSWER2: Tiles on all paths", len(tilesAll))
}

func main() {
	part1and2()
}
