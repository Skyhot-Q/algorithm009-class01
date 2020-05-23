package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//a + b =target
	record := make(map[int]int)
	for i, n := range nums {
		t := target - n
		if j, ok := record[t]; ok {
			return []int{j, i}
		} else {
			record[n] = i
		}
	}
	return []int{-1, -1}
}

type twoSumTestData struct {
	nums   []int
	target int
}

func main() {
	testCase := []twoSumTestData{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
	}

	for _, d := range testCase {
		result := twoSum(d.nums, d.target)
		fmt.Println("Result: ", result)
	}
}
