package main

import "fmt"
//
//func findMin(nums []int) int {
//	nlen := len(nums)
//	if nlen == 1 {
//		return nums[0]
//	}
//	if nlen == 0 {
//		return -1
//	}
//	right, left, mid := nlen-1, 0, 0
//
//	if nums[right] > nums[0] {
//		return nums[0]
//	}
//
//	for left <= right {
//		mid = left + (right-left)/2
//
//		if nums[mid] > nums[mid+1] {
//			return nums[mid+1]
//		}
//
//		if nums[mid-1] > nums[mid] {
//			return nums[mid]
//		}
//
//		if nums[mid] > nums[0] {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//	return -1
//}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right :=  0, len(nums) -1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	testCases := [][]int{
		{3, 4, 5, 1, 2},
		{3, 1, 2},
		{1},
		{1,2},
		{2,1},
		{2,0,1},
		{},
	}
	for _, c := range testCases {
		r := findMin(c)
		fmt.Println("Result:", r)
	}
}
