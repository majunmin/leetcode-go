package leetcode_0127

import (
	"math"
)

var (
	minLen = math.MaxInt
)

// https://leetcode.cn/problems/word-ladder/description/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	return doubleBfsSolution(beginWord, endWord, wordList)
}

func BFSSolution2(beginWord string, endWord string, wordList []string) int {
	// param check
	if beginWord == endWord {
		return 0
	}
	step := 0
	var queue []string
	set := make(map[string]bool)
	for _, word := range wordList {
		if word == beginWord {
			continue
		}
		set[word] = true
	}
	if _, exist := set[endWord]; !exist {
		return 0
	}

	queue = append(queue, beginWord)
	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := 0; i < size; i++ {
			curWord := queue[i]
			if curWord == endWord {
				return step
			}
			for next := range set {
				if !diffOne(curWord, next) {
					continue
				}
				if set[next] {
					queue = append(queue, next)
					set[next] = false
				}
			}
		}
		queue = queue[size:]
	}
	return 0
}

func BFSSolution(beginWord string, endWord string, wordList []string) int {
	// param check
	if beginWord == endWord {
		return 0
	}
	step := 0
	var queue []string
	visited := make(map[string]bool)
	visited[beginWord] = true

	queue = append(queue, beginWord)
	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := 0; i < size; i++ {
			curWord := queue[i]
			if curWord == endWord {
				return step
			}
			for _, next := range wordList {
				if visited[next] {
					continue
				}
				if !diffOne(curWord, next) {
					continue
				}
				// terminate
				visited[next] = true
				queue = append(queue, next)
			}
		}
		queue = queue[size:]
	}
	return 0
}

func doubleBfsSolution(beginWord string, endWord string, wordList []string) int {
	// param check
	if beginWord == endWord {
		return 0
	}
	if len(wordList) == 0 {
		return 0
	}

	wordSet := make(map[string]bool)
	for _, word := range wordList {
		if word == beginWord {
			continue
		}
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	wordSet[beginWord] = false
	wordSet[endWord] = false
	step := 1
	for len(startQueue) > 0 && len(endQueue) > 0 {
		startSet := make(map[string]bool, len(startQueue))
		endSet := make(map[string]bool, len(endQueue))
		for _, word := range startQueue {
			startSet[word] = true
		}
		for _, word := range endQueue {
			endSet[word] = true
		}
		size := len(startQueue)
		for i := 0; i < size; i++ {
			curWord := startQueue[i]

			for idx := range curWord {
				for j := 0; j < 26; j++ {
					next := curWord[0:idx] + string(rune('a'+j)) + curWord[idx+1:]
					// terminate
					if endSet[next] {
						return step + 1
					}

					if wordSet[next] {
						startQueue = append(startQueue, next)
						wordSet[next] = false
					}
				}
			}

		}
		step++
		startQueue = startQueue[size:]
		// 选择短的一遍进行遍历
		if len(endQueue) < len(startQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return 0

}

// DFSSolution
func DFSSolution(beginWord string, endWord string, wordList []string) int {
	visited := make(map[string]bool)
	dfs(beginWord, endWord, wordList, visited, 1)
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func dfs(beginWord string, endWord string, wordList []string, visited map[string]bool, step int) {
	// terminate
	if beginWord == endWord {
		if step < minLen {
			minLen = step
		}
		return
	}

	for _, word := range wordList {
		if visited[word] {
			continue
		}
		if !diffOne(beginWord, word) {
			continue
		}

		visited[word] = true
		dfs(word, endWord, wordList, visited, step+1)
		visited[word] = false
	}
}

func diffOne(word string, word2 string) bool {
	if word == word2 {
		return false
	}
	cnt := 0
	for i := 0; i < len(word); i++ {
		if word[i] != word2[i] {
			cnt++
			if cnt >= 2 {
				return false
			}
		}
	}
	return true
}
