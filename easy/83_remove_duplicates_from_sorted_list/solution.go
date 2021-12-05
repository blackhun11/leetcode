package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := []int{1, 1, 1, 1, 2, 2, 3, 3, 3}

	head := new(ListNode)

	temp := head

	for i := 0; i < len(input); i++ {
		temp.Next = &ListNode{
			Val: input[i],
		}
		temp = temp.Next
	}

	res := deleteDuplicates(head.Next)
	arr := []int{}

	for res != nil {
		arr = append(arr, res.Val)
		res = res.Next
	}

	fmt.Println(arr)

}

func deleteDuplicates(head *ListNode) *ListNode {

	val := 0

	temp := head

	for temp != nil && temp.Next != nil {
		val = temp.Val
		if val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return head

}
