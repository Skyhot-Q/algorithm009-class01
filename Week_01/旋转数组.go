package main

import (
	"fmt"
)

// 环转替换
//func rotate(nums []int, k int) {
//	numLen := len(nums)
//	k = k % numLen
//	var count int
//	for start := 0; count < numLen; start++ {
//		cur := start
//		prev := nums[start]
//		for {
//			next := (cur + k) % numLen
//			temp := nums[next]
//			nums[next] = prev
//			prev = temp
//			cur = next
//			count++
//			if start == cur {
//				break
//			}
//		}
//	}
//}

// 反转
func rotate(nums []int, k int) {
	numsLen := len(nums)
	k = k % numsLen
	reverse(nums, 0, numsLen-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, numsLen-1)
}

//func reverse(num []int, start int, end int) {
//	for start < end {
//		temp := num[start]
//		num[start] = num[end]
//		num[end] = temp
//		start++
//		end--
//	}
//}

func reverse(num []int, start int, end int) {
	for start < end {
		tmp := num[start]
		num[start] = num[end]
		num[end] = tmp
		start++
		end--
	}
}
func main() {
	testCase := []map[int][]int{
		{3: {1, 2, 3, 4, 5, 6, 7}},
		{3: {1, 2}},
	}
	for _, x := range testCase {
		for k, nums := range x {
			rotate(nums, k)
			fmt.Println("Result:", nums)
		}
	}
}
