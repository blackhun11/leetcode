package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	input := []int{1, 2, 6, 3, 4, 5, 6}

	head := new(ListNode)

	temp := head

	for i := 0; i < len(input); i++ {
		temp.Next = &ListNode{
			Val: input[i],
		}
		temp = temp.Next
	}
	res := removeElements2(head.Next, 6)

	arr := []int{}

	for res != nil {
		arr = append(arr, res.Val)
		res = res.Next
	}

	fmt.Println(arr)
}

func removeElements(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		head = head.Next
	}
	tmp := head

	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {

	newHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	tmp := newHead

	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return newHead.Next
}
