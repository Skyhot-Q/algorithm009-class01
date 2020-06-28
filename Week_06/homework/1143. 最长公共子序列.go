package main

func longestCommonSubsequence(text1 string, text2 string) int {
	/*
		  " a b c d e
		" 0 0 0 0 0 0
		a 0 1 1 1 1 1
		c 0 1 1 2 2 2
		e 0 1 1 2 2 3
	*/

	// 找重复性
	// 定义关系
	// DP方程

	// 为什么是要 len(text)+1，其实是因为考虑空字符串的时候
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
