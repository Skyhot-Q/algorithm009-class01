package main

import (
	"container/heap"
	"fmt"
)

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

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

func getLeastNumbers(arr []int, k int) []int {
	h := IntHeap(arr)
	heap.Init(&h)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&h).(int)
	}
	return result
}

type testData struct {
	arr []int
	k   int
}

func main() {
	testCase := []testData{
		{[]int{2, 7, 11, 15}, 3},
		{[]int{3, 2, 1}, 2},
		{[]int{0, 1, 2, 1}, 2},
		{[]int{0, 1, 2, 1}, 1},
		{[]int{0,0,1,2,4,2,2,3,1,4}, 8},
	}
	for _, c := range testCase {
		r := getLeastNumbers(c.arr, c.k)
		fmt.Println("Result:", r)
	}
}
