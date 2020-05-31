package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}
type Order struct {
	level [][]int
}

func (o *Order) levelOrder(node *Node, layer int) {
	if node != nil {
		if len(o.level) == layer {
			o.level = append(o.level, []int{})
		}
		o.level[layer] = append(o.level[layer], node.Val)
		for i := 0; i < len(node.Children); i++ {
			o.levelOrder(node.Children[i], layer+1)
		}
	}
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	o := Order{result}
	o.levelOrder(root, 0)
	return o.level
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
	r := levelOrder(testNode)
	fmt.Println("Result", r)
}
