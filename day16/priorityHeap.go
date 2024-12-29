package main

import (
	"container/heap"
)

type item struct {
	cost int
	state
	index int
}

type priorityHeap []*item

func (h priorityHeap) Len() int           { return len(h) }
func (h priorityHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h priorityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *priorityHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*item)
	item.index = n
	*h = append(*h, item)
}

func (h *priorityHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *priorityHeap) update(item *item, cost int) {
	item.cost = cost
	heap.Fix(h, item.index)
}
