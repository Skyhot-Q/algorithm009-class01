package main

import (
	"fmt"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]struct{})
	for _, s := range wordList {
		if s == beginWord {
			continue
		}
		wordSet[s] = struct{}{}
	}

	if _, ok := wordSet[endWord]; !ok {
		return nil
	}

	startQueue := []string{beginWord}
	wordLen := len(beginWord)
	result := make([][]string, 0)
	for i := 0; len(startQueue) > 0; i++ {
		result = append(result, []string{})
		queueLen := len(startQueue)
		for i := 0; i < queueLen; i++ {
			if startQueue[i] == endWord {
				break
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
					result[i] = append(result[i], w)
					delete(wordSet, w)
				}
			}
		}
		startQueue = startQueue[queueLen:]
	}
	return result
}

func main() {
	//r := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	r := findLadders("a", "c", []string{"a", "b", "c"})
	fmt.Println("Result:", r)

}
