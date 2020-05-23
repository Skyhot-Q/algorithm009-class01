package main

import (
	"fmt"
)

type MyCircularDeque struct {
	capacity int
	len      int
	queue    []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
		len:      0,
		queue:    make([]int, 0, k),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.len++
	this.queue = append([]int{value}, this.queue...)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.len++
	this.queue = append(this.queue, []int{value}...)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.queue = this.queue[1:this.len]
	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.len--
	this.queue = this.queue[0:this.len]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	f := this.queue[0]
	return f
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	r := this.queue[this.len-1]
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
	k        int
	firstVal []int
	lastVal  []int
}

func main() {
	testCase := []dequeTestData{
		{3, []int{1, 2, 3}, []int{4, 5, 6}},
	}

	for _, d := range testCase {
		obj := Constructor(d.k)
		param_1_1 := obj.InsertFront(d.firstVal[0])
		fmt.Println("param_1_1: ", param_1_1)
		param_1_2 := obj.InsertFront(d.firstVal[1])
		fmt.Println("param_1_1: ", param_1_2)

		param_2_1 := obj.InsertLast(d.lastVal[0])
		fmt.Println("param_2_1: ", param_2_1)
		param_2_2 := obj.InsertLast(d.lastVal[1])
		fmt.Println("param_2_2: ", param_2_2)
		param_2_3 := obj.InsertLast(d.lastVal[2])
		fmt.Println("param_2_3: ", param_2_3)

		param_3 := obj.DeleteFront()
		fmt.Println("param_3: ", param_3)

		param_4 := obj.DeleteLast()
		fmt.Println("param_4: ", param_4)

		param_5 := obj.GetFront()
		fmt.Println("param_5: ", param_5)

		param_1_1 = obj.InsertFront(d.firstVal[0])
		fmt.Println("param_1_1: ", param_1_1)
		param_1_2 = obj.InsertFront(d.firstVal[1])
		fmt.Println("param_1_1: ", param_1_2)
		param_6 := obj.GetRear()
		fmt.Println("param_6: ", param_6)

		param_7 := obj.IsEmpty()
		fmt.Println("param_7: ", param_7)

		param_8 := obj.IsFull()
		fmt.Println("param_8: ", param_8)

	}
}
