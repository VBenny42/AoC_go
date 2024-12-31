package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/davecgh/go-spew/spew"
)

type (
	grid  [][]rune
	coord struct {
		x, y int
	}
)

type day12 struct {
	regions      map[rune][]map[coord]struct{}
	notNeighbors [][]map[direction]struct{}
}

type direction [2]int

var (
	UP    direction = [2]int{0, -1}
	DOWN  direction = [2]int{0, 1}
	LEFT  direction = [2]int{-1, 0}
	RIGHT direction = [2]int{1, 0}
)

func (c coord) getNeighbor(g grid, dir direction) (coord, error) {
	value := g[c.y][c.x]
	var neighbor coord
	switch dir {
	case UP:
		if c.y == 0 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x, c.y - 1}
	case DOWN:
		if c.y == len(g)-1 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x, c.y + 1}
	case LEFT:
		if c.x == 0 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x - 1, c.y}
	case RIGHT:
		if c.x == len(g[0])-1 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x + 1, c.y}
	}
	if value != g[neighbor.y][neighbor.x] {
		return coord{}, fmt.Errorf("No neighbor in that direction")
	}
	return neighbor, nil
}

func buildRegions(g grid) (map[rune][]map[coord]struct{}, [][]map[direction]struct{}) {
	m, n := len(g[0]), len(g)
	regions := make(map[rune][]map[coord]struct{})
	visited := make([][]bool, n)
	notNeighbors := make([][]map[direction]struct{}, n)

	for i := range visited {
		visited[i] = make([]bool, m)
		notNeighbors[i] = make([]map[direction]struct{}, m)
		for j := range notNeighbors[i] {
			notNeighbors[i][j] = make(map[direction]struct{})
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] {
				continue
			}

			region := make(map[coord]struct{})
			stack := []coord{{j, i}}

			for len(stack) > 0 {
				position := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				region[position] = struct{}{}
				visited[position.y][position.x] = true

				for _, direction := range []direction{UP, RIGHT, DOWN, LEFT} {
					nextPosition, err := position.getNeighbor(g, direction)
					if err == nil && !visited[nextPosition.y][nextPosition.x] {
						stack = append(stack, nextPosition)
					} else if err != nil {
						notNeighbors[position.y][position.x][direction] = struct{}{}
					}
				}
			}

			if regions[g[i][j]] == nil {
				regions[g[i][j]] = []map[coord]struct{}{region}
			} else {
				regions[g[i][j]] = append(regions[g[i][j]], region)
			}
		}
	}

	return regions, notNeighbors
}

func (d *day12) calculatePerimeter(region map[coord]struct{}) int {
	perimeter := 0
	for c := range region {
		perimeter += len(d.notNeighbors[c.y][c.x])
	}
	return perimeter
}

func countSides(region map[coord]struct{}) int {
	sideCount := 0

	for _, dir := range []direction{UP, RIGHT, DOWN, LEFT} {
		dX, dY := dir[0], dir[1]
		visited := make(map[coord]struct{})

		for c := range region {
			if _, ok := visited[c]; ok {
				continue
			}

			neighbor := coord{c.x + dX, c.y + dY}
			if _, ok := region[neighbor]; ok {
				continue
			}
			sideCount++

			for _, d := range []int{-1, 1} {
				fX, fY := c.x, c.y
				for {
					_, fOk := region[coord{fX, fY}]
					_, nOk := region[coord{fX + dX, fY + dY}]
					if !fOk || nOk {
						break
					}
					visited[coord{fX, fY}] = struct{}{}
					fX += d * dY
					fY += d * dX
				}
			}
		}
	}

	return sideCount
}

func (d *day12) part1() {
	price := 0
	for _, region := range d.regions {
		for _, r := range region {
			price += d.calculatePerimeter(r) * len(r)
		}
	}

	fmt.Println("ANSWER1: price:", price)
}

func (d *day12) part2() {
	price := 0
	for _, region := range d.regions {
		for _, r := range region {
			price += countSides(r) * len(r)
		}
	}
	fmt.Println("ANSWER2: price:", price)
}

func parse() *day12 {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	grid := [][]rune{}

	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}

	regions, notNeighbors := buildRegions(grid)

	return &day12{regions, notNeighbors}
}

func main() {
	d := parse()
	d.part1()
	d.part2()
}