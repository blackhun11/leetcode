package main

import "fmt"

func main() {
	fmt.Println((productExceptSelf([]int{4, 5, 1, 8, 2})))
	fmt.Println((productExceptSelf2([]int{4, 5, 1, 8, 2})))

}

// brute force solution
// time complexity = o(n)2
// space complexity = o(n)
func productExceptSelf(nums []int) []int {

	l := len(nums)

	answer := make([]int, l)

	for i := 0; i < l; i++ {
		ans := 1
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}

			ans *= nums[j]
		}
		answer[i] = ans
	}

	return answer
}

// tricky solution, loop foward, and backward
// time complexity = o(m)
// space complexity = o(!)
func productExceptSelf2(nums []int) []int {
	l := len(nums)
	answer := make([]int, l)

	total := 1

	for i := 0; i < l; i++ {
		answer[i] = total
		total *= nums[i]
	}

	total = 1

	for i := l - 1; i >= 0; i-- {
		answer[i] *= total
		total *= nums[i]
	}

	return answer
}
