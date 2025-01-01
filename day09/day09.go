package main

import (
	"container/heap"
	"fmt"
	"os"
)

// TIL runes are just int32, id's were stored as ints, so freeSpace was colliding with id's
// Max is is 10000, so I just used 1000001 as a sentinel value
const freeSpace = 1000001

type day09 struct {
	diskmap []int
}

func (d *day09) convert() []int {
	isFreeSpace := false
	id := 0

	diskmap := make([]int, 0)

	for _, i := range d.diskmap {
		if isFreeSpace {
			for j := 0; j < i; j++ {
				diskmap = append(diskmap, freeSpace)
			}
		} else {
			for j := 0; j < i; j++ {
				diskmap = append(diskmap, id)
			}
			id++
		}
		isFreeSpace = !isFreeSpace
	}
	return diskmap
}

func (d *day09) convertWithHeaps() ([]int, []intHeap) {
	isFreeSpace := false
	id := 0

	diskmap := make([]int, 0)
	heaps := make([]intHeap, 10)
	for i := range heaps {
		heap.Init(&heaps[i])
	}

	for _, i := range d.diskmap {
		if isFreeSpace {
			heap.Push(&heaps[i], len(diskmap))
			for j := 0; j < i; j++ {
				diskmap = append(diskmap, freeSpace)
			}
		} else {
			for j := 0; j < i; j++ {
				diskmap = append(diskmap, id)
			}
			id++
		}
		isFreeSpace = !isFreeSpace
	}
	return diskmap, heaps
}

func makeContiguous(diskmap []int) []int {
	firstFreeBlock := 0
	lastFileBlock := len(diskmap) - 1
	for firstFreeBlock < lastFileBlock {
		for firstFreeBlock < len(diskmap) && diskmap[firstFreeBlock] != freeSpace {
			firstFreeBlock++
		}
		for lastFileBlock >= 0 && diskmap[lastFileBlock] == freeSpace {
			lastFileBlock--
		}

		if firstFreeBlock < lastFileBlock {
			diskmap[firstFreeBlock], diskmap[lastFileBlock] = diskmap[lastFileBlock], diskmap[firstFreeBlock]
		}
	}
	return diskmap
}

func makeContiguousHeaps(diskmap []int, heaps []intHeap) []int {
	index := len(diskmap) - 1

	for index >= 0 {
		if diskmap[index] == freeSpace {
			index--
			continue
		}

		id := diskmap[index]
		fileWidth := 0
		for index >= 0 && diskmap[index] == id {
			fileWidth++
			index--
		}

		bestWidth := -1
		smallestIndex := len(diskmap)

		for width := fileWidth; width < len(heaps); width++ {
			if heaps[width].Len() > 0 {
				if smallestIndex > heaps[width][0] {
					smallestIndex = heaps[width][0]
					bestWidth = width
				}
			}
		}

		if smallestIndex == len(diskmap) {
			continue
		}
		if smallestIndex > index {
			continue
		}

		heap.Pop(&heaps[bestWidth])
		for j := 0; j < fileWidth; j++ {
			diskmap[smallestIndex+j] = id
			diskmap[index+j+1] = freeSpace
		}
		heap.Push(&heaps[bestWidth-fileWidth], smallestIndex+fileWidth)
	}

	return diskmap
}

func checksum(diskmap []int) int {
	checksum := 0
	for i, id := range diskmap {
		if id == freeSpace {
			continue
		}
		checksum += i * id
	}
	return checksum
}

func (d *day09) part1() {
	diskmap := d.convert()
	diskmap = makeContiguous(diskmap)

	fmt.Println("ANSWER1: checksum:", checksum(diskmap))
}

func (d *day09) part2() {
	diskmap, heaps := d.convertWithHeaps()
	diskmap = makeContiguousHeaps(diskmap, heaps)

	fmt.Println("ANSWER2: checksum:", checksum(diskmap))
}

func parse() *day09 {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	line := string(file)[:len(file)-1]
	diskmap := make([]int, len(line))

	for i, c := range line {
		value := int(c - '0')
		diskmap[i] = value
	}

	return &day09{diskmap}
}

func main() {
	d := parse()
	d.part1()
	d.part2()
}
