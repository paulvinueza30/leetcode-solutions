package main

import (
	"container/heap"
)

type Point struct {
	dist int;
	point []int
}
type pointHeap []Point

func (h pointHeap) Len() int{
	return len(h)
}

func (h pointHeap) Less(i int, j int) bool {
	return h[i].dist > h[j].dist
}

func (h pointHeap) Swap(i int , j int) {
	h[i] , h[j] = h[j] , h[i]
}

func (h *pointHeap) Push(x any) {
	*h = append(*h, x.(Point))
}
func (h *pointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int{
	pointHeap := &pointHeap{}
	heap.Init(pointHeap)
	for _, p := range points{
		dist := (p[0] * p[0]) + (p[1] * p[1])
		heap.Push(pointHeap, Point{dist: dist, point: p})
	    if pointHeap.Len() > k{
            heap.Pop(pointHeap)
        }
	}
	
		
	res := make([][]int, 0 , k)	
	for pointHeap.Len() >= 1 {
		res = append(res, pointHeap.Pop().(Point).point)
	}
	return res
}