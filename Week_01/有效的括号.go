package main

import "fmt"

// https://leetcode-cn.com/problems/valid-parentheses/

// 第一遍： 2020/05/24 22:24
//func isValid(s string) bool {
// 符号映射表 -》 左右符合映射
// 顺序关闭 -》 stack

//symbol := "{}[]()"
//symbolMap := make(map[int32]int32)
//for i:=0; i<len(symbol); i+=2 {
//	fmt.Println(int32(symbol[i]), int32(symbol[i+1]))
//	symbolMap[int32(symbol[i])] = int32(symbol[i+1])
//}
//
//symbolMap := map[int32]int32{
//123: 125,
//91:93,
//40:41,
//}
//
//
//stack := make([]int32, 1)
//for _, x := range s {
//stackLen := len(stack)
//if x != symbolMap[stack[stackLen-1]] {
//stack = append(stack, x)
//} else {
//stack = stack[0:stackLen-1]
//}
//}
//if len(stack) > 1 {
//return false
//} else {
//return true
//}
//}

// 第二遍： 2020/05/26 21:43
func isValid(s string) bool {
	symbol := "(){}[]"
	symbolMap := make(map[int32]int32)
	var pre int32
	for i, v := range symbol {
		if i%2 == 0 {
			pre = v
		} else {
			symbolMap[pre] = v
		}
	}
	stack := make([]int32, 1, len(s))

	for _, v := range s {
		stackLen := len(stack)
		if v != symbolMap[stack[stackLen-1]] {
			stack = append(stack, v)
		} else {
			stack = stack[:stackLen-1]
		}
	}

	if len(stack) > 1 {
		return false
	} else {
		return true
	}
}
func main() {
	testCase := []string{
		"[](){}",
		"()",
		"{}",
		"[]",
		"[({",
	}
	for _, x := range testCase {
		r := isValid(x)
		fmt.Println(r)
	}
}
