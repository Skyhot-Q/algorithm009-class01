package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 枚举target
	// 通过left，right 的边界移动来进行值记录
	// 注意除重是否完成

	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}
	result := make([][]int, 0)

	sort.Ints(nums)
	for i, num := range nums[:numsLen-2] {
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1]{
			continue
		}
		target := -num
		left := i+1
		right := numsLen - 1
		for left < right {
			tmpNum := nums[left] + nums[right]
			if tmpNum == target{
				result = append(result, []int{num, nums[left], nums[right]})
				left ++
				right --
				for right > left && nums[right] == nums[right+1]{
					right --
				}

				for left <= right && nums[left] == nums[left-1]{
					left ++
				}
			} else if tmpNum > target {
				right--
			} else if tmpNum < target {
				left++
			}
		}
	}
	return result
}

func main () {
	testCase := [][]int{
		{-1, 0, 1, 2, -1, -4},
	}

	for _, c := range testCase {
		 r := threeSum(c)
		 fmt.Println("result:", r)
	}
}

