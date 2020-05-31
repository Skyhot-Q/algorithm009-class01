package main

import "fmt"

//https://leetcode-cn.com/problems/group-anagrams/solution/zi-mu-yi-wei-ci-fen-zu-by-leetcode/

// 第一遍： 2020/05/26 23:30 -> 利用26长度数组作为map的key保证唯一性，快速判断、检索，分类
func groupAnagrams(strs []string) [][]string {
	// 设定一个Map来用于记录 唯一key 与 字符串数组的对应关系
	// 唯一key采用26长度的int数组， 因为字母减去'a'的数字大小不会超过26, 而且golang中数组是值对象可以作为key存储
	// 遍历 strs的值，遍历值，转化出[26]int的数组key
	// 从map中寻找，找到则加入当前str值，否则在map中新增key并存储当前的str值进去
	k:=make(map[[26]int][]string)

	for _, str := range strs {
		key := [26]int{}
		for _, b := range str {
			key[b-'a']++
		}
		if x, ok := k[key]; ok {
			k[key]  = append(x, str)
		} else {
			k[key] = []string{str}
		}
	}

	res := make([][]string, 0, len(k))
	for _, j := range k{
		res = append(res, j)
	}
	return res
}

func main(){
 	testCase := [][]string{
 		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}

	for _, c := range testCase{
		r := groupAnagrams(c)
		fmt.Println("result:", r)
	}
}