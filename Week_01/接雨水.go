package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/trapping-rain-water/


// Date: 2020/05/23
//func trap(height []int) int {
//	heightLen := len(height)
//	if heightLen == 0 {
//		return 0
//	}
//	left := 0
//	right := heightLen - 1
//
//	ans := 0
//	leftMax := 0
//	rightMax := 0
//	for left != right {
//		if height[left] < height[right] {
//			if height[left] >= leftMax {
//				leftMax = height[left]
//			} else {
//				ans += (leftMax - height[left])
//			}
//			left++
//		} else {
//			if height[right] >= rightMax {
//				rightMax = height[right]
//			} else {
//				ans += (rightMax - height[right])
//			}
//			right--
//		}
//	}
//	return ans
//}

// Date: 2020/05/24 19:52
func trap(height []int) int {
	// left bound right bound
	// left max  right max
	// if height left < height right  --->
	// else <---
	// util 找到 height right > height left
	// util right == left break loop
	heightLen := len(height)
	if heightLen == 0 {
		return 0
	}
	leftMax := 0
	rightMax := 0
	leftBound := 0
	rightBound := heightLen - 1
	res := 0
	for leftBound < rightBound {
		if height[leftBound] < height[rightBound] {
			if height[leftBound] >= leftMax {
				leftMax = height[leftBound]
			} else {
				res += (leftMax - height[leftBound])
			}
			leftBound++

		} else {
			if height[rightBound] >= rightMax {
				rightMax = height[rightBound]
			} else {
				res += (rightMax - height[rightBound])
			}
			rightBound--
		}
	}

	return res
}

func main() {
	testCase := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{1, 0, 1, 0, 4, 0, 2},
	}
	for _, t := range testCase {
		result := trap(t)
		fmt.Println("Result:", result)
	}
}
