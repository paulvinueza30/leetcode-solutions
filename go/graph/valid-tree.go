package main

import (
	"container/list"
)

func validTree(n int, edges [][]int) bool {
	g := make(map[int][]int)
	v := make(map[int]bool)
	q := list.New()

	if len(edges) != n-1 {
		return false
	}

	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	q.PushBack([]int{0, -1})

	for q.Len() > 0 {
		front := q.Front()
		temp := front.Value.([]int)
		curr, par := temp[0], temp[1]
		q.Remove(front)
		v[curr] = true
		for _, nei := range g[curr] {
			if !v[curr] {
				q.PushBack([]int{nei, curr})
				continue
			}
			if par != nei {
				return false
			}
		}
	}
	return n == 0
}
