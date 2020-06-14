package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count :=0
	i := 0
	j := 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i]  {
			count++
			i++
		}
		j++
	}
	return count
}

func main() {
	testCases := [][][]int{
		{{1,2,3}, {1,1}},
		{{1,2}, {1,2,3}},
	}
	for _, c:= range testCases{
		r := findContentChildren(c[0], c[1])
		fmt.Println("Result:", r)
	}
}