package lastStoneWeight

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	heapStones := IntHeap(stones)
	heap.Init(&heapStones)
	for {
		biggerStone := heap.Pop(&heapStones).(int)
		smallerStone := heap.Pop(&heapStones).(int)
		newStone := biggerStone - smallerStone
		if newStone != 0 {
			heap.Push(&heapStones, newStone)
		}

		stoneLen := len(heapStones)
		if stoneLen == 1 {
			return heap.Pop(&heapStones).(int)
		} else if stoneLen == 0 {
			return 0
		}
	}
	return 0
}
