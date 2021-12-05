package main

import "fmt"

func main() {
	Constructor([]int{-4, -5})
	// fmt.Println(obj.SumRange(0, 2))
	// fmt.Println(obj.SumRange(2, 5))
	// fmt.Println(obj.SumRange(0, 5))

	construct([]int{-4, -5})
	// fmt.Println(obj2.SumRange(0, 2))
	// fmt.Println(obj2.SumRange(2, 5))
	// fmt.Println(obj2.SumRange(0, 5))

}

type NumArray struct {
	num []int
}

// DP bottom UP
func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	fmt.Println(arr)

	return NumArray{
		num: arr,
	}
}

func construct(nums []int) NumArray {

	total := 0
	for _, v := range nums {
		total += v
	}

	res := make([]int, len(nums))
	topDown(nums, res, total, len(nums)-1)
	fmt.Println(res)
	return NumArray{
		num: res,
	}

}

func topDown(nums, res []int, total, idx int) {
	res[idx] = total

	if idx == 0 {
		return
	}
	total -= nums[idx]
	idx--
	topDown(nums, res, total, idx)
}

func (this *NumArray) SumRange(left int, right int) int {

	if left == 0 {
		return this.num[right]
	}

	return this.num[right] - this.num[left-1]
}
