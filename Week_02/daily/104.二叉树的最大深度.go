package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// root 向下衍生，使用前序遍历如果左右节点都为空时跳出，每次往下加一，时间复杂度O(N)
	if root == nil {
		return 0
	}
	leftDepth :=  float64(maxDepth(root.Left))
	rightDepth := float64(maxDepth(root.Right))

	return int(math.Max(leftDepth, rightDepth) + 1)
}

func createTreeNode(l []int) *TreeNode {
	n := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	return n
}

func main() {
	p := createTreeNode([]int{3, 9, 20, 15, 7})
	r := maxDepth(p)
	fmt.Println(r)
}
