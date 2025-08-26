package main

import "container/heap"
type Task struct {
	task byte;
	count int;
}
type taskHeap []Task

func (h taskHeap) Len() int{
	return len(h)
}

func (h taskHeap) Less(i int, j int) bool {
	return h[i].count > h[j].count
}

func (h taskHeap) Swap(i int , j int) {
	h[i] , h[j] = h[j] , h[i]
}

func (h *taskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}
func (h *taskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Have a min heap for tasks with task and count
// Have a dict for keeping track of count as you put in heap
// Have a dict for keeping track of lastCycle for specific task as you run sim
// Run the sim use above dict to decide whether to skip keep a running count
// return count

func leastInterval(tasks []byte, n int) int{
	h := &taskHeap{}	
	heap.Init(h)
	counts := make(map[byte]int)
	for _, t := range tasks{
		_ , ok := counts[t]
		if !ok {
			counts[t] = 0
		}
		counts[t]++
	}
	for t , c := range counts{
		heap.Push(h, Task{task: t, count: c})
	}
	currTime := 0
	for h.Len() > 0 {
		temp := []Task{}
		for range n + 1{
			if h.Len() == 0{
				break
			}
			temp = append(temp, heap.Pop(h).(Task))
			currTime++
		}
		for _, t := range temp{
			if t.count - 1 > 0 {
				heap.Push(h, Task{task: t.task, count: t.count - 1})
			}
		}
		if h.Len() > 0 {
			currTime += (n + 1 - len(temp))
		}
	}
	return currTime
}