package main

import "fmt"

// https://leetcode.com/problems/single-number/submissions/
func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber2([]int{4, 1, 2, 1, 2}))

}

// using bit manipulation
func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result = result ^ v
	}
	return result
}

// using hashmap
func singleNumber2(nums []int) int {

	count := 0

	lookup := make(map[int]struct{}, len(nums)/2+1)

	for _, v := range nums {
		if _, exist := lookup[v]; exist {
			count -= v
			continue
		}
		lookup[v] = struct{}{}
		count += v

	}

	return count
}
