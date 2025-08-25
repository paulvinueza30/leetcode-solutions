package main

import "container/heap"

type stoneHeap []int

func (h stoneHeap) Len() int{
	return len(h)
}

func (h stoneHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h stoneHeap) Swap(i int , j int) {
	h[i] , h[j] = h[j] , h[i]
}

func (h *stoneHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *stoneHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	stoneHeap:= &stoneHeap{}
	heap.Init(stoneHeap)
	for _, s := range stones {
		heap.Push(stoneHeap, s)
	}
	
	for len(*stoneHeap) > 1 {
		stoneA := heap.Pop(stoneHeap).(int)
		stoneB := heap.Pop(stoneHeap).(int)
		if stoneA != stoneB{
			heap.Push(stoneHeap, stoneA - stoneB)
		}
	}
	if len(*stoneHeap) == 1 {
		return heap.Pop(stoneHeap).(int)
	}
	return 0
}

