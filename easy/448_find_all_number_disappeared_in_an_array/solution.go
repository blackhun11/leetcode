package main

import (
	"math"
)

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
func main() {
	findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
}

func findDisappearedNumbers(nums []int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		val := math.Abs(float64(nums[i]))

		if nums[int(val)-1] > 0 {
			nums[int(val)-1] = -nums[int(val)-1]
		}
	}
	for k, v := range nums {
		if v > 0 {
			res = append(res, k+1)
		}
	}
	return res
}
