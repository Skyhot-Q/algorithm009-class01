package main

import (
	"fmt"
)

func minMutation(start string, end string, bank []string) int {
	// A 变到 B 所需要的步数
	// 每次变化都要bank里面就有相应的值
	// 基因串每个字符只有四个选择 ACGT
	if start == end {
		return 0
	}

	stack := []string{start}
	used := make([]int, len(bank))
	nums := 0
	for len(stack) > 0 {
		nums++
		level := len(stack)
		for i := 0; i < level; i++ {
			for j, b := range bank {
				if b == start {
					continue
				}
				if used[j] != 0 {
					continue
				}
				diff := 0
				for m := 0; m < 8; m++ {
					if b[m]!=stack[i][m]{
						diff++
					}
					if diff>=2{
						break
					}
				}
				if diff == 1 {
					used[j]=nums
					if bank[j]==end{
						return nums
					}
					stack=append(stack,bank[j])
				}
			}
		}
		stack = stack[level:]
	}
	return -1
}

type testData struct {
	start string
	end   string
	bank  []string
}

func main() {
	testCases := []testData{
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}},
		{"AAAACCCC", "CCCCCCCC", []string{"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC"}},
	}

	for _, c := range testCases {
		r := minMutation(c.start, c.end, c.bank)
		fmt.Println("Result:", r)
	}
}
