package main

type Node struct {
	Val      int
	Children []*Node
}

func myPreorder(root *Node, result *[]int) {
	if root != nil {

		*result = append(*result, root.Val)
		for _, c := range root.Children {
			myPreorder(c, result)
		}
	}
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	myPreorder(root, &result)
	return result
}
