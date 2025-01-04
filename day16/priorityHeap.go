package day16

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
func (h priorityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *priorityHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*item)
	item.index = n
	*h = append(*h, item)
}

func (h *priorityHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

func (h *priorityHeap) update(item *item, cost int) {
	item.cost = cost
	heap.Fix(h, item.index)
}
