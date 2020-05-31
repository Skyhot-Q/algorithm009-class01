package main

import (
	"fmt"
)

// 第一遍： 2020/05/26 22:30 ->自己尝试习惯写法，效率不高，占用空间高
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t){
//		return false
//	}
//	sMap:= make(map[int32]int)
//	for _, i := range s{
//		sMap[i]+=1
//	}
//	for _, i := range t{
//		if c, ok := sMap[i]; ok {
//			if c==0 {
//				return false
//			} else {
//				sMap[i]--
//			}
//		} else {
//			return false
//		}
//	}
//	return true
//}

// 第一遍： 2020/05/26 22:30 -> 双百写法
func isAnagram(s string, t string) bool {
	// 利用字母转数字后减去 初始的'a'，得到的数字不超过26的特点，初始化一个26长度的数组做索引使用
	// 因保证长度一致下，遍历两个字符串，遍历s时将字母在数组中+1， 而遍历t时将字母在数组中-1
	// 最后遍历数组，判断是否全为0，是则true，否则false
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for i, v := range s {
		arr[v-'a']++
		arr[int32(t[i])-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

type testData struct {
	s string
	t string
}

func main() {
	testCase := []testData{
		{s: "anagram", t: "nagaram"},
		{s: "rat", t: "car"},
	}

	for _, c := range testCase {
		r := isAnagram(c.s, c.t)
		fmt.Println("Result:", r)
	}

}
