package main

import "fmt"

//func moveZeroes(nums []int) {
//	slow := 0
//	for fast, n := range nums{
//		if n != 0 {
//			nums[slow] = n
//			if slow != fast {
//				nums[fast] = 0
//			}
//			slow++
//		}
//	}
//}

func moveZeroes(nums []int) {
	slow := 0
	for fast, n := range  nums {
		if n != 0 {
			nums[slow] = n
			if slow != fast {
				nums[fast] = 0
			}
			slow++
		}
	}
}
func main() {
	testCase := [][]int{
		{0, 1, 0, 3, 12},
		{3, 2, 4},
	}

	for _, d := range testCase {
		moveZeroes(d)
		fmt.Println("Result: ", d)
	}
}
