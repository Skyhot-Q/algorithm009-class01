package main

import "strings"

func replaceSpace(s string) string {
	// 维护一个stack
	// 遇到空格清空，写入结果

	result := strings.Builder{}

	for  i:=0; i< len(s); i++ {
		if s[i] != ' '{
			result.WriteByte(s[i])
		} else {
			result.WriteByte('%')
			result.WriteByte('2')
			result.WriteByte('0')
		}
	}

	return result.String()
}
