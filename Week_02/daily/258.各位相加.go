package main

import "fmt"

func addDigits(num int) int {
	for num>= 10 {
		fmt.Println(num%10, num/10)
		num = num%10 +num/10

	}
	return num
}


func main() {
	testCase := []int{
		38,
		49,
		3801020,
		80,
		10,
	}
	for _, c := range testCase {
		x := addDigits(c)
		fmt.Println(x)
	}
}