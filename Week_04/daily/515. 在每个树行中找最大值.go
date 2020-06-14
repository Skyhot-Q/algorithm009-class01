package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result :=make([]int, 0)

	for i:=0; len(queue) > 0 ; i++{
		p := make([]*TreeNode, 0)
		max := math.MinInt64
		for _, q := range queue {
			if q.Val > max {
				max = q.Val
			}
			if q.Left != nil {
				p = append(p, q.Left)
			}
			if q.Right != nil {
				p = append(p, q.Right)
			}
		}
		queue = p
		result = append(result, max)
	}
	return result
}


func main() {
	x := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -1,
		},
	}
	r := largestValues(x)
	fmt.Println("Result: ", r)
}