package main

import "fmt"

func f (nums[]int, target int) int {
	s := len(nums)
	s2 := s/2
	if s == 1 {
		if target != nums[0] {
			return -1
		}
	} else if s ==0 {
		return 0
	}

	if target == nums[s2] {
		left := -1
		right := -1
		for l := 0; l < s2; l++ {
			if target == nums[l] {
				left =l
				break
			}
		}
		for l := s-1; l > s2; l-- {
			if target == nums[l] {
				right = l
				break
			}
		}
		if right == -1 {
			right = s2
		}
		if left == -1 {
			left = s2
		}
		return right - left + 1
	} else if target < nums[s2] {
		return f(nums[0:s2], target)
	} else if target > nums[s2] {
		return f(nums[s2:s], target)
	} else {
		return -1
	}
}

func main() {
	x := f([]int{5,7,7,8,8,10}, 6)

	fmt.Println(x)
}