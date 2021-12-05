package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := []int{1, 2, 2, 1}

	head := new(ListNode)
	temp := head

	for i := 0; i < len(input); i++ {
		temp.Next = &ListNode{
			Val: input[i],
		}
		temp = temp.Next
	}
	res := middleNode(head.Next)

	arr := []int{}

	for res != nil {
		arr = append(arr, res.Val)
		res = res.Next
	}

	fmt.Println(arr)

}

/*
1 2 3 4 5
s s s
f   f   f

1 2 3 4 5 6
s s s s
f   f   f   f
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
