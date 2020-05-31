package main

import (
	"container/heap"
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/chou-shu-lcof/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func nthUglyNumber(n int) int {
	// 维护一个小顶堆
	// 初始化小顶堆的初始值为1
	// 循环n，当n时退出，并输出uglyNumber
	// 每次都会将2、3、5 各自与uglyNumber相乘，并存放进堆，达到堆顶为最小值以便下次循环使用
	h := IntHeap{1}
	heap.Init(&h)
	uglyNum := -1
	for n > 0 {
		for h[0] == uglyNum {
			fmt.Println(h[0], uglyNum)
			heap.Pop(&h)
		}
		uglyNum = heap.Pop(&h).(int)
		if uglyNum*2 <= math.MaxInt32 {
			heap.Push(&h, uglyNum*2)
		}
		if uglyNum*3 <= math.MaxInt32 {
			heap.Push(&h, uglyNum*3)
		}
		if uglyNum*5 <= math.MaxInt32 {
			heap.Push(&h, uglyNum*5)
		}
		n--
	}
	return uglyNum
}

type testData struct {
	arr []int
	k   int
}

func main() {
	testCase := []int{
		//19,
		103,
	}
	for _, c := range testCase {
		r := nthUglyNumber(c)
		fmt.Println("Result:", r)
	}
}
