package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type integerQueue []*int
type treeQueue []*TreeNode

/*
	  3
	9  20
	  15 7

	   1
	 2   3
   4  5 6  7
*/

func main() {
	input := []int{3, 9, 20, -1, -1, 15, 7}

	intQue := make(integerQueue, 0)
	treeQue := make(treeQueue, 0)

	for i := 1; i < len(input); i++ {
		if input[i] == -1 {
			intQue.push(nil)
			continue
		}
		intQue.push(&input[i])
	}

	root := &TreeNode{
		Val: input[0],
	}

	treeQue.push(root)

	for !intQue.isEmpty() {
		var left, right *int

		if !intQue.isEmpty() {
			left = intQue.pop()
		}

		if !intQue.isEmpty() {
			right = intQue.pop()
		}

		current := treeQue.pop()

		if left != nil {
			leftTree := &TreeNode{
				Val: *left,
			}
			current.Left = leftTree
			treeQue.push(leftTree)
		}

		if right != nil {
			rightTree := &TreeNode{
				Val: *right,
			}
			current.Right = rightTree
			treeQue.push(rightTree)
		}
	}

	// fmt.Println(averageOfLevels(root))
	root.bfs()
}

func (q *TreeNode) traverseInorderDFS() {
	if q == nil {
		return
	}

	q.Left.traverseInorderDFS()
	fmt.Print(q.Val, " ")
	q.Right.traverseInorderDFS()
}

func (q *TreeNode) traversePreorderDFS() {
	if q == nil {
		return
	}

	fmt.Print(q.Val, " ")
	q.Left.traversePreorderDFS()
	q.Right.traversePreorderDFS()
}

func (q *TreeNode) traversePostorderDFS() {
	if q == nil {
		return
	}

	q.Left.traversePostorderDFS()
	q.Right.traversePostorderDFS()
	fmt.Print(q.Val, " ")
}
func (q *integerQueue) push(i *int) {
	*q = append(*q, i)
}

func (q *integerQueue) pop() *int {

	arr := ([]*int(*q))
	res := arr[0]

	*q = arr[1:]

	return res
}

func (q *integerQueue) isEmpty() bool {
	return len(*q) == 0
}

func (q *treeQueue) isEmpty() bool {
	return len(*q) == 0
}

func (q *treeQueue) push(tree *TreeNode) {
	*q = append(*q, tree)
}

func (q *treeQueue) pop() *TreeNode {

	arr := ([]*TreeNode(*q))
	res := arr[0]

	*q = arr[1:]

	return res
}

func averageOfLevels(root *TreeNode) []float64 {
	sum := map[int]int{}
	count := map[int]int{}
	traverseDFS(root, sum, count, 0)

	res := make([]float64, len(sum))

	for i := 0; i < len(sum); i++ {
		res[i] = float64(sum[i]) / float64(count[i])
	}

	return res

}

func traverseDFS(root *TreeNode, sum, count map[int]int, level int) {

	if root == nil {
		return
	}

	count[level]++
	sum[level] += root.Val

	traverseDFS(root.Left, sum, count, level+1)
	traverseDFS(root.Right, sum, count, level+1)
}

func (q *TreeNode) bfs() {

	queue := treeQueue{}
	queue.push(q)

	for !queue.isEmpty() {
		current := queue.pop()
		fmt.Print(current.Val, " ")
		if current.Left != nil {
			queue.push(current.Left)
		}

		if current.Right != nil {
			queue.push(current.Right)
		}
	}

}
