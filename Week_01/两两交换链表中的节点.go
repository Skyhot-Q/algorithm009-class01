package main

import "fmt"

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h := &ListNode{Val: -1, Next: head}
	preHead := h
	cur := h.Next
	for cur != nil && cur.Next != nil {
		preHead.Next = cur.Next
		cur.Next = cur.Next.Next
		preHead.Next.Next = cur
		preHead = cur
		cur = cur.Next
	}
	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createListNodeFromIntArray(l []int) *ListNode {
	p := &ListNode{}
	result := p
	for _, i := range l {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}
	return result.Next
}

func printListNodeValue(l *ListNode) {
	p := l
	fmt.Print("ListNodeValue: start -> ")
	for p != nil {
		fmt.Print(p.Val, " -> ")
		p = p.Next
	}
	fmt.Println("end")
}

func main() {
	testCase := [][]int{
		{1, 2, 3, 0, 0, 4},
		{2,0},
	}
	for _, t := range testCase {
		r := swapPairs(createListNodeFromIntArray(t))
		printListNodeValue(r)
	}
}
