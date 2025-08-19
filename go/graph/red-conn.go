package main

// Build the graph
// Loop backwords thru edges
// Take out the curr edge
// check for cycle
// if cycle return edge

func hasCycle(curr int, par int, g map[int][]int, v map[int]bool) bool {
	v[curr] = true
	for _, nei := range g[curr] {
		if !v[nei] {
			if hasCycle(nei, curr, g, v) {
				return true
			}
		} else if nei != par {
			return true
		}
	}
	return false
}

func findReduantConnection(edges [][]int) []int {
	g := make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	n := len(g)
	var res = []int{}
	for i := len(edges); i >= 0; i-- {
		currEdge := edges[i]
		v1, v2 := currEdge[0], currEdge[1]
		v1List, v2List := g[v1], g[v2]
		var fv1List, fv2List []int
		for _, v := range v1List {
			if v != v2 {
				fv1List = append(fv1List, v)
			}
		}
		for _, v := range v2List {
			if v != v1 {
				fv2List = append(fv2List, v)
			}
		}
		g[v1] = fv1List
		g[v2] = fv2List
		v := make(map[int]bool)
		if !hasCycle(1, -1, g, v) && len(v) == n {
			return currEdge
		}
		g[v1] = v1List
		g[v2] = v2List
	}
	return res
}
