package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}

// greedy algorithm
func maxSubArray(nums []int) int {

	max := nums[0]

	tempMax := max

	for i := 1; i < len(nums); i++ {

		max += nums[i]

		if nums[i] > max {
			max = nums[i]
		}

		if max > tempMax {
			tempMax = max
		}
	}

	return tempMax
}

// brute force
func maxSubArray2(nums []int) int {

	max := math.MinInt32

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

// using kadane algorithm
/*
Kadane’s algorithm works by maintaining the start position of a subarray and repeatedly looking
at the next element in the array and deciding to either.
*/

// -2,1,4,5,6
// [-2,1,-3,4,-1,2,1,-5,4]
// [-2,1,-2,4,3,5,6,1,5]
func maxSubArray3(nums []int) int {

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// divide and conquer
func maxSubArray4(nums []int) int {
	return maxSubArraySum(nums, 0, len(nums)-1)
}

func maxSubArraySum(arr []int, left, right int) int {

	if left == right {
		return arr[left]
	}

	middle := (left + right) / 2

	return max(
		max(maxSubArraySum(arr, left, middle),
			maxSubArraySum(arr, middle+1, right)), maxCrossingSum(arr, left, middle, right),
	)
}

func maxCrossingSum(arr []int, left, middle, right int) int {

	sum := 0
	leftSum := math.MinInt32

	for i := middle; i >= left; i-- {
		sum = sum + arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	rightSum := math.MinInt32

	for i := middle + 1; i <= right; i++ {
		sum = sum + arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return max(leftSum+rightSum, max(leftSum, rightSum))

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
