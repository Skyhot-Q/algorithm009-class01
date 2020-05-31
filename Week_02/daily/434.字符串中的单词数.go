package main

import "fmt"

func countSegments(s string) int {
	word := make([]int32, 0, len(s))
	wc := 0
	for _, x := range s {
		if x != ' ' {
			word = append(word, x)
		} else {
			if len(word) > 0 {
				wc++
				word = word[0:0]
			}
		}
	}
	if len(word) > 0 {
		wc ++
	}
	return wc
}

func main() {
	testCase := []string{
		"Hello, my name is John",
	}
	for _, c := range testCase {
		r := countSegments(c)
		fmt.Println("Result:", r)
	}
}
