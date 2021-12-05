package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	input := []int{1, 2, 3, 4, 5}

	head := new(ListNode)

	temp := head

	for i := 0; i < len(input); i++ {
		head.Next = &ListNode{
			Val: input[i],
		}
		head = head.Next
	}

	res := reverseList(temp.Next)

	arr := []int{}

	for res != nil {
		arr = append(arr, res.Val)
		res = res.Next
	}

	fmt.Println(arr)
}

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
