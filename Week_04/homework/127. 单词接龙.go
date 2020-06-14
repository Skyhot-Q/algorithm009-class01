package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]struct{})
	for _, s := range wordList {
		if s == beginWord {
			continue
		}
		wordSet[s] = struct{}{}
	}

	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	wordLen := len(beginWord)
	step := 0
	for len(startQueue) > 0 && len(endQueue) > 0 {
		step++
		queueLen := len(startQueue)
		for i := 0; i < queueLen; i++ {
			if startQueue[i] == endWord {
				return step
			}
			for w, _ := range wordSet {
				c := 0
				for k := 0; k < wordLen; k++ {
					if startQueue[i][k] != w[k] {
						c++
					}
				}
				if c == 1 {
					startQueue = append(startQueue, w)
					delete(wordSet, w)
				}
			}
		}
		startQueue = startQueue[queueLen:]
	}
	return 0
}
//
//type item struct {
//	Val  string
//	Step int
//}
//
//func visitWordNode(queue *[]item, wordLen int, wordSet map[string][]string, visited map[string]int, otherVisited map[string]int) int {
//	wordQueue := *queue
//	word := wordQueue[0]
//	for i := 0; i < wordLen; i++ {
//		iWord := word.Val[:i] + "*" + word.Val[i+1:] // -> 枚举 26
//		for _, w := range wordSet[iWord] {
//			if v, ok := otherVisited[w]; ok {
//				return word.Step + v
//			}
//			if _, ok := visited[w]; !ok {
//				visited[w] = word.Step + 1
//				wordQueue = append(wordQueue, item{w, word.Step + 1})
//				*queue = wordQueue[1:]
//			}
//		}
//	}
//	return -1
//}
//
//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	wordLen := len(beginWord)
//	wordSet := make(map[string][]string)
//	isEnd := 0
//	for _, word := range wordList {
//		if word == beginWord {
//			continue
//		}
//		if word == endWord {
//			isEnd++
//		}
//		for i := 0; i < wordLen; i++ {
//			key := word[:i] + "*" + word[i+1:]
//			if _, ok := wordSet[key]; ok {
//				wordSet[key] = []string{word}
//			} else {
//				wordSet[key] = append(wordSet[key], word)
//			}
//		}
//	}
//	if isEnd == 0{
//		return 0
//	}
//
//	ans := -1
//
//	queueBegin := []item{{beginWord, 1}}
//	queueEnd := []item{{endWord, 1}}
//
//	visitedBegin := map[string]int{beginWord: 1}
//	visitedEnd := map[string]int{endWord: 1}
//
//	for len(queueBegin) > 0 && len(queueEnd) > 0 {
//		ans = visitWordNode(&queueBegin, wordLen, wordSet, visitedBegin, visitedEnd)
//		if ans != -1 {
//			return ans
//		}
//		ans = visitWordNode(&queueEnd, wordLen, wordSet, visitedEnd, visitedBegin)
//		if ans != -1 {
//			return ans
//		}
//	}
//
//	return 0
//}

func main() {
	//r := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	r := ladderLength("a", "c", []string{"a","b","c"})
	fmt.Println("Result:", r)

}
