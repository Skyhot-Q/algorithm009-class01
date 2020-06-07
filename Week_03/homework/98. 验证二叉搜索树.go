package main

import (
	"fmt"
	"math"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidNode(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper{
		return false
	}
	return isValidNode(root.Left, lower, root.Val) && isValidNode(root.Right, root.Val, upper)
}

func isValidBST(root *TreeNode) bool {
	return isValidNode(root, math.MinInt64, math.MaxInt64)
}

func main() {
	r := isValidBST(&TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 5},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 20},
		},
	})
	fmt.Println(r)
}
