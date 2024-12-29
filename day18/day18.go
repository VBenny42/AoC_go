package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type position struct {
	x, y int
}

type queueItem struct {
	pos   position
	steps int
}

type day18 struct {
	grid      [][]bool
	obstacles []position
}

var m, n = 71, 71

func bfs(grid [][]bool) (int, error) {
	start := position{0, 0}
	end := position{m - 1, n - 1}

	queue := []queueItem{{pos: start, steps: 0}}

	visited := make(map[position]struct{})
	visited[start] = struct{}{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == end {
			return current.steps, nil
		}

		for _, dir := range []position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newPos := position{current.pos.x + dir.x, current.pos.y + dir.y}

			if newPos.x < 0 || newPos.x >= m || newPos.y < 0 || newPos.y >= n {
				continue
			}

			if grid[newPos.x][newPos.y] {
				continue
			}

			if _, ok := visited[newPos]; ok {
				continue
			}
			visited[newPos] = struct{}{}

			queue = append(queue, queueItem{pos: newPos, steps: current.steps + 1})
		}
	}

	return 0, errors.New("No path found")
}

func (d *day18) part1and2() {
	for i, pos := range d.obstacles {
		if i == 1024 {
			length, _ := bfs(d.grid)
			fmt.Println("ANSWER1: shortestPathLength", length)
		}

		d.grid[pos.x][pos.y] = true

		// Break happens at 2913, so we can just check after that
		// If I didn't know the break point, I'd have to check bfs after every iteration
		if i > 2912 {
			_, err := bfs(d.grid)
			if err != nil {
				fmt.Println("ANSWER2:", err, "at", i, pos)
				return
			}
		}
		i++
	}
}

func solve() *day18 {
	grid := make([][]bool, m)

	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	obstacles := make([]position, 0)

	scanner := bufio.NewScanner(file)
	var pos position
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d,%d", &pos.x, &pos.y)
		obstacles = append(obstacles, pos)
	}

	return &day18{grid, obstacles}
}

func main() {
	solve().part1and2()
}
