package main

import "fmt"
func p(s string, isDeleted bool) bool {
	low := 0
	high := len(s) - 1
	for low = 0; low < high ; low ++{
		if s[low] != s[high] {
			if isDeleted{
				return false
			} else {
				return p(s[low:high], true) || p(s[low+1:high+1], true)
			}
		}
		high--
	}
	return true
}


func validPalindrome(s string) bool {
	return p(s, false)
}

func main() {

	x := validPalindrome("abc")
	fmt.Println(x)
}