package main

import "fmt"

// https://leetcode.com/problems/missing-number/
func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

// using hashmap
func missingNumber(nums []int) int {

	lookup := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if v == 0 {
			continue
		}

		lookup[v] = struct{}{}
	}

	for i := 1; i <= len(nums); i++ {
		if _, exist := lookup[i]; !exist {
			return i
		}
	}

	return 0
}

// using simple math
func missingNumber2(nums []int) int {

	totalV := 0
	totalK := 0
	for k, v := range nums {
		totalV += v
		totalK += k + 1
	}

	return totalK - totalV
}
