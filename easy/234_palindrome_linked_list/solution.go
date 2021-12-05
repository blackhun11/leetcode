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

	res := isPalindrome(head.Next)

	fmt.Println(res)

}

/*
  	1 2 2 1
	s s s
    p p
	f   f   f
	1 2 3 2 1
	s s s s
	p p
	f   f   f
*/

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	var prev *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// cut first to half
	prev.Next = nil

	// odd
	if fast != nil {
		slow = slow.Next
	}
	// reverse first to half
	reversedHead := reverse(head)

	// compare first to half with half to end
	for reversedHead != nil && slow != nil {

		if reversedHead.Val != slow.Val {
			return false
		}

		reversedHead = reversedHead.Next
		slow = slow.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {

	var prev *ListNode = nil

	for head != nil {
		current := head.Next
		head.Next = prev
		prev = head
		head = current
	}

	return prev
}
