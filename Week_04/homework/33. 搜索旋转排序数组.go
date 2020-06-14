package main

import "fmt"

func search(nums []int, target int) int {
	//二分查找法，判断方式采用mid、左右边界对比，判断target所处那边进行查找
	numLen := len(nums)
	if numLen == 0 {
		return -1
	}
	left, right, mid := 0, numLen-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

type testItem struct {
	nums   []int
	target int
}

func main() {
	testCases := []testItem{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for _, c := range testCases {
		r := search(c.nums, c.target)
		fmt.Println("Result:", r)
	}
}
