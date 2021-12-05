package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := createList([]int{1, 2, 4})
	l2 := createList([]int{1, 3, 4})

	res := mergeTwoLists(l1, l2)

	arr := []int{}

	for res != nil {
		arr = append(arr, res.Val)
		res = res.Next
	}

	fmt.Println(arr)

}

func createList(arr []int) *ListNode {

	head := new(ListNode)

	temp := head

	for i := 0; i < len(arr); i++ {
		temp.Next = &ListNode{
			Val: arr[i],
		}

		temp = temp.Next
	}

	return head.Next

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := new(ListNode)

	temp := head

	for l1 != nil || l2 != nil {

		if l1 == nil {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
			continue
		}

		if l2 == nil {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
			continue
		}

		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next

		}
	}

	return head.Next
}
