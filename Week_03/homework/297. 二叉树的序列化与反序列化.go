package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	data []string
}

func Constructor() Codec {
	return Codec{
		data: []string{},
	}
}

func (this *Codec) s(root *TreeNode) {
	if root == nil {
		this.data = append(this.data, "null")
	} else {
		this.data = append(this.data, strconv.Itoa(root.Val))
		this.s(root.Left)
		this.s(root.Right)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.s(root)
	return "[" + strings.Join(this.data, ",") + "]"
}

func (this *Codec) d() *TreeNode {
	// t
	if len(this.data) == 0 {
		return nil
	}
	if this.data[0] == "null" {
		this.data = this.data[1:]
		return nil
	}

	// p
	val, _ := strconv.Atoi(this.data[0])
	t := &TreeNode{
		Val: val,
	}
	this.data = this.data[1:]
	// d
	t.Left = this.d()
	t.Right = this.d()
	return t
	// r
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1]
	this.data = strings.Split(data, ",")
	n := this.d()
	return n
}
