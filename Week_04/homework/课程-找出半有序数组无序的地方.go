package main

import "fmt"

func searchTurnPoint(nums []int) int {
	nLen := len(nums)
	if nLen == 0 {
		return -1
	}
	left, right := 0, nLen-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid-1] > nums[mid] {
			return mid - 1
		}
		if nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] >= nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	testCases := [][]int{
		//{5,6,7,8,0,1,2,3,4},
		{5, 6, 7, 0, 1, 3},
		{},
		{1},
		{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5},
	}
	for _, c := range testCases {
		r := searchTurnPoint(c)
		fmt.Println("Result:", r)
	}
}
