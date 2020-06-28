package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid) - 1
	if m < 0 {
		return -1
	}
	n := len(obstacleGrid[0]) - 1

	dp := make([][]int, m+1)
	for i := m; i >= 0; i-- {
		dp[i] = make([]int, n+1)
		for j := n; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if n == j && m == i {
				dp[i][j] = 1
				continue
			}

			if j == n {
				dp[i][j] = dp[i+1][j]
				continue
			}

			if i == m {
				dp[i][j] = dp[i][j+1]
				continue
			}

			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}
