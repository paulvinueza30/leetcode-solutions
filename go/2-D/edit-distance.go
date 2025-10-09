package main

func minDistance(word1 , word2 string) int {
	var dfs func( w1Idx int, w2Idx int ) int 
	dfs = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2){
			return len(word1) - i
		}
		if word1[i] == word2[j]{
			return dfs(i + 1, j + 1)
		}
		
		return 1 + min(dfs(i + 1, j), dfs(i, j + 1), dfs(i + 1, j + 1))
	}
	return dfs(0,0)
}

func minDistance2(word1, word2 string) int {
	cached := make(map[[2]int] int)
	var dfs func( w1Idx int, w2Idx int ) int 
	dfs = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2){
			return len(word1) - i
		}
		cache := [2]int{i, j}
		if v , ok := cached[cache]; ok{
			return v			
		}
		if word1[i] == word2[j]{
			return dfs(i + 1, j + 1)
		}
		
		cached[cache] = 1 + min(dfs(i + 1, j), dfs(i, j + 1), dfs(i + 1, j + 1))
		return cached[cache]
	}
	return dfs(0,0)
}