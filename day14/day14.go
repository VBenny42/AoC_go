package day14

import (
	"fmt"
	"os"
	"strconv"

	"github.com/VBenny42/AoC/2024/golang/utils"
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
	topLeft = iota
	topRight
	bottomLeft
	bottomRight
)

type day14 struct {
	robots []robot
	grid   grid
}

func (d *day14) moveRobot(r *robot) {
	newPosX := ((r.position.x+r.velocity.x)%cols + cols) % cols
	newPosY := ((r.position.y+r.velocity.y)%rows + rows) % rows
	d.grid[r.position.y][r.position.x]--
	d.grid[newPosY][newPosX]++
	r.position.x = newPosX
	r.position.y = newPosY
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

func (g *grid) getSafetyFactor() int {
	getBoundingBox := func(quadrant int) (int, int, int, int) {
		var startRow, startCol, endRow, endCol int
		switch quadrant {
		case topLeft:
			startRow = 0
			startCol = 0
			endRow = rows / 2
			endCol = cols / 2
		case topRight:
			startRow = 0
			startCol = cols/2 + 1
			endRow = rows / 2
			endCol = cols
		case bottomLeft:
			startRow = rows/2 + 1
			startCol = 0
			endRow = rows
			endCol = cols / 2
		case bottomRight:
			startRow = rows/2 + 1
			startCol = cols/2 + 1
			endRow = rows
			endCol = cols
		}
		return startRow, startCol, endRow, endCol
	}

	safetyFactor := 1

	for q := topLeft; q <= bottomRight; q++ {
		startRow, startCol, endRow, endCol := getBoundingBox(q)
		numRobots := 0
		for i := startRow; i < endRow; i++ {
			for j := startCol; j < endCol; j++ {
				numRobots += g[i][j]
			}
		}
		safetyFactor *= numRobots
	}

	return safetyFactor
}

func (d *day14) Part1() int {
	for j := 0; j < 100; j++ {
		for i := range d.robots {
			d.moveRobot(&d.robots[i])
		}
	}

	return d.grid.getSafetyFactor()
}

func (d *day14) Part2() int {
	minSafetyFactor, minIteration := 1<<31-1, 0
	for i := 100; i < 10000; i++ {
		for j := range d.robots {
			d.moveRobot(&d.robots[j])
		}

		// if i == 7752 {
		// 	d.grid.printScaledBitmap(fmt.Sprintf("output%d.pbm", i), 7)
		// }

		safetyFactor := d.grid.getSafetyFactor()
		if safetyFactor < minSafetyFactor {
			minSafetyFactor = safetyFactor
			minIteration = i
		}
	}

	return minIteration
}

func Parse(filename string) *day14 {
	data := utils.SplitLines(filename)

	grid := grid{}
	robots := []robot{}
	var posX, posY, velX, velY int

	for _, line := range data {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &velX, &velY)
		robots = append(robots, robot{
			position: struct{ x, y int }{posX, posY},
			velocity: struct{ x, y int }{velX, velY},
		})
		grid[posY][posX]++
	}

	return &day14{robots, grid}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: safetyFactor:", d.Part1())
	fmt.Println("ANSWER2: minSafetyFactor is at iteration:", d.Part2()+1)
}
