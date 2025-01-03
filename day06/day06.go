package day06

import (
	"fmt"
	"sync"

	"github.com/VBenny42/AoC_go/utils"
)

type coord struct {
	x, y int
}
type (
	visitedDirection [5]bool
	grid             [][]visitedDirection
)

type day06 struct {
	grid              grid
	start             coord
	distinctPositions []coord
}

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
	box   = 4
)

var rotateMap = map[int]int{
	up:    right,
	right: down,
	down:  left,
	left:  up,
}

func (g grid) getNextPosition(c coord, d int) (coord, error) {
	m, n := len(g[0]), len(g)
	var next coord
	switch d {
	case up:
		if c.y == 0 {
			return coord{}, fmt.Errorf("can't move up")
		}
		next = coord{c.x, c.y - 1}
	case right:
		if c.x == m-1 {
			return coord{}, fmt.Errorf("can't move right")
		}
		next = coord{c.x + 1, c.y}
	case down:
		if c.y == n-1 {
			return coord{}, fmt.Errorf("can't move down")
		}
		next = coord{c.x, c.y + 1}
	case left:
		if c.x == 0 {
			return coord{}, fmt.Errorf("can't move left")
		}
		next = coord{c.x - 1, c.y}
	}
	return next, nil
}

func (g grid) markVisited(c coord) {
	currentDirection := up
	current := c
	for {
		next, err := g.getNextPosition(current, currentDirection)
		if err != nil {
			g[current.y][current.x][currentDirection] = true
			break
		}
		g[current.y][current.x][currentDirection] = true
		if g[next.y][next.x][box] {
			currentDirection = rotateMap[currentDirection]
			continue
		}
		current = next
	}
}

func (d *day06) doesInduceLoop(obstruction coord) bool {
	visited := make(map[coord][4]bool)
	currentDirection := up
	current := d.start
	for {
		next, err := d.grid.getNextPosition(current, currentDirection)
		if err != nil {
			return false
		}
		if next == obstruction || d.grid[next.y][next.x][box] {
			currentDirection = rotateMap[currentDirection]
			continue
		}
		if visited[next][currentDirection] {
			return true
		}
		directions := visited[next]
		directions[currentDirection] = true
		visited[next] = directions
		current = next
	}
}

func (d *day06) part1() {
	d.grid.markVisited(d.start)
	distinctPositions := make([]coord, 0)
	for y, row := range d.grid {
		for x, cell := range row {
			if cell[up] || cell[right] || cell[down] || cell[left] {
				distinctPositions = append(distinctPositions, coord{x, y})
			}
		}
	}
	d.distinctPositions = distinctPositions
	fmt.Println("ANSWER1: distinct positions:", len(distinctPositions))
}

func (d *day06) part2Channels() {
	sum := 0

	var wg sync.WaitGroup
	wg.Add(len(d.distinctPositions))
	ch := make(chan bool)

	for _, c := range d.distinctPositions {
		go func(c coord) {
			ch <- d.doesInduceLoop(c)
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		if r {
			sum++
		}
	}

	fmt.Println("ANSWER2: number of positions that induce a loop:", sum)
}

func parse(filename string) *day06 {
	data := utils.SplitLines(filename)

	grid := grid{}
	var start coord

	for y, line := range data {
		row := make([]visitedDirection, len(line))
		for x, r := range line {
			direction := visitedDirection{}
			switch r {
			case '#':
				direction[box] = true
			case '^':
				direction[up] = true
				start = coord{x, y}
			}
			row[x] = direction
		}
		grid = append(grid, row)
	}

	return &day06{grid, start, nil}
}

func Solve(filename string) {
	d := parse(filename)
	d.part1()
	d.part2Channels()
}
