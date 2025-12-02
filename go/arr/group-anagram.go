package main

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[c-'a'] += 1
		}
		groups[count] = append(groups[count], str)
	}
	res := make([][]string, 0)
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}
