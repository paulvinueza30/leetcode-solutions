package main

func dfs3(curr int, g map[int][]int, v map[int]int, c *[]int) bool {
	v[curr] = 1
	*c = append(*c, curr)
	for _, nei := range g[curr] {
		if v[nei] == 0 {
			dfs3(nei, g, v, c)
		}
	}
	return true
}

func countComponents(n int, edges [][]int) int {
	res := 0
	g := make(map[int][]int)
	v := make(map[int]int)

	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	for i := range n {
		if v[i] != 0 {
			var c = []int{}
			dfs3(i, g, v, &c)
			complete := true
			for _, node := range c {
				if len(g[node]) != len(c)-1 {
					complete = false
				}

			}
			if complete {
				res++
			}
		}
	}

	return res
}
