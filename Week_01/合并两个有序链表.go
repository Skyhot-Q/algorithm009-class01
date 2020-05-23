package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			preHead.Next = l2
			l2 = l2.Next
		} else {
			preHead.Next = l1
			l1 = l1.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return result.Next
}

type testData struct {
	l1 []int
	l2 []int
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
	testCase := []testData{
		{[]int{1, 2, 4}, []int{1, 3, 4}},
		{[]int{2, 2, 5}, []int{1, 3, 4, 9}},
		{[]int{}, []int{1, 3, 4, 9}},
		{[]int{1, 4, 5}, []int{}},
		{nil, []int{1, 2, 3, 4}},
	}
	for _, x := range testCase {
		l1 := createListNodeFromIntArray(x.l1)
		l2 := createListNodeFromIntArray(x.l2)
		l := mergeTwoLists(l1, l2)
		printListNodeValue(l)
	}
}
