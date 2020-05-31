package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

type Order struct {
	data []int
}

func (o *Order) preorder(node *Node) {
	if node != nil {
		o.data = append(o.data, node.Val)
		for i:=0 ; i< len(node.Children); i++ {
			o.preorder(node.Children[i])
		}
	}
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	o := Order{result}
	o.preorder(root)
	return o.data
}

func main() {
	testNode := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}},
			{Val: 2},
			{Val: 4},
		},
	}
	r := preorder(testNode)
	fmt.Println("Result", r)
}
