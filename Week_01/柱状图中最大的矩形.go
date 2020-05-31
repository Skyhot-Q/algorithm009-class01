package main

import "fmt"

func largestRectangleArea(heights []int) int {
	heights = append([]int{-2}, heights...)
	heights = append(heights, []int{-1}...)
	heightsLen := len(heights)
	stack := []int{0}
	res := 0
	i := 1
	for i < heightsLen {
		stackLen := len(stack)
		top := stackLen - 1
		if heights[i] > heights[stack[top]] {
			stack = append(stack, i)
			i++
		} else {
			temp := heights[stack[top]] * (i - stack[top-1] - 1)
			stack = stack[0:top]
			if temp > res {
				res = temp
			}
		}
	}
	return res
}

func main() {
	testCase := [][]int{
		{2, 1, 5, 6, 2, 3},
		{0},
		{},
		{1},
	}

	for _, d := range testCase {
		r := largestRectangleArea(d)
		fmt.Println("Result: ", r)
	}
}
