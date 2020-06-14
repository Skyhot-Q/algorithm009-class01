package main

import (
	"fmt"
)
//
//func mySqrt(x int) int {
//	// 是单调递增，有上下界
//	if x==0|| x==1{
//		return x
//	}
//	left, right := 1, x
//	var mid int
//	for left <= right {
//		mid = left + (right-left)/2
//		if mid*mid == x {
//			return mid
//		}
//
//		if mid*mid > x {
//			right=mid -1
//		} else {
//			left= mid+1
//		}
//	}
//	return right
//}

func mySqrt(x int) int{
	if x < 2 {
		return x
	}
	cur := x
	for cur* cur > x {
		cur = (cur + x/cur)/2
	}
	return cur
}


func main(){
	testCases := []int {
		4,8,
	}
	for _,c := range testCases{
		r := mySqrt(c)
		fmt.Println("result:", r)
	}
}