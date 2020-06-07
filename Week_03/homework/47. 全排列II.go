package main

import (
	"fmt"
)

type Result struct {
	data [][]int
	uniq map[float64]struct{}
}

func (r *Result) search(nums []int, index int) {
	if index == len(nums) {
		fmt.Println("-->", nums)
		tmp := make([]int, len(nums), len(nums))
		copy(tmp, nums)
		r.data = append(r.data, tmp)
		return
	}

	m := make(map[int]int)
	for i := index; i < len(nums); i++ {
		//fmt.Println(i, index, nums)
		if _, ok := m[nums[i]]; ok {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		r.search(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
		m[nums[i]] = 1
	}
}

func permuteUnique(nums []int) [][]int {
	r := Result{make([][]int, 0), make(map[float64]struct{})}
	r.search(nums, 0)
	return r.data
}

func main() {
	testCase := [][]int{
		{1, 1, 2},
	}
	for _, c := range testCase {
		r := permuteUnique(c)
		fmt.Println("Result:", r)
	}
}
