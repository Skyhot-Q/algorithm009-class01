package main

import "fmt"

//https://leetcode-cn.com/problems/two-sum/description/

// 第一遍： 2020/05/20 23:50
func twoSum(nums []int, target int) []int {

	// record作为历史数据缓存字典记录下标与值的对应关系
	// 利用  a + b= target ==> target - a = b 公式
	// 使用 b 在 历史缓存字典中查找，找到则返回
	// 否则将 a 和他的 i 记录到历史数据缓存中

	record := make(map[int]int)
	for i, a := range nums {
		b := target - a
		if v, ok := record[b]; ok {
			return []int{v, i}
		} else {
			record[a] = i
		}
	}
	return []int{-1, -1}
}

type testData struct {
	nums   []int
	target int
}

func main() {
	testCase := []testData{
		{[]int{2, 7, 11, 15}, 9},
	}
	for _, c := range testCase {
		r := twoSum(c.nums, c.target)
		fmt.Println("Result:", r)
	}
}
