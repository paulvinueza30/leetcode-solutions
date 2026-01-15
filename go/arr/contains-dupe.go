package main

import "fmt"

func containsDuplicate(arr []int) bool {
	check := 0
	for _, val := range arr {
		if (check & (1 << val)) != 0 {
			return true
		}
		check |= (1 << val)
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 2, 3}))
	fmt.Println(containsDuplicate([]int{1, 2, 2, 3}))
	fmt.Println(containsDuplicate([]int{1, 5, 2, 3, 3}))
	fmt.Println(containsDuplicate([]int{1, 5, 2, 3}))
}
