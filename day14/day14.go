package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type robot struct {
	position struct {
		x, y int
	}
	velocity struct {
		x, y int
	}
}

const (
	rows = 103
	cols = 101
)

type grid [rows][cols]int

const (
	TOP_LEFT = iota
	TOP_RIGHT
	BOTTOM_LEFT
	BOTTOM_RIGHT
)

type day14 struct {
	robots []robot
	grid   grid
}

func (d *day14) moveRobot(r *robot) {
	newPosX := ((r.position.x+r.velocity.x)%cols + cols) % cols
	newPosY := ((r.position.y+r.velocity.y)%rows + rows) % rows
	d.grid[r.position.y][r.position.x] -= 1
	d.grid[newPosY][newPosX] += 1
	r.position.x = newPosX
	r.position.y = newPosY
}

func (g *grid) getQuadrant(quadrant int) [][]int {
	var startRow, startCol, endRow, endCol int
	switch quadrant {
	case TOP_LEFT:
		startRow = 0
		startCol = 0
		endRow = rows / 2
		endCol = cols / 2
	case TOP_RIGHT:
		startRow = 0
		startCol = cols/2 + 1
		endRow = rows / 2
		endCol = cols
	case BOTTOM_LEFT:
		startRow = rows/2 + 1
		startCol = 0
		endRow = rows
		endCol = cols / 2
	case BOTTOM_RIGHT:
		startRow = rows/2 + 1
		startCol = cols/2 + 1
		endRow = rows
		endCol = cols
	}
	quadrantGrid := make([][]int, endRow-startRow)
	for i := 0; i < endRow-startRow; i++ {
		quadrantGrid[i] = make([]int, endCol-startCol)
		for j := 0; j < endCol-startCol; j++ {
			quadrantGrid[i][j] = g[i+startRow][j+startCol]
		}
	}
	return quadrantGrid
}

func (g *grid) printScaledBitmap(filename string, scaleFactor int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	line := fmt.Sprintf("P1\n%d %d\n", cols*scaleFactor, rows*scaleFactor)
	file.WriteString(line)

	for i := 0; i < rows; i++ {
		for k := 0; k < scaleFactor; k++ { // Repeat each row scaleFactor times
			for j := 0; j < cols; j++ {
				for l := 0; l < scaleFactor; l++ { // Repeat each column scaleFactor times
					file.WriteString(strconv.Itoa(g[i][j]))
				}
			}
			file.WriteString("\n") // New line after each scaled row
		}
	}
}

func printGrid(g *grid) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g[i][j] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func (d *day14) part1() {
	for _, r := range d.robots {
		for i := 0; i < 100; i++ {
			d.moveRobot(&r)
		}
	}

	safetyFactor := 1
	for q := TOP_LEFT; q <= BOTTOM_RIGHT; q++ {
		quadrantGrid := d.grid.getQuadrant(q)
		numRobots := 0
		for i := 0; i < len(quadrantGrid); i++ {
			for j := 0; j < len(quadrantGrid[i]); j++ {
				numRobots += quadrantGrid[i][j]
			}
		}
		safetyFactor *= numRobots
	}

	fmt.Println("ANSWER1: safetyFactor:", safetyFactor)
}

func (d *day14) part2() {
	for i := 0; i < 10000; i++ {
		for _, r := range d.robots {
			d.moveRobot(&r)
		}

		// TODO: Grid looked the same for all iterations, no idea why
		// Couldn't get it working like python version,
		// correct iteration should be 7752
		if i >= 7750 && i <= 7755 {
			d.grid.printScaledBitmap(fmt.Sprintf("output%d.pbm", i), 7)
		}

	}
}

func parse() *day14 {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	grid := grid{}

	s := bufio.NewScanner(file)
	var posX, posY, velX, velY int
	robots := []robot{}
	for s.Scan() {
		fmt.Sscanf(s.Text(), "p=%d,%d v=%d,%d", &posX, &posY, &velX, &velY)
		robots = append(robots, robot{
			position: struct{ x, y int }{posX, posY},
			velocity: struct{ x, y int }{velX, velY},
		})
		grid[posY][posX] += 1
	}

	return &day14{robots, grid}
}

func main() {
	parse().part1()
	// parse().part2()
}
