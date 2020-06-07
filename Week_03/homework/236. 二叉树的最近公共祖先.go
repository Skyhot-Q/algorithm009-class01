package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAnpcestor(root, p, q *TreeNode) *TreeNode {
	// T
	if root == nil {
		return nil
	}

	// P
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}

	// D
	l := lowestCommonAnpcestor(root.Left, p, q)
	r := lowestCommonAnpcestor(root.Right, p, q)

	// M
	if l != nil && r != nil {
		return root
	}

	if l == nil {
		return r
	}
	return l
}
