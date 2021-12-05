package main

import (
	"fmt"
)

func main() {

	fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
	fmt.Println(peakIndexInMountainArray2([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))

}

// using bisect left
func peakIndexInMountainArray(arr []int) int {

	lo := 0
	hi := len(arr)

	for lo < hi {
		mid := lo + (hi-lo)/2

		if arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// using normal binary tree in case unique
func peakIndexInMountainArray2(arr []int) int {

	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		}

		if arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}
