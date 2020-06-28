package main

import (
	"fmt"
	"math"
)


func maxSumSubmatrix(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	max := math.MinInt64
	for i1 := 1; i1 <= rows; i1++ {
		for j1 := 1; j1 <= cols; j1++ {
			dp := make([][]int, rows+1)
			for r := 0; r < rows+1; r++ {
				dp[r] = make([]int, cols+1)
			}
			dp[i1][j1] = matrix[i1-1][j1-1]
			for i2 := i1; i2 <= rows; i2++ {
				for j2 := j1; j2 <= cols; j2++ {
					dp[i2][j2] = dp[i2-1][j2] + dp[i2][j2-1] - dp[i2-1][j2-1] + matrix[i2-1][j2-1]
					if dp[i2][j2] <= k && dp[i2][j2] > max {
						max = dp[i2][j2]
					}
				}
			}
		}
	}
	return max
}



func main() {
	r := maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2)
	fmt.Println("Result:", r)
}
