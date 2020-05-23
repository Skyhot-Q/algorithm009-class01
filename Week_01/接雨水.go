package main

import (
	"fmt"
)

func trap(height []int) int {
	heightLen := len(height)
	if heightLen == 0 {
		return 0
	}
	left := 0
	right := heightLen - 1

	ans := 0
	leftMax := 0
	rightMax := 0
	for left != right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += (leftMax - height[left])
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += (rightMax - height[right])
			}
			right--
		}
	}
	return ans
}

func main() {
	testCase := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}
	for _, t := range testCase {
		result := trap(t)
		fmt.Println("Result:", result)
	}
}
