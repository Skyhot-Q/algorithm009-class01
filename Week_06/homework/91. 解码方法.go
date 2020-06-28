package main

func numDecodings(s string) int {
	pre, cur := 1, 1
	if s[0] == '0' {
		return 0
	}

	for i := 1; i < len(s); i++ {
		temp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' {
			cur = cur + pre
		} else if s[i-1] == '2' {
			if s[i] < '7' {
				cur = cur + pre
			}
		}
		pre = temp
	}
	return cur
}