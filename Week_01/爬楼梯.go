package main

import (
	"fmt"
	"math"
)

// 记忆迭代
//func climbStair(i, n int, memo []int) int {
//	if i == n {
//		return 1
//	}
//	if i > n {
//		return 0
//	}
//
//	if memo[i] > 0 {
//		return memo[i]
//	}
//
//	memo[i] = climbStair(i+1, n, memo) + climbStair(i+2, n, memo)
//	return memo[i]
//}
//
//func ClimbStairs(n int) int{
//	memo := make([]int, n)
//	return climbStair(0, n, memo)
//}



//斐波那契公式
func ClimbStairs(n int) int{
	sqrt5 := math.Sqrt(5)
	n1 := float64(n+1)
	x := (math.Pow(((1+sqrt5)/2) , n1) - math.Pow(((1-sqrt5)/2), n1))/sqrt5
	return int(math.Round(x))
}

//
//func ClimbStairs(n int) int {
//	if n == 1 {
//		return n
//	}
//	dp := make([]int, n+1)
//	dp[1] = 1
//	dp[2] = 2
//	for i := 3; i <= n; i++ {
//		dp[i] = dp[i-1] + dp[i-2]
//	}
//	return dp[n]
//}



func main(){
	x := ClimbStairs(2)
	fmt.Printf("count: %d\n", x)
}