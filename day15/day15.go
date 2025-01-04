package day15

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type (
	grid  [][]rune
	coord struct {
		x, y int
	}
)

const (
	wall  = '#'
	open  = '.'
	robot = '@'
	box   = 'O'
	lBox  = '['
	rBox  = ']'
)

type direction [2]int

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

type day15 struct {
	g         grid
	movements []direction
}

func (d *day15) findRobot() (coord, error) {
	for y, row := range d.g {
		for x, c := range row {
			if c == robot {
				return coord{x, y}, nil
			}
		}
	}
	return coord{-1, -1}, fmt.Errorf("robot not found")
}

func (d *day15) boxesToMove(c coord, dir direction, boxesSoFar []coord) ([]coord, error) {
	newX := c.x + dir[0]
	newY := c.y + dir[1]
	boxesSoFar = append(boxesSoFar, c)

	if d.g[newY][newX] == wall {
		return nil, fmt.Errorf("Hit wall")
	}
	if d.g[newY][newX] == open {
		return boxesSoFar, nil
	}
	if d.g[newY][newX] == box {
		return d.boxesToMove(coord{newX, newY}, dir, boxesSoFar)
	}
	return nil, fmt.Errorf("Shouldn't get here")
}

func (d *day15) boxesToMove2(boxSides [2]coord, dir direction, boxesSoFar [][2]coord) ([][2]coord, error) {
	boxesSoFar = append(boxesSoFar, boxSides)
	if dir == left || dir == right {
		var x, y int
		var bracket rune
		if dir == left {
			x, y = boxSides[0].x, boxSides[0].y
			bracket = rBox
		} else {
			x, y = boxSides[1].x, boxSides[1].y
			bracket = lBox
		}
		newX, newY := x+dir[0], y+dir[1]
		if d.g[newY][newX] == wall {
			return nil, fmt.Errorf("Hit wall")
		}
		if d.g[newY][newX] == open {
			return boxesSoFar, nil
		}
		if d.g[newY][newX] == bracket {
			var adjacentBox [2]coord
			if dir == left {
				adjacentBox = [2]coord{{newX - 1, newY}, {newX, newY}}
			} else {
				adjacentBox = [2]coord{{newX, newY}, {newX + 1, newY}}
			}
			return d.boxesToMove2(adjacentBox, dir, boxesSoFar)
		}
	} else {
		newLeftX, newLeftY := boxSides[0].x+dir[0], boxSides[0].y+dir[1]
		newRightX, newRightY := boxSides[1].x+dir[0], boxSides[1].y+dir[1]

		if d.g[newLeftY][newLeftX] == wall || d.g[newRightY][newRightX] == wall {
			return nil, fmt.Errorf("Hit wall")
		}

		if d.g[newLeftY][newLeftX] == open && d.g[newRightY][newRightX] == open {
			return boxesSoFar, nil
		}

		if d.g[newLeftY][newLeftX] == lBox && d.g[newRightY][newRightX] == rBox {
			return d.boxesToMove2([2]coord{{newLeftX, newLeftY}, {newRightX, newRightY}}, dir, boxesSoFar)
		}

		if d.g[newLeftY][newLeftX] == rBox {
			leftBoxes, err := d.boxesToMove2([2]coord{{newLeftX - 1, newLeftY}, {newLeftX, newLeftY}}, dir, [][2]coord{})
			if err != nil {
				return nil, err
			}
			boxesSoFar = append(boxesSoFar, leftBoxes...)
		}
		if d.g[newRightY][newRightX] == lBox {
			rightBoxes, err := d.boxesToMove2([2]coord{{newRightX, newRightY}, {newRightX + 1, newRightY}}, dir, [][2]coord{})
			if err != nil {
				return nil, err
			}
			boxesSoFar = append(boxesSoFar, rightBoxes...)
		}
		return boxesSoFar, nil
	}

	return nil, fmt.Errorf("Shouldn't get here")
}

func (d *day15) moveRobot(c coord, dir direction) (coord, error) {
	newX := c.x + dir[0]
	newY := c.y + dir[1]

	if d.g[newY][newX] == wall {
		return c, nil
	}

	if d.g[newY][newX] == open {
		d.g[c.y][c.x] = open
		d.g[newY][newX] = robot
		return coord{newX, newY}, nil
	}

	if d.g[newY][newX] == box {
		boxes, err := d.boxesToMove(coord{newX, newY}, dir, []coord{})
		if err != nil {
			return c, nil
		}

		switch dir {
		case up:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i].y < boxes[j].y
			})
		case down:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i].y > boxes[j].y
			})
		case left:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i].x < boxes[j].x
			})
		case right:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i].x > boxes[j].x
			})
		}

		for _, b := range boxes {
			newX := b.x + dir[0]
			newY := b.y + dir[1]
			d.g[b.y][b.x] = open
			d.g[newY][newX] = box
		}
		d.g[c.y][c.x] = open
		d.g[newY][newX] = robot
		return coord{newX, newY}, nil
	}
	return c, fmt.Errorf("Shouldn't get here")
}

