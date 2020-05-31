package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/remove-outermost-parentheses/

func removeOuterParentheses(S string) string {

	// 外层的(在哪里?
	// 利用stack的特性，如果stack里面为空就代表里面是完整的（），可以记录入库
	// 也就是遇到)时而且栈顶为(时，pop出栈顶，直到 stack为空，从外层内的括号到当前下标的值可以记录入库
	// 其余情况都加入stack中

	if len(S) == 0 {
		return ""
	}
	//  这里是因为，题目时正确的括号，所以如果时)开头那就是永远都是错误的了，可以利用这点，将 stack和start初始化
	stack := []byte{S[0]}
	start := 1
	ans := strings.Builder{}
	for i := 1; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == '(' && S[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				ans.WriteString(S[start:i])
				start = i + 2
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return ans.String()
}

func main() {
	testCase := []string{
		"(()())(())",
		"(()())(())(()(()))",
		"()()",
		")(())(())(())",
	}

	for _, c := range testCase {
		r := removeOuterParentheses(c)
		fmt.Println("Result:", r)
	}

}
