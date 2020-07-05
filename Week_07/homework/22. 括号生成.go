package main

import "fmt"

func dfs(ans *[]string, result string, n int, left, right int) {
	if left + right == n*2 {
		*ans = append(*ans, result)
		return
	}

	if left < n {
		dfs(ans, result+"(", n, left+1, right)
	}

	if right < left {
		dfs(ans, result+")", n, left, right+1)
	}
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	dfs(&ans, "", n, 0, 0)
	return ans
}

func main() {
	testCases := []int{
		3,
	}
	for _, c := range testCases {
		r := generateParenthesis(c)
		fmt.Println("Result:", r)
	}
}
