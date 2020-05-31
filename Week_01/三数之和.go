package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/3sum/

//func threeSum(nums []int) [][]int {
//	// 排序
//	// 枚举target
//	// left bound， right bound
//	// 如果left - right
//	// 记得去重
//	result := make([][]int, 0)
//	numsLen := len(nums)
//	if numsLen < 3 {
//		return result
//	}
//	sort.Ints(nums)
//	for i := 0; i < numsLen-2; i++ {
//		if nums[i] > 0 {
//			break
//		}
//		if i > 0 && nums[i-1] == nums[i] {
//			continue
//		}
//		target := -nums[i]
//		left := i + 1
//		right := numsLen - 1
//		for left < right {
//			sum := nums[left] + nums[right]
//			if sum == target {
//				result = append(result, []int{nums[i], nums[left], nums[right]})
//				right--
//				left++
//				for left < right && nums[right+1] == nums[right] {
//					right--
//				}
//				for left < right && nums[left-1] == nums[left] {
//					left++
//				}
//			} else if sum > target {
//				right--
//			} else {
//				left++
//			}
//		}
//	}
//	return result
//}

// Date: 2020/05/25 23:00
func threeSum(nums []int) [][]int {
	nLen := len(nums)
	sort.Ints(nums)

	result := make([][]int, 0)
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && nums[i-1] == v {
			continue
		}
		target := -v
		left := i + 1
		right := nLen - 1
		for left < right {
			n := nums[left] + nums[right]
			if n == target {
				result = append(result, []int{v, nums[left], nums[right]})
				right--
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if n < target {
				left++
			} else if n > target {
				right--
			}
		}
	}
	return result
}

func main() {
	testCase := [][]int{
		{2, 1, 5, 6, 2, 3},
		{-1, 0, 1, 2, -1, -4},
		{},
		{0, 0, 0},
		{1},
		{1, 2, -2, -1},
		{-13, 5, 13, 12, -2, -11, -1, 12, -3, 0, -3, -7, -7, -5, -3, -15, -2, 14, 14, 13, 6, -11, -11, 5, -15, -14, 5, -5, -2, 0, 3, -8, -10, -7, 11, -5, -10, -5, -7, -6, 2, 5, 3, 2, 7, 7, 3, -10, -2, 2, -12, -11, -1, 14, 10, -9, -15, -8, -7, -9, 7, 3, -2, 5, 11, -13, -15, 8, -3, -7, -12, 7, 5, -2, -6, -3, -10, 4, 2, -5, 14, -3, -1, -10, -3, -14, -4, -3, -7, -4, 3, 8, 14, 9, -2, 10, 11, -10, -4, -15, -9, -1, -1, 3, 4, 1, 8, 1},
	}

	for _, d := range testCase {
		r := threeSum(d)
		fmt.Println("Result: ", r)
	}
}
