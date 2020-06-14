package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left, right := 0, num/2
	for left <= right {
		x :=  left + (right - left)/2
		guess :=  x*x
		if guess == num {
			return true
		}
		if guess > num {
			right = x-1
		} else {
			left = x+1
		}
	}
	return false
}

// 牛顿迭代法
//func isPerfectSquare(num int) bool {
//	if num < 2{
//		return true
//	}
//
//	x := num/2
//	for x*x > num{
//		x  = (x+num/x) /2
//	}
//	return x*x == num
//}

func main() {
	testCases := []int{
		2147483647,
	}
	for _, c := range testCases {
		r := isPerfectSquare(c)
		fmt.Println("Result: ", r)
	}
}
