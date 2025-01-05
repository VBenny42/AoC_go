package day20

import (
	"fmt"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type (
	grid  [][]rune
	coord struct{ x, y int }
)

type day20 struct {
	grid
}

func isValid(g grid, c coord) bool {
	return c.y >= 0 && c.y < len(g) && c.x >= 0 && c.x < len(g[0]) && g[c.y][c.x] != '#'
}

func bfs(g grid, start, end coord) []coord {
	type state struct {
		c coord
		p []coord
	}

	var q []state

	enqueue := func(next coord, path []coord) {
		newPath := make([]coord, len(path)+1)
		copy(newPath, path)
		newPath[len(path)] = next
		q = append(q, state{next, newPath})
	}
	dequeue := func() state {
		s := q[0]
		q = q[1:]
		return s
	}

	visited := make(map[coord]struct{})

	enqueue(start, []coord{})

	for len(q) > 0 {
		curr := dequeue()

		if curr.c == end {
			return curr.p
		}

		if _, ok := visited[curr.c]; ok {
			continue
		}
		visited[curr.c] = struct{}{}

		for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := coord{curr.c.x + dir[0], curr.c.y + dir[1]}
			if !isValid(g, next) {
				continue
			}
			enqueue(next, curr.p)
		}
	}

	return nil
}

func getCoord(g grid, c rune) coord {
	for y, row := range g {
		for x, v := range row {
			if v == c {
				return coord{x, y}
			}
		}
	}
	return coord{}
}

func manhattanDistance(c1, c2 coord) int {
	return utils.Abs(c1.x-c2.x) + utils.Abs(c1.y-c2.y)
}

func (d *day20) part1and2() {
	start := getCoord(d.grid, 'S')
	end := getCoord(d.grid, 'E')

	path := bfs(d.grid, start, end)

	threshold := 100

	twoCheats, twentyCheats := 0, 0
	for i := 0; i < len(path)-threshold; i++ {
		for j := i + threshold; j < len(path); j++ {
			cheatDuration := manhattanDistance(path[i], path[j])
			if cheatDuration <= 2 && ((j-i)-cheatDuration >= threshold) {
				twoCheats++
			}
			if cheatDuration <= 20 && ((j-i)-cheatDuration >= threshold) {
				twentyCheats++
			}
		}
	}
	fmt.Println("ANSWER1: twoCheats:", twoCheats)
	fmt.Println("ANSWER2: twentyCheats:", twentyCheats)
}

func parse(filename string) *day20 {
	data := utils.SplitLines(filename)

	var grid grid

	for _, line := range data {
		grid = append(grid, []rune(line))
	}

	return &day20{grid}
}

func Solve(filename string) {
	parse(filename).part1and2()
}
