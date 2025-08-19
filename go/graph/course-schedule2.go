package main

func dfs2(curr int, v map[int]int, g map[int][]int, r *[]int) bool {
	if v[curr] == 1 {
		return false
	}
	if v[curr] == -1 {
		return true
	}
	v[curr] = 1

	for _, nei := range g[curr] {
		if !dfs2(nei, v, g, r) {
			return false
		}
	}
	v[curr] = -1
	*r = append(*r, curr)
	return true
}
func findOrder(numCourses int, prereqs [][]int) []int {
	var res []int
	graph := make(map[int][]int)
	visited := make(map[int]int)

	for _, courses := range prereqs {
		courseToTake := courses[0]
		prereq := courses[1]
		graph[courseToTake] = append(graph[courseToTake], prereq)
	}

	for i := range graph {
		if !dfs2(i, visited, graph, &res) {
			return nil
		}

	}

	for i := range numCourses {
		if visited[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
