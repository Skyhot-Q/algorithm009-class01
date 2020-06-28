package main

import "fmt"

func uniquePaths(m int, n int) int {
	// 1. 找出重复性
	// 2. 定义状态公式
	// 3. DP方程

	f := make([][]int, m)
	for i := m - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			if m-1 == i || n-1 == j {
				f[i][j] = 1
			} else {
				f[i][j] = f[i+1][j] + f[i][j+1]
			}
		}
	}
	return f[0][0]
}

func main() {
	r := uniquePaths(3, 2)
	fmt.Println("Result: ", r)
}
