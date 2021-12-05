package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := []int{3, 2, 0, -4}

	head := new(ListNode)
	temp := head

	sec := new(ListNode)

	for i := 0; i < len(input); i++ {
		temp.Next = &ListNode{
			Val: input[i],
		}
		temp = temp.Next
		if i == 1 {
			sec = temp
		}
	}

	temp.Next = sec

	res := hasCycle(head.Next)

	fmt.Println(res)

}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}
