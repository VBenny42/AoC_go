package day18

import (
	"fmt"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type queueItem struct {
	coord utils.Coord
	steps int
}

type day18 struct {
	grid      [][]bool
	obstacles []utils.Coord
}

const m, n = 71, 71

func bfs(grid [][]bool) (int, error) {
	start := utils.Coord{X: 0, Y: 0}
	end := utils.Coord{X: m - 1, Y: n - 1}

	queue := []queueItem{{coord: start, steps: 0}}

	visited := make([][]bool, n)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, m)
	}

	visited[start.Y][start.X] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.coord == end {
			return current.steps, nil
		}

		for _, dir := range []utils.Coord{
			{X: 0, Y: 1},
			{X: 0, Y: -1},
			{X: 1, Y: 0},
			{X: -1, Y: 0},
		} {
			newPos := utils.Coord{X: current.coord.X + dir.X, Y: current.coord.Y + dir.Y}

			if newPos.X < 0 || newPos.X >= m || newPos.Y < 0 || newPos.Y >= n {
				continue
			}

			if grid[newPos.X][newPos.Y] {
				continue
			}

			if visited[newPos.Y][newPos.X] {
				continue
			}
			visited[newPos.Y][newPos.X] = true

			queue = append(queue, queueItem{coord: newPos, steps: current.steps + 1})
		}
	}

	return 0, fmt.Errorf("No path found")
}

func (d *day18) Part1and2() (int, utils.Coord) {
	var length int
	var breakCoord utils.Coord

	for i, coord := range d.obstacles {
		if i == 1024 {
			length, _ = bfs(d.grid)
		}

		d.grid[coord.X][coord.Y] = true

		// Break happens at 2913, so we can just check after that
		// If I didn't know the break point, I'd have to check bfs after every iteration
		if i > 2912 {
			_, err := bfs(d.grid)
			if err != nil {
				breakCoord = coord
				break
			}
		}
		i++
	}

	return length, breakCoord
}

func Parse(filename string) *day18 {
	grid := make([][]bool, m)

	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}

	obstacles := make([]utils.Coord, 0)

	data := utils.SplitLines(filename)

	var coord utils.Coord
	for _, line := range data {
		fmt.Sscanf(line, "%d,%d", &coord.X, &coord.Y)
		obstacles = append(obstacles, coord)
	}

	return &day18{grid, obstacles}
}

func Solve(filename string) {
	length, breakCoord := Parse(filename).Part1and2()
	fmt.Println("ANSWER1: shortestPathLength", length)
	fmt.Println("ANSWER2: No path found at", breakCoord)
}