func (d *day15) moveRobot2(c coord, dir direction) (coord, error) {
	newX := c.x + dir[0]
	newY := c.y + dir[1]

	if d.g[newY][newX] == wall {
		return c, nil
	}

	if d.g[newY][newX] == open {
		d.g[c.y][c.x] = open
		d.g[newY][newX] = robot
		return coord{newX, newY}, nil
	}

	if d.g[newY][newX] == lBox || d.g[newY][newX] == rBox {
		var boxPair [2]coord
		if d.g[newY][newX] == lBox {
			boxPair = [2]coord{{newX, newY}, {newX + 1, newY}}
		} else {
			boxPair = [2]coord{{newX - 1, newY}, {newX, newY}}
		}
		boxes, err := d.boxesToMove2(boxPair, dir, [][2]coord{})
		if err != nil {
			return c, nil
		}
		switch dir {
		case up:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i][0].y < boxes[j][0].y
			})
		case down:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i][0].y > boxes[j][0].y
			})
		case left:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i][0].x < boxes[j][0].x
			})
		case right:
			sort.Slice(boxes, func(i, j int) bool {
				return boxes[i][0].x > boxes[j][0].x
			})
		}
		for _, b := range boxes {
			boxL, boxR := b[0], b[1]
			newBoxL, newBoxR := coord{boxL.x + dir[0], boxL.y + dir[1]}, coord{boxR.x + dir[0], boxR.y + dir[1]}
			d.g[boxL.y][boxL.x] = open
			d.g[boxR.y][boxR.x] = open
			d.g[newBoxL.y][newBoxL.x] = lBox
			d.g[newBoxR.y][newBoxR.x] = rBox
		}
		d.g[c.y][c.x] = open
		d.g[newY][newX] = robot
		return coord{newX, newY}, nil
	}
	return c, fmt.Errorf("Shouldn't get here")
}

func (d *day15) findBoxes() []coord {
	boxes := make([]coord, 0)
	for y, row := range d.g {
		for x, c := range row {
			if c == box || c == lBox {
				boxes = append(boxes, coord{x, y})
			}
		}
	}
	return boxes
}

func (d *day15) scaleGrid() {
	newGrid := make(grid, 0)
	for _, row := range d.g {
		newRow := make([]rune, 0)
		for _, c := range row {
			switch c {
			case box:
				newRow = append(newRow, lBox, rBox)
			case robot:
				newRow = append(newRow, robot, open)
			case wall:
				newRow = append(newRow, wall, wall)
			case open:
				newRow = append(newRow, open, open)
			}
		}
		newGrid = append(newGrid, newRow)
	}
	d.g = newGrid
}

func (c coord) gpsCoordinate() int {
	return c.y*100 + c.x
}

func (d *day15) Part1() int {
	robot, err := d.findRobot()
	if err != nil {
		fmt.Println("Error finding robot", err)
		return -1
	}

	for _, move := range d.movements {
		robot, _ = d.moveRobot(robot, move)
	}

	sumGpsCoordinates := 0
	for _, box := range d.findBoxes() {
		sumGpsCoordinates += box.gpsCoordinate()
	}

	return sumGpsCoordinates
}

func (d *day15) Part2() int {
	d.scaleGrid()

	robot, err := d.findRobot()
	if err != nil {
		fmt.Println("Error finding robot", err)
		return -1
	}

	for _, move := range d.movements {
		robot, _ = d.moveRobot2(robot, move)
	}

	sumGpsCoordinates := 0
	for _, box := range d.findBoxes() {
		sumGpsCoordinates += box.gpsCoordinate()
	}

	return sumGpsCoordinates
}

func Parse(filename string) *day15 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return nil
	}
	defer file.Close()

	g := make(grid, 0)

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		g = append(g, []rune(line))
	}

	movements := make([]direction, 0)
	for s.Scan() {
		line := s.Text()
		for _, c := range line {
			switch c {
			case '^':
				movements = append(movements, up)
			case 'v':
				movements = append(movements, down)
			case '<':
				movements = append(movements, left)
			case '>':
				movements = append(movements, right)
			}
		}
	}

	return &day15{g, movements}
}

func (d *day15) printGrid() {
	for _, row := range d.g {
		fmt.Println(string(row))
	}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: sumGpsCoordinates:", Parse(filename).Part1())
	fmt.Println("ANSWER2: sumGpsCoordinates:", Parse(filename).Part2())
}
