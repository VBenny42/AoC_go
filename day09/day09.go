package day09

import (
	"container/heap"
	"fmt"

	"github.com/VBenny42/AoC/2024/golang/utils"
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

func (d *day09) Part1() int {
	diskmap := d.convert()
	diskmap = makeContiguous(diskmap)

	return checksum(diskmap)
}

func (d *day09) Part2() int {
	diskmap, heaps := d.convertWithHeaps()
	diskmap = makeContiguousHeaps(diskmap, heaps)

	return checksum(diskmap)
}

func Parse(filename string) *day09 {
	data := utils.JoinFile(filename)

	diskmap := make([]int, len(data))

	for i, c := range data {
		value := int(c - '0')
		diskmap[i] = value
	}

	return &day09{diskmap}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: checksum:", d.Part1())
	fmt.Println("ANSWER2: checksum:", d.Part2())
}
