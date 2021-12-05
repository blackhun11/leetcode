package main

import "fmt"

func main() {
	// fmt.Println(search([]int{1, 2, 3, 4, 5}, 2))
	// fmt.Println(search2([]int{1, 2, 3, 4, 5}, 2))
	search3([]int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}, 1)

}

// [-1,0,3,5,9,12] l = 0 r = 5 m = 2
// [-1,0,3,5,9,12] l = 3 r = 5 m = 4

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1
}

func search2(nums []int, target int) int {

	index := bisectLeft(nums, target)

	if nums[index] != target {
		return -1
	}

	return index

}

func search3(nums []int, target int) int {

	index := bisectRight(nums, target)

	fmt.Println(index)
	return index

}

func bisectLeft(nums []int, target int) int {

	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2 // for handling overflow

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func bisectRight(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2 // for handling overflow

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
