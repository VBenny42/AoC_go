package day12

import (
	"fmt"

	"github.com/VBenny42/AoC/2024/golang/utils"
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
	up    direction = [2]int{0, -1}
	down  direction = [2]int{0, 1}
	left  direction = [2]int{-1, 0}
	right direction = [2]int{1, 0}
)

func (c coord) getNeighbor(g grid, dir direction) (coord, error) {
	value := g[c.y][c.x]
	var neighbor coord
	switch dir {
	case up:
		if c.y == 0 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x, c.y - 1}
	case down:
		if c.y == len(g)-1 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x, c.y + 1}
	case left:
		if c.x == 0 {
			return coord{}, fmt.Errorf("No neighbor in that direction")
		}
		neighbor = coord{c.x - 1, c.y}
	case right:
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

				for _, direction := range []direction{up, right, down, left} {
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

	for _, dir := range []direction{up, right, down, left} {
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

func (d *day12) Part1() int {
	price := 0
	for _, region := range d.regions {
		for _, r := range region {
			price += d.calculatePerimeter(r) * len(r)
		}
	}

	return price
}

func (d *day12) Part2() int {
	price := 0
	for _, region := range d.regions {
		for _, r := range region {
			price += countSides(r) * len(r)
		}
	}
	return price
}

func Parse(filename string) *day12 {
	data := utils.SplitLines(filename)
	grid := [][]rune{}

	for _, row := range data {
		grid = append(grid, []rune(row))
	}

	regions, notNeighbors := buildRegions(grid)

	return &day12{regions, notNeighbors}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: price:", d.Part1())
	fmt.Println("ANSWER2: price:", d.Part2())
}
