package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Order struct {
	data []int
}

func (o *Order) preorder(node *TreeNode) {
	if node != nil {
		// root left right
		o.data = append(o.data, node.Val)
		o.preorder(node.Left)
		o.preorder(node.Right)
	}
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	o := Order{data: result}
	o.preorder(root)
	return o.data
}

func main() {
	testNode := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	r := preorderTraversal(testNode)
	fmt.Println("Result", r)
}
