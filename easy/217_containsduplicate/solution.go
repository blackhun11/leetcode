package main

import "fmt"

// https://leetcode.com/problems/contains-duplicate/
func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

// using map
func containsDuplicate(nums []int) bool {

	lookup := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, exist := lookup[v]; exist {
			return true
		}

		lookup[v] = struct{}{}
	}

	return false
}
