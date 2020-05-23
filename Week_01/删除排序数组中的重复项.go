package main

import "fmt"

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	k := 0
	for i := 1; i < numsLen; i++ {
		if nums[k] != nums[i] {
			nums[k+1] = nums[i]
			k = k + 1
		}
	}
	return k + 1
}

func main() {
	testCase := [][]int{
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 2},
	}
	for _, nums := range testCase {
		x := removeDuplicates(nums)
		fmt.Println("Len:", x,", Array:", nums[0:x])
	}
}
