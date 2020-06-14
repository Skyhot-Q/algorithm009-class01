package main

import "fmt"

//
//func searchMatrix(matrix [][]int, target int) bool {
//	for _, x := range matrix {
//		if len(x) == 0 {
//			continue
//		}
//		if target > x[len(x)-1] {
//			continue
//		}
//		left, right, mid := 0, len(x)-1, 0
//		for left <= right {
//			mid = left + (right-left)/2
//			if target == x[mid] {
//				return true
//			}
//			if target >= x[mid] {
//				left = mid + 1
//			} else {
//				right = mid - 1
//			}
//		}
//		return false
//	}
//	return false
//}

func searchMatrix(matrix [][]int, target int) bool {
	mLen := len(matrix)
	if mLen == 0 {
		return false
	}
	nLen := len(matrix[0])
	left, right := 0, mLen*nLen-1
	mid := 0
	for left <= right {
		mid = left + (right-left)/2
		c := mid % nLen
		r := mid / nLen
		if target == matrix[r][c] {
			return true
		}
		if target > matrix[r][c] {
			left = mid +1
		} else {
			right = mid -1
		}
	}
	return false
}

type testItem struct {
	matrix [][]int
	target int
}

func main() {
	testCases := []testItem{
		{[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		}, 3},
	}

	for _, c := range testCases {
		r := searchMatrix(c.matrix, c.target)
		fmt.Println("Result:", r)
	}
}
