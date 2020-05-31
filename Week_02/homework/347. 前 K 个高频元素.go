package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode-cn.com/problems/chou-shu-lcof/

type PriorityQueue []*Item
type Item struct {
	value    int
	priority int
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	h := PriorityQueue{}
	x := make(map[int]int)
	for _, n := range nums {
		x[n] += 1
	}

	for k, v := range x {
		heap.Push(&h, &Item{value: k, priority: v})
	}

	ans := make([]int, 0)
	for k > 0 {
		ans = append(ans, heap.Pop(&h).(*Item).value)
		k--
	}
	return ans
}

type testData struct {
	arr []int
	k   int
}

func main() {
	testCase := []testData{
		{[]int{1,1,1,2,2,3}, 2},
	}
	for _, c := range testCase {
		r := topKFrequent(c.arr, c.k)
		fmt.Println("Result:", r)
	}
}
