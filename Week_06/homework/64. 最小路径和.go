package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	// 1. 找出重复性
	// 2. 定义状态公式
	// 3. DP方程
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j != len(grid[0])-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]

			} else if j == len(grid[0])-1 && i != len(grid)-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]

			} else if j != len(grid[0])-1 && i != len(grid)-1 {
				x := grid[i+1][j]
				y := grid[i][j+1]
				if x < y {
					grid[i][j] = grid[i][j] + x
				} else {
					grid[i][j] = grid[i][j] + y
				}
			}
		}
	}
	return grid[0][0]
}


func main() {
	testCases := [][][]int{
		{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
			//{1, 2},
			//{1, 2},
		},
	}
	for _, c := range testCases {
		r := minPathSum(c)
		fmt.Println("result:", r)
	}
}
