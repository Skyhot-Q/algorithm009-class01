package main

import (
	"fmt"
	"strings"
)


func wordPattern(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	recordMap := make(map[string]int32)
	uniqMap := make(map[int32]int)

	if len(pattern) != len(strList) {
		return false
	}

	for i, x := range strList {
		if s, ok := recordMap[x]; !ok {
			if _, ok := uniqMap[int32(pattern[i])]; ok {
				return false
			}
			recordMap[x] = int32(pattern[i])
			uniqMap[int32(pattern[i])] = 1
		} else {
			if s != int32(pattern[i]) {
				return false
			}
		}
	}
	return true
}


func main() {
	testCase := [][]string{
		{"abba", "dog cat cat dog"},
		{"abba", "dog cat cat fish"},
		{"aaaa", "dog cat cat dog"},
		{"abba", "dog dog dog dog"},
	}
	for _, t := range testCase {
		result := wordPattern(t[0], t[1])
		fmt.Println("Result:", result)
	}
}
