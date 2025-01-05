package day04

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type (
	grid  [][]rune
	coord struct {
		x, y int
	}
)

type day04 struct {
	grid    grid
	aCoords []coord
	xCoords []coord
}

type direction int

const (
	up direction = iota
	down
	left
	right
	upLeft
	upRight
	downLeft
	downRight
)

func (d *day04) getAdjacentLetters(c coord) map[direction]coord {
	adjacents := map[direction]coord{}
	if c.x > 0 {
		adjacents[left] = coord{c.x - 1, c.y}
	}
	if c.x < len(d.grid[0])-1 {
		adjacents[right] = coord{c.x + 1, c.y}
	}
	if c.y > 0 {
		adjacents[up] = coord{c.x, c.y - 1}
	}
	if c.y < len(d.grid)-1 {
		adjacents[down] = coord{c.x, c.y + 1}
	}
	if c.x > 0 && c.y > 0 {
		adjacents[upLeft] = coord{c.x - 1, c.y - 1}
	}
	if c.x < len(d.grid[0])-1 && c.y > 0 {
		adjacents[upRight] = coord{c.x + 1, c.y - 1}
	}
	if c.x > 0 && c.y < len(d.grid)-1 {
		adjacents[downLeft] = coord{c.x - 1, c.y + 1}
	}
	if c.x < len(d.grid[0])-1 && c.y < len(d.grid)-1 {
		adjacents[downRight] = coord{c.x + 1, c.y + 1}
	}

	return adjacents
}

func (d *day04) isXmasMatch(c coord, dir direction, currentMatch string) bool {
	xmas := "XMAS"

	adjacents := d.getAdjacentLetters(c)
	adjacentCoord, ok := adjacents[dir]
	if !ok {
		return false
	}

	letter := d.grid[c.y][c.x]
	adjacentLetter := d.grid[adjacentCoord.y][adjacentCoord.x]

	potentialMatch := currentMatch + string(letter) + string(adjacentLetter)

	if strings.HasPrefix(xmas, potentialMatch) {
		if len(potentialMatch) == len(xmas) {
			return true
		}
		return d.isXmasMatch(adjacentCoord, dir, currentMatch+string(letter))
	}

	return false
}

func (d *day04) isXMasMatch(c coord) bool {
	adjacents := d.getAdjacentLetters(c)

	for _, corner := range []direction{upLeft, upRight, downLeft, downRight} {
		if _, ok := adjacents[corner]; !ok {
			return false
		}
	}

	upLeftValue := d.grid[adjacents[upLeft].y][adjacents[upLeft].x]
	downRightValue := d.grid[adjacents[downRight].y][adjacents[downRight].x]

	upRightValue := d.grid[adjacents[upRight].y][adjacents[upRight].x]
	downLeftValue := d.grid[adjacents[downLeft].y][adjacents[downLeft].x]

	return ((upLeftValue == 'M' && downRightValue == 'S') ||
		(upLeftValue == 'S' && downRightValue == 'M')) &&
		((upRightValue == 'M' && downLeftValue == 'S') ||
			(upRightValue == 'S' && downLeftValue == 'M'))
}

func (d *day04) Part1() int {
	sum := 0
	for _, xCoord := range d.xCoords {
		for dir := up; dir <= downRight; dir++ {
			if d.isXmasMatch(xCoord, dir, "") {
				sum++
			}
		}
	}
	return sum
}

func (d *day04) Part2() int {
	sum := 0
	for _, aCoord := range d.aCoords {
		if d.isXMasMatch(aCoord) {
			sum++
		}
	}
	return sum
}

func Parse(filename string) *day04 {
	data := utils.SplitLines(filename)

	grid := grid{}

	aCoords := []coord{}
	xCoords := []coord{}

	for y, line := range data {
		row := []rune(line)
		for x, r := range line {
			if r == 'A' {
				aCoords = append(aCoords, coord{x, y})
			}
			if r == 'X' {
				xCoords = append(xCoords, coord{x, y})
			}
		}
		grid = append(grid, row)
	}

	return &day04{grid, aCoords, xCoords}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: xmas matches:", d.Part1())
	fmt.Println("ANSWER2: x-mas matches:", d.Part2())
}
