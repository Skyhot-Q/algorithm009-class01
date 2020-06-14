package main

import "fmt"

func lemonadeChange(bills []int) bool {
	p := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}
	for _, b := range bills {
		x := b - 5
		for x >= 20 && p[20] > 0 {
			x = x - 20
			p[20] -= 1
		}
		for x >= 10 && p[10] > 0 {
			x = x - 10
			p[10] -= 1
		}
		for x >= 5 && p[5] > 0 {
			x = x - 5
			p[5] -= 1
		}
		if x > 0 {
			return false
		}
		p[b] += 1
	}
	return true
}

func main() {
	testCase := [][]int{
		{10, 10},
		{5, 5, 5, 10, 20},
		{5, 5, 10},
	}
	for _, t := range testCase{
		r :=  lemonadeChange(t)
		fmt.Println("Result: ", r)
	}
}
