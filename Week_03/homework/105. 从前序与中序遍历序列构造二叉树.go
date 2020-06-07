package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 先找出preorder根节点
	// 用inorder根据根节点划分左右树
	// 递归左右树，构建TreeNode
	// 要点：root在inorder的位置就代表了preorder有多少个值在左边，其余都是右边
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var root int
	for i, d := range inorder {
		if preorder[0] == d {
			root = i
			break
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:root+1], inorder[0:root]),
		Right: buildTree(preorder[root+1:], inorder[root+1:]),
	}
}

func main() {
	testCase := [][][]int{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
	}
	for _, c := range testCase {
		r := buildTree(c[0], c[1])
		fmt.Println("Result:", r)
	}
}
