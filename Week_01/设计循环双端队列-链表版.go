package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

type MyCircularDeque struct {
	capacity int
	len      int
	head     *ListNode
	tail     *ListNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	d := MyCircularDeque{
		capacity: k,
		len:      0,
		head:     &ListNode{val: -1, prev: nil, next: nil},
		tail:     &ListNode{val: -1, prev: nil, next: nil},
	}
	d.head.next = d.tail
	d.tail.prev = d.head
	return d
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.len++
	n := &ListNode{val: value, prev: this.head, next: this.head.next}
	this.head.next.prev = n
	this.head.next = n
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.len++
	n := &ListNode{val: value, prev: this.tail.prev, next: this.tail}
	this.tail.prev.next = n
	this.tail.prev = n
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head.next = this.head.next.next
	this.head.next.prev = this.head
	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail.prev = this.tail.prev.prev
	this.tail.prev.next = this.tail
	this.len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	r := this.head.next.val
	return r
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	r := this.tail.prev.val
	return r
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.len == 0 {
		return true
	} else {
		return false
	}
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.len == this.capacity {
		return true
	} else {
		return false
	}
}

type dequeTestData struct {
	op    string
	value []int
}

func main() {
	testCase := [][]dequeTestData{
		{
			{"MyCircularDeque", []int{3}}, //首位必须为MyCircularDeque，初始化deque使用
			{"insertLast", []int{1}},
			{"insertLast", []int{2}},
			{"insertFront", []int{3}},
			{"insertFront", []int{4}},
			{"getRear", []int{}},
			{"isFull", []int{}},
			{"isEmpty", []int{}},
			{"deleteLast", []int{}},
			{"insertFront", []int{4}},
			{"getFront", []int{3}},
		},
	}

	var obj MyCircularDeque
	for _, t := range testCase {
		for _, d := range t {
			switch d.op {
			case "MyCircularDeque":
				obj = Constructor(d.value[0])

			case "insertLast":
				r := obj.InsertLast(d.value[0])
				fmt.Println("result: ", r)

			case "insertFront":
				r := obj.InsertFront(d.value[0])
				fmt.Println("result: ", r)

			case "getRear":
				r := obj.GetRear()
				fmt.Println("result: ", r)

			case "isFull":
				r := obj.IsFull()
				fmt.Println("result: ", r)

			case "isEmpty":
				r := obj.IsEmpty()
				fmt.Println("result: ", r)

			case "deleteLast":
				r := obj.DeleteLast()
				fmt.Println("result: ", r)

			case "deleteFront":
				r := obj.DeleteFront()
				fmt.Println("result: ", r)

			case "getFront":
				r := obj.GetFront()
				fmt.Println("result: ", r)
			}
		}
	}
}
