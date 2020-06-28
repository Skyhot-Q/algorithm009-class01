package main

import (
	"fmt"
)

/*
	1. BFS
	2. DP
		a. subproblems
		b. DP array： f(n)= min(f(n-k), for k in [1,2,5]) + 1
		c. DP方程
 */

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}

	return dp[amount]
}

func main() {
	testCases := []int{1, 2, 5}
	c := coinChange(testCases, 11)
	fmt.Println("Result:", c)
}
