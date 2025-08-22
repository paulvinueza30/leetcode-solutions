package main

import (
	"container/heap"
)

type intHeap []int

func (h intHeap) Len() int{
	return len(h)
}

func (h intHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i int , j int) {
	h[i] , h[j] = h[j] , h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	h *intHeap
	k int
}

func Constructor(k int, nums [] int) KthLargest {
	h := &intHeap{}
	heap.Init(h)
	for _ , num := range nums{
		heap.Push(h, num)
		if h.Len() > k{
			heap.Pop(h)
		}
	}
	return KthLargest{h: h, k: k}
}

func (t *KthLargest) Add(val int) int{
	heap.Push(t.h, val)
	if t.h.Len() > t.k{
		heap.Pop(t.h)
	}
	return (*t.h)[0]
}
