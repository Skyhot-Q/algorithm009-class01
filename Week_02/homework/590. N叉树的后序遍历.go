package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}
type Order struct {
	data []int
}

func (o *Order) postorder(node *Node) {
	if node != nil {
		for i:=0 ; i< len(node.Children); i++ {
			o.postorder(node.Children[i])
		}
		o.data = append(o.data, node.Val)
	}
}

func postorder(root *Node) []int {
	result := make([]int, 0)
	o := Order{result}
	o.postorder(root)
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
	r := postorder(testNode)
	fmt.Println("Result", r)
}
