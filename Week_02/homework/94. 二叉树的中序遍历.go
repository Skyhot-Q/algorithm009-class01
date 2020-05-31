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

func (o *Order) inorder(node *TreeNode) {
	if node != nil {
		//left root right
		o.inorder(node.Left)
		o.data = append(o.data, node.Val)
		o.inorder(node.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	o := Order{data: result}
	o.inorder(root)
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
	r := inorderTraversal(testNode)
	fmt.Println("Result", r)
}
