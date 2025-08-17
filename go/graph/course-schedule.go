package main

func dfs(curr int, g map[int][]int, v map[int]int) bool {
	if v[curr] == 1 {
		return false
	}
	if v[curr] == -1 {
		return true
	}

	v[curr] = 1
	for _, nei := range g[curr] {
		if !dfs(nei, g, v) {
			return false
		}
	}
	v[curr] = -1
	return true
}

func canFinish(numCourses int, prereqs [][]int) bool {
	graph := make(map[int][]int)
	visited := make(map[int]int)

	for _, courseList := range prereqs {
		courseToTake := courseList[0]
		prereq := courseList[1]
		graph[prereq] = append(graph[prereq], courseToTake)
	}

	for i := range numCourses {
		if !dfs(i, graph, visited) {
			return false
		}
	}
	return true
}
