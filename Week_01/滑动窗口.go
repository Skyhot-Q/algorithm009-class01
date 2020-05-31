package main

import "fmt"

// https://leetcode-cn.com/problems/sliding-window-maximum/

//

// 单调递减栈： 压入的元素必须比当前栈顶小
// nums[i] > deque[k]， pop deque[k] ->  nums[i] > deque[k-1], pop deque[k-1;
// i - 3 - 1 > deque[0] , pop deque[0]
//
//i := 0

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	deque := make([]int, 0)
	res := make([]int, 0)
	for i, n := range nums {
		for len(deque) > 0 && n > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i-k >= deque[0] {
			deque = deque[1:]
		}

		if i-k+1 >=0 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

type testData struct {
	nums []int
	k    int
}

func main() {
	testCase := []testData{
		//{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
		{nums: []int{1, -1}, k: 1},
		{nums: []int{7, 2, 4}, k: 2},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
	}

	for _, c := range testCase {
		r := maxSlidingWindow(c.nums, c.k)
		fmt.Println("Result: ", r)
	}
}
