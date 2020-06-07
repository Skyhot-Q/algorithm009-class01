package main

import "fmt"

func dfs(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		tmp := make([]int, len(nums), len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, res, index+1)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

func permute(nums []int) [][]int {
	res := make([][]int,0)
	dfs(nums, &res, 0)
	return res
}

func main() {
	testCase := [][]int{
		{1, 2, 3},
	}
	for _, c := range testCase {
		r := permute(c)
		fmt.Println("Result:", r)
	}
}
