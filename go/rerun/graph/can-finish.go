package graph

func canFinish(numCourses int, prereq [][]int) bool {
	adj := make(map[int][]int)
	indegree := make(map[int]int)
	for _, pair := range prereq {
		adj[pair[1]] = append(adj[pair[1]], pair[0])
		indegree[pair[0]] += 1
	}

	q := make([]int, 0)
	for c := range numCourses {
		if indegree[c] == 0 {
			q = append(q, c)
		}
	}
	count := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		count++
		for _, neigh := range adj[curr] {
			indegree[neigh]--
			if indegree[neigh] == 0 {
				q = append(q, neigh)
			}
		}

	}

	return count == numCourses
}
