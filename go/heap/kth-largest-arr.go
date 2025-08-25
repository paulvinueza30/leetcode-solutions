package main

import "container/heap"

type numHeap []int

func (h numHeap) Len() int{
	return len(h)
}

func (h numHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h numHeap) Swap(i int , j int) {
	h[i] , h[j] = h[j] , h[i]
}

func (h *numHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *numHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &numHeap{}
	heap.Init(h)
	for _ , num := range nums {
		if h.Len() < k{
		heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
}
	return (*h)[0]
}